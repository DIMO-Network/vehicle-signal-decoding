package commands

import (
	"context"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/models"
	"testing"

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

	s.handler = NewRunTestSignalCommandHandler(s.pdb.DBS, s.mockUserDeviceService)
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
		userDeviceID       = "1"
		autoPIUnitID       = "1"
	)

	userDeviceMock := &models.UserDeviceAutoPIUnit{
		UserDeviceID:       userDeviceID,
		DeviceDefinitionID: deviceDefinitionID,
	}

	s.mockUserDeviceService.EXPECT().GetUserDeviceServiceByAutoPIUnitID(s.ctx, gomock.Any()).Return(userDeviceMock, nil).Times(1)

	for _, scenario := range []tableTestCases{
		{
			description: "Run test signal success",
			command:     &RunTestSignalCommandRequest{},
			expected:    "dbcName",
			isError:     false,
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
