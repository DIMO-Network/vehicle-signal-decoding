package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/null/v8"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UpdatePidCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdatePidCommandHandler(dbs func() *db.ReaderWriter) UpdatePidCommandHandler {
	return UpdatePidCommandHandler{DBS: dbs}
}

type UpdatePidCommandRequest struct {
	ID                   int64
	TemplateName         string
	Header               []byte
	Mode                 []byte
	Pid                  []byte
	Formula              string
	IntervalSeconds      int32
	Protocol             *string
	SignalName           string
	CanFlowControlClear  *bool
	CanFlowControlIDPair *string
	Enabled              bool
	VSSCovesaSignalName  *string
	Unit                 *string
}

type UpdatePidCommandResponse struct {
	ID int64
}

func (h UpdatePidCommandHandler) Execute(ctx context.Context, req *UpdatePidCommandRequest) (*UpdatePidCommandResponse, error) {

	pid, err := models.PidConfigs(models.PidConfigWhere.ID.EQ(req.ID)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("pid config not found for ID: %d", req.ID),
			}
		}
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	pid.ID = req.ID
	pid.TemplateName = req.TemplateName
	pid.Header = req.Header
	pid.Mode = req.Mode
	pid.Pid = req.Pid
	pid.Formula = req.Formula
	pid.IntervalSeconds = int(req.IntervalSeconds)
	pid.Protocol = null.StringFromPtr(req.Protocol)
	pid.SignalName = req.SignalName
	pid.Enabled = req.Enabled
	pid.CanFlowControlClear = null.BoolFromPtr(req.CanFlowControlClear)
	pid.CanFlowControlIDPair = null.StringFromPtr(req.CanFlowControlIDPair)
	pid.VSSCovesaName = null.StringFromPtr(req.VSSCovesaSignalName)
	pid.Unit = null.StringFromPtr(req.Unit)

	if _, err := pid.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdatePidCommandResponse{ID: pid.ID}, nil
}
