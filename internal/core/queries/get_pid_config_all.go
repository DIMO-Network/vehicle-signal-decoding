package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetPidsQueryRequest struct {
	TemplateName string
}

// GetPidsByTemplate gets all pids in a template and their children from inherited parent templates
func GetPidsByTemplate(ctx context.Context, dbs func() *db.ReaderWriter, request *GetPidsQueryRequest) (models.PidConfigSlice, *models.Template, error) {
	templateNames, template, err := GetAllParentTemplates(ctx, dbs, request.TemplateName)
	if err != nil {
		return nil, nil, err
	}

	pidConfigs, err := models.PidConfigs(
		qm.WhereIn("template_name IN ?", ToAnySlice(templateNames)...),
	).All(ctx, dbs().Reader)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, nil, errors.Wrap(err, "Failed to retrieve PID Configs")
	}
	// dedupe any signals, child overrides parent
	seen := make(map[string]bool)
	result := []*models.PidConfig{}
	dupes := make(map[string]bool)
	// first get the dupes and the seend
	for _, pid := range pidConfigs {
		if _, ok := seen[pid.SignalName]; !ok {
			seen[pid.SignalName] = true
		} else {
			dupes[pid.SignalName] = true
		}
	}

	// make sure child always gets added
	for _, pid := range pidConfigs {
		if _, ok := dupes[pid.SignalName]; ok {
			if pid.TemplateName == request.TemplateName {
				result = append(result, pid)
			}
		} else {
			result = append(result, pid)
		}
	}

	return result, template, nil
}

// GetAllParentTemplates gets all the parent templates for a template, if none returns empty just the current template name
func GetAllParentTemplates(ctx context.Context, dbs func() *db.ReaderWriter, templateName string) ([]string, *models.Template, error) {
	template, err := models.FindTemplate(ctx, dbs().Reader, templateName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, fmt.Errorf("no template with name: %s found", templateName)
		}
		return nil, nil, errors.Wrapf(err, "failed to retrieve Template %s", templateName)
	}

	var templates []models.Template
	currentTemplate := template
	for {
		templates = append(templates, *currentTemplate)

		if currentTemplate.ParentTemplateName.Valid {
			currentTemplate, err = models.FindTemplate(ctx, dbs().Reader, currentTemplate.ParentTemplateName.String)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					break
				}
				return nil, nil, errors.Wrapf(err, "failed to retrieve parent Template %s", currentTemplate.ParentTemplateName.String)
			}
		} else {
			break
		}
	}

	templateNames := make([]string, len(templates))
	for i, tmpl := range templates {
		templateNames[i] = tmpl.TemplateName
	}
	return templateNames, template, nil
}

func ToAnySlice(slice []string) []any {
	anySlice := make([]any, len(slice))
	for i, v := range slice {
		anySlice[i] = v
	}
	return anySlice
}
