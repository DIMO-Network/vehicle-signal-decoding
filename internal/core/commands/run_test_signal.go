package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source run_test_signal.go -destination mocks/run_test_signal_mock.go

type RunTestSignalCommandHandler interface {
	Execute(ctx context.Context, command *RunTestSignalCommandRequest) error
}

type runTestSignalCommandHandler struct {
	DBS    func() *db.ReaderWriter
	logger zerolog.Logger
}

// NewRunTestSignalCommandHandler creates a handler for run_test_signal commands.
// Device lookup by AutoPI unit ID is no longer supported (devices-api deprecated); Execute returns an error for that path.
func NewRunTestSignalCommandHandler(dbs func() *db.ReaderWriter, logger zerolog.Logger) RunTestSignalCommandHandler {
	return runTestSignalCommandHandler{DBS: dbs, logger: logger}
}

type RunTestSignalCommandRequest struct {
	AutoPIUnitID string
	Time         time.Time
	Signals      map[string]RunTestSignalItemCommandRequest
}

type RunTestSignalItemCommandRequest struct {
	Value any    `json:"value"`
	Time  string `json:"_stamp"` //nolint
}

type RunTestSignalCommandResponse struct {
}

func (h runTestSignalCommandHandler) Execute(_ context.Context, command *RunTestSignalCommandRequest) error {
	// Device lookup by AutoPI unit ID is no longer supported (devices-api deprecated; Identity API does not expose it).
	return fmt.Errorf("run_test_signal by AutoPI unit ID is not supported: device lookup by autopi_unit_id %s is unavailable", command.AutoPIUnitID)
}
