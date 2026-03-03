package commands

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/DIMO-Network/shared/pkg/db"
	dbtesthelper "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/testcontainers/testcontainers-go"
)

type RunTestSignalTestSuite struct {
	suite.Suite
	*require.Assertions

	pdb       db.Store
	container testcontainers.Container
	ctx       context.Context
	handler   RunTestSignalCommandHandler
}

func TestRunTestSignalCommandHandler(t *testing.T) {
	suite.Run(t, new(RunTestSignalTestSuite))
}

func (s *RunTestSignalTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)
	s.handler = NewRunTestSignalCommandHandler(s.pdb.DBS, zerolog.New(os.Stdout))
}

func (s *RunTestSignalTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, dbName, s.T())
}

func (s *RunTestSignalTestSuite) Test_RunTestSignal_ReturnsNotSupported() {
	// RunTestSignalCommandHandler no longer looks up device by AutoPI unit ID (devices-api deprecated).
	err := s.handler.Execute(s.ctx, &RunTestSignalCommandRequest{
		AutoPIUnitID: "3",
		Signals: map[string]RunTestSignalItemCommandRequest{
			"canbus_vin_toyota580v1": {Time: "2023-01-30T15:12:17.464970", Value: "0"},
		},
	})
	s.Error(err)
	s.Contains(err.Error(), "not supported")
	s.Contains(err.Error(), "autopi_unit_id")
}
