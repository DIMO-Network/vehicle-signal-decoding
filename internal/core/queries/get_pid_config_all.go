package queries

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetPidAllQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetPidAllQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetPidAllQueryHandler {
	return GetPidAllQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetPidAllQueryRequest struct {
	TemplateName string
}

// Handle gets all pids in a template and their children from inherited parent templates
func (h *GetPidAllQueryHandler) Handle(ctx context.Context, request *GetPidAllQueryRequest) (models.PidConfigSlice, error) {
	template, err := models.FindTemplate(ctx, h.DBS().Reader, request.TemplateName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no template with name: %s found", request.TemplateName)
		}
		return nil, errors.Wrapf(err, "failed to retrieve Template %s", request.TemplateName)
	}

	var templates []models.Template
	currentTemplate := template
	for {
		templates = append(templates, *currentTemplate)

		if currentTemplate.ParentTemplateName.Valid {
			currentTemplate, err = models.FindTemplate(ctx, h.DBS().Reader, currentTemplate.ParentTemplateName.String)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					break
				}
				return nil, errors.Wrapf(err, "failed to retrieve parent Template %s", currentTemplate.ParentTemplateName.String)
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
	).All(ctx, h.DBS().Reader)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "Failed to retrieve PID Configs")
	}
	// todo what about deduping
	return pidConfigs, nil
}
