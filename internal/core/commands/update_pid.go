package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

type UpdatePidCommandResponse struct {
	Name string
}

func (h UpdatePidCommandHandler) Execute(ctx context.Context, req *UpdatePidCommandRequest) (*UpdatePidCommandResponse, error) {

	pid, err := models.PidConfigs(models.PidConfigWhere.TemplateName.EQ(req.TemplateName)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("pid config not found name: %s", req.TemplateName),
			}
		}
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	pid.TemplateName = req.TemplateName
	pid.Header = req.Header
	pid.Mode = req.Mode
	pid.Pid = req.Pid
	pid.Formula = req.Formula
	pid.IntervalSeconds = int(req.IntervalSeconds)
	pid.Protocol = req.Protocol
	pid.SignalName = req.SignalName

	if _, err := pid.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdatePidCommandResponse{Name: pid.TemplateName}, nil
}
