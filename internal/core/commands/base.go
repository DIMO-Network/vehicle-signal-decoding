package commands

import (
	"context"
	"fmt"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"

	"github.com/DIMO-Network/shared/db"
)

type BaseCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewBaseCommandHandler(dbs func() *db.ReaderWriter) BaseCommandHandler {
	return BaseCommandHandler{DBS: dbs}
}

type BaseCommandRequest struct {
	Input string
}

type BaseCommandResponse struct {
	Result bool
}

func (h BaseCommandHandler) Execute(ctx context.Context, command *BaseCommandRequest) (*BaseCommandResponse, error) {
	if len(command.Input) == 0 {
		return nil, &exceptions.ValidationError{
			Err: fmt.Errorf("custom error %s", command.Input),
		}
	}

	return &BaseCommandResponse{Result: true}, nil
}
