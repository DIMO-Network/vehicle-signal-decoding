package commands

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
)

type BaseCommandHandler struct {
}

type BaseCommandRequest struct {
	Input string
}

type BaseCommandResponse struct {
	Result bool
}

func (h BaseCommandHandler) Execute(_ context.Context, command *BaseCommandRequest) (*BaseCommandResponse, error) {
	if len(command.Input) == 0 {
		return nil, &exceptions.ValidationError{
			Err: fmt.Errorf("custom error %s", command.Input),
		}
	}

	return &BaseCommandResponse{Result: true}, nil
}
