package commands

import (
	"context"

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
	TemplateName    string
	Header          []byte
	Mode            []byte
	Pid             []byte
	Formula         string
	IntervalSeconds int32
	Protocol        string
	SignalName      string
	BytesReturned   int32
}

type CreatePidCommandResponse struct {
	Name string
}

func (h CreatePidCommandHandler) Execute(ctx context.Context, req *CreatePidCommandRequest) (*CreatePidCommandResponse, error) {

	exists, err := models.PidConfigs(models.PidConfigWhere.TemplateName.EQ(req.TemplateName)).Exists(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error checking if pid config exists: %s", req.TemplateName),
		}
	}
	if exists {
		return nil, &exceptions.ConflictError{
			Err: errors.Wrapf(err, "pid with name %s already exists", req.TemplateName),
		}
	}

	pid := &models.PidConfig{
		TemplateName:    req.TemplateName,
		Header:          req.Header,
		Mode:            req.Mode,
		Pid:             req.Pid,
		Formula:         req.Formula,
		IntervalSeconds: int(req.IntervalSeconds),
		Protocol:        req.Protocol,
		SignalName:      req.SignalName,
	}

	err = pid.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error inserting pid with template name: %s", req.TemplateName),
		}
	}

	return &CreatePidCommandResponse{Name: pid.TemplateName}, nil
}
