package commands

import (
	"context"
	"os"
	"testing"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"

	"github.com/rs/zerolog"

	mockService "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/mocks"
	dbtesthelper "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/stretchr/testify/require"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/testcontainers/testcontainers-go"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type RunTestSignalTestSuite struct {
	suite.Suite
	*require.Assertions

	ctrl                  *gomock.Controller
	pdb                   db.Store
	container             testcontainers.Container
	ctx                   context.Context
	handler               RunTestSignalCommandHandler
	mockUserDeviceService *mockService.MockUserDevicesService
}

func TestRunTestSignalCommandHandler(t *testing.T) {
	suite.Run(t, new(RunTestSignalTestSuite))
}

func (s *RunTestSignalTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.ctx = context.Background()
	s.mockUserDeviceService = mockService.NewMockUserDevicesService(s.ctrl)
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)

	s.handler = NewRunTestSignalCommandHandler(s.pdb.DBS, zerolog.New(os.Stdout), s.mockUserDeviceService)
}

func (s *RunTestSignalTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, dbName, s.T())
	s.ctrl.Finish()
}

func (s *RunTestSignalTestSuite) Test_RunTestSignal() {
	type tableTestCases struct {
		description string
		command     *RunTestSignalCommandRequest
		isError     bool
	}

	const (
		autoPIUnitID = "3"
	)
	deviceDefinitionID := ksuid.New().String()
	userDeviceID := ksuid.New().String()

	userDeviceMock := &appmodels.UserDeviceAutoPIUnit{
		UserDeviceID:       userDeviceID,
		DeviceDefinitionID: deviceDefinitionID,
	}

	s.mockUserDeviceService.EXPECT().GetUserDeviceByAutoPIUnitID(s.ctx, gomock.Any()).Return(userDeviceMock, nil).Times(1)

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
			isError: false,
		},
	} {
		s.T().Run(scenario.description, func(t *testing.T) {
			err := s.handler.Execute(s.ctx, scenario.command)
			if scenario.isError {
				s.Error(err)
			} else {
				s.NoError(err)
				all, err := models.TestSignals().All(context.Background(), s.pdb.DBS().Reader)
				require.NoError(t, err)
				assert.Len(t, all, 3)
				assert.Equal(t, userDeviceID, all[0].UserDeviceID)
			}
		})
	}
}
