package commands

import (
	"context"

	"github.com/DIMO-Network/shared/db"
)

type RunTestSignalCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewRunTestSignalCommandHandler(dbs func() *db.ReaderWriter) RunTestSignalCommandHandler {
	return RunTestSignalCommandHandler{DBS: dbs}
}

type RunTestSignalCommandRequest struct {
}

type RunTestSignalCommandResponse struct {
}

func (h RunTestSignalCommandHandler) Execute(ctx context.Context, command *RunTestSignalCommandRequest) error {
	return nil
}
