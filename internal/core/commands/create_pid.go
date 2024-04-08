package commands

import (
	"context"

	"github.com/volatiletech/null/v8"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/common"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
)

type CreatePidCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreatePidCommandHandler(dbs func() *db.ReaderWriter) CreatePidCommandHandler {
	return CreatePidCommandHandler{DBS: dbs}
}

type CreatePidCommandRequest struct {
	ID                   int64
	TemplateName         string
	Header               []byte
	Mode                 []byte
	Pid                  []byte
	Formula              string
	IntervalSeconds      int32
	Protocol             string
	SignalName           string
	CanFlowControlClear  *bool
	CanFlowControlIDPair *string
}

type CreatePidCommandResponse struct {
	ID int64
}

func (h CreatePidCommandHandler) Execute(ctx context.Context, req *CreatePidCommandRequest) (*CreatePidCommandResponse, error) {

	exists, err := models.PidConfigs(
		models.PidConfigWhere.SignalName.EQ(req.SignalName),
		models.PidConfigWhere.TemplateName.EQ(req.TemplateName),
	).Exists(ctx, h.DBS().Reader)

	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error checking if pid config exists: %s", req.TemplateName),
		}
	}
	if exists {
		return nil, &exceptions.ConflictError{
			Err: errors.Errorf("pid config already exists: %s", req.TemplateName),
		}
	}

	pid := &models.PidConfig{
		TemplateName:         req.TemplateName,
		Header:               req.Header,
		Mode:                 req.Mode,
		Pid:                  req.Pid,
		Formula:              common.PrependFormulaTypeDefault(req.Formula),
		IntervalSeconds:      int(req.IntervalSeconds),
		Protocol:             req.Protocol,
		SignalName:           req.SignalName,
		CanFlowControlClear:  null.BoolFromPtr(req.CanFlowControlClear),
		CanFlowControlIDPair: null.StringFromPtr(req.CanFlowControlIDPair),
	}

	err = pid.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error inserting pid with id: %d", req.ID),
		}
	}

	return &CreatePidCommandResponse{ID: pid.ID}, nil
}
