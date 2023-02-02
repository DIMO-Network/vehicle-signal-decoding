package commands

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	dbtesthelper "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"

	"github.com/DIMO-Network/shared/db"
	"github.com/testcontainers/testcontainers-go"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CreateTestSignalTestSuite struct {
	suite.Suite
	*require.Assertions

	ctrl      *gomock.Controller
	pdb       db.Store
	container testcontainers.Container
	ctx       context.Context
	handler   CreateTestSignalCommandHandler
}

func TestCreateTestSignalCommandHandler(t *testing.T) {
	suite.Run(t, new(CreateTestSignalTestSuite))
}

func (s *CreateTestSignalTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.ctx = context.Background()
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)

	s.handler = NewCreateTestSignalCommandHandler(s.pdb.DBS)
}

func (s *CreateTestSignalTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, s.T())
	s.ctrl.Finish()
}

func (s *CreateTestSignalTestSuite) Test_CreateTestSignal() {
	type tableTestCases struct {
		description string
		command     *CreateTestSignalCommandRequest
		expected    string
		isError     bool
	}

	const (
		deviceDefinitionID = "1"
		userDeviceID       = "1"
		autoPIUnitID       = "1"
		value              = ""
		approved           = true
	)

	for _, scenario := range []tableTestCases{
		{
			description: "Create test signal success",
			command: &CreateTestSignalCommandRequest{
				DeviceDefinitionID: deviceDefinitionID,
				UserDeviceID:       userDeviceID,
				AutoPIUnitID:       autoPIUnitID,
				Value:              value,
				Approved:           approved,
			},
			expected: "dbcName",
			isError:  false,
		},
	} {
		s.T().Run(scenario.description, func(t *testing.T) {
			result, err := s.handler.Execute(s.ctx, scenario.command)
			if scenario.isError {
				s.Nil(result)
				s.Error(err)
			} else {
				assert.Equal(t, scenario.expected, result.ID)
			}

		})
	}
}
