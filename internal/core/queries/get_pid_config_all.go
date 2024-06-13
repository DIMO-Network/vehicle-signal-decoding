package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
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

func (h *GetPidAllQueryHandler) Handle(ctx context.Context, request *GetPidAllQueryRequest) (*grpc.GetPidListResponse, error) {

	template, err := models.Templates(models.TemplateWhere.TemplateName.EQ(request.TemplateName)).One(ctx, h.DBS().Reader)

	if err != nil {
		return nil, fmt.Errorf("invalid template: %w", err)
	}

	currentTemplatePids, err := h.getPidsByTemplate(ctx, template.TemplateName)

	if err != nil {
		return nil, fmt.Errorf("failed to get PidConfigs: %w", err)
	}

	parentTemplatePids := make([]*grpc.PidSummary, 0)

	if template.ParentTemplateName.Valid && len(template.ParentTemplateName.String) > 0 {
		parentTemplatePids, err = h.getPidsByTemplate(ctx, template.ParentTemplateName.String)
		if err != nil {
			return nil, fmt.Errorf("failed to get PidConfigs: %w", err)
		}
	}

	pidSummaries := make([]*grpc.PidSummary, 0)

	pidSummaries = append(pidSummaries, currentTemplatePids...)

	for _, item := range parentTemplatePids {
		item.IsParentPid = true
		pidSummaries = append(pidSummaries, item)
	}

	result := &grpc.GetPidListResponse{
		Pid: pidSummaries,
	}

	return result, nil
}

func (h *GetPidAllQueryHandler) getPidsByTemplate(ctx context.Context, templateName string) ([]*grpc.PidSummary, error) {
	allPidConfigs, err := models.PidConfigs(
		models.PidConfigWhere.TemplateName.EQ(templateName),
		qm.OrderBy(models.PidConfigColumns.SignalName),
	).All(ctx, h.DBS().Reader)

	if err != nil {
		return nil, fmt.Errorf("failed to get PidConfigs: %w", err)
	}

	pidSummaries := make([]*grpc.PidSummary, 0)

	for _, item := range allPidConfigs {
		pidSummaries = append(pidSummaries, &grpc.PidSummary{
			Id:                   item.ID,
			TemplateName:         item.TemplateName,
			Header:               item.Header,
			Mode:                 item.Mode,
			Pid:                  item.Pid,
			Formula:              item.Formula,
			IntervalSeconds:      int32(item.IntervalSeconds),
			Protocol:             item.Protocol.String,
			SignalName:           item.SignalName,
			CanFlowControlClear:  item.CanFlowControlClear.Bool,
			CanFlowControlIdPair: item.CanFlowControlIDPair.String,
			Enabled:              item.Enabled,
			VssCovesaName:        item.VSSCovesaName.String,
			Unit:                 item.Unit.String,
		})
	}

	return pidSummaries, nil
}
