package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetPidsQueryRequest struct {
	TemplateName string
}

// GetPidsByTemplate gets all pids in a template and their children from inherited parent templates
func GetPidsByTemplate(ctx context.Context, dbs func() *db.ReaderWriter, request *GetPidsQueryRequest) (models.PidConfigSlice, *models.Template, error) {
	template, err := models.FindTemplate(ctx, dbs().Reader, request.TemplateName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, fmt.Errorf("no template with name: %s found", request.TemplateName)
		}
		return nil, nil, errors.Wrapf(err, "failed to retrieve Template %s", request.TemplateName)
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

	templateNames := make([]interface{}, len(templates))
	for i, tmpl := range templates {
		templateNames[i] = tmpl.TemplateName
	}

	pidConfigs, err := models.PidConfigs(
		qm.WhereIn("template_name IN ?", templateNames...),
	).All(ctx, dbs().Reader)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, nil, errors.Wrap(err, "Failed to retrieve PID Configs")
	}
	// todo what about deduping
	return pidConfigs, template, nil
}
