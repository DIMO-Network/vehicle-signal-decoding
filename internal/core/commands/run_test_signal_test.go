package commands

import (
	"context"
	"os"
	"testing"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"

	"github.com/rs/zerolog"

	mockService "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/mocks"
	dbtesthelper "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/stretchr/testify/require"

	"github.com/DIMO-Network/shared/db"
	"github.com/testcontainers/testcontainers-go"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type RunTestSignalTestSuite struct {
	suite.Suite
	*require.Assertions

	ctrl                  *gomock.Controller
	pdb                   db.Store
	container             testcontainers.Container
	ctx                   context.Context
	handler               RunTestSignalCommandHandler
	mockUserDeviceService *mockService.MockUserDeviceService
}

func TestRunTestSignalCommandHandler(t *testing.T) {
	suite.Run(t, new(RunTestSignalTestSuite))
}

func (s *RunTestSignalTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.ctx = context.Background()
	s.mockUserDeviceService = mockService.NewMockUserDeviceService(s.ctrl)
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)

	s.handler = NewRunTestSignalCommandHandler(s.pdb.DBS, zerolog.New(os.Stdout), s.mockUserDeviceService)
}

func (s *RunTestSignalTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, s.T())
	s.ctrl.Finish()
}

func (s *RunTestSignalTestSuite) Test_RunTestSignal() {
	type tableTestCases struct {
		description string
		command     *RunTestSignalCommandRequest
		expected    string
		isError     bool
	}

	const (
		deviceDefinitionID = "1"
		userDeviceID       = "2"
		autoPIUnitID       = "3"
	)

	userDeviceMock := &services.UserDeviceAutoPIUnit{
		UserDeviceID:       userDeviceID,
		DeviceDefinitionID: deviceDefinitionID,
	}

	s.mockUserDeviceService.EXPECT().GetUserDeviceServiceByAutoPIUnitID(s.ctx, gomock.Any()).Return(userDeviceMock, nil).Times(1)

	eventSignals1 := map[string]RunTestSignalItemCommandRequest{}
	eventSignals1["canbus_vin_toyota580v1"] = RunTestSignalItemCommandRequest{
		Time:  "2023-01-30T15:12:17.464970",
		Value: "0",
	}
	eventSignals1["canbus_vin_toyota580v2"] = RunTestSignalItemCommandRequest{
		Time:  "2023-01-30T15:12:17.464970",
		Value: 1,
	}
	eventSignals1["canbus_vin_toyota580v3"] = RunTestSignalItemCommandRequest{
		Time:  "2023-01-30T15:12:17.464970",
		Value: 1.4,
	}

	for _, scenario := range []tableTestCases{
		{
			description: "Run test signal success",
			command: &RunTestSignalCommandRequest{
				AutoPIUnitID: autoPIUnitID,
				Signals:      eventSignals1,
			},
			expected: "dbcName",
			isError:  false,
		},
	} {
		s.T().Run(scenario.description, func(t *testing.T) {
			err := s.handler.Execute(s.ctx, scenario.command)
			if scenario.isError {
				s.Error(err)
			} else {
				s.NoError(err)

			}
		})
	}
}
