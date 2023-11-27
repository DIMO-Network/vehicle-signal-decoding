package commands

import (
	"context"
	"testing"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/stretchr/testify/require"

	dbtesthelper "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"

	"github.com/DIMO-Network/shared/db"
	"github.com/testcontainers/testcontainers-go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type UpdateTestSignalTestSuite struct {
	suite.Suite
	*require.Assertions

	ctrl      *gomock.Controller
	pdb       db.Store
	container testcontainers.Container
	ctx       context.Context
	handler   UpdateTestSignalCommandHandler
}

func TestUpdateTestSignalCommandHandler(t *testing.T) {
	suite.Run(t, new(UpdateTestSignalTestSuite))
}

func (s *UpdateTestSignalTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.ctx = context.Background()
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)

	s.handler = NewUpdateTestSignalCommandHandler(s.pdb.DBS)
}

func (s *UpdateTestSignalTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, s.T())
	s.ctrl.Finish()
}

func (s *UpdateTestSignalTestSuite) Test_UpdateTestSignal() {
	type tableTestCases struct {
		description string
		command     *UpdateTestSignalCommandRequest
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

	testSignal := setupCreateTestSignal(s.T(), "db_code_name", s.pdb)

	for _, scenario := range []tableTestCases{
		{
			description: "Update test signal success",
			command: &UpdateTestSignalCommandRequest{
				ID:                 testSignal.ID,
				DeviceDefinitionID: deviceDefinitionID,
				UserDeviceID:       userDeviceID,
				AutoPIUnitID:       autoPIUnitID,
				Value:              value,
				Approved:           approved,
			},
			expected: testSignal.ID,
			isError:  false,
		},
		{
			description: "Update test signal with not found dbc_code",
			command: &UpdateTestSignalCommandRequest{
				DeviceDefinitionID: deviceDefinitionID,
				UserDeviceID:       userDeviceID,
				AutoPIUnitID:       autoPIUnitID,
				Value:              value,
				Approved:           approved,
			},
			expected: "",
			isError:  true,
		},
	} {
		s.T().Run(scenario.description, func(t *testing.T) {
			result, err := s.handler.Execute(s.ctx, scenario.command)
			if scenario.isError {
				s.Nil(result)
				s.Error(err)
			} else {
				assert.Equal(t, scenario.expected, testSignal.ID)
			}

		})
	}
}

func setupCreateTestSignal(t *testing.T, name string, pdb db.Store) models.TestSignal {
	dbCode := setupCreateDbcCode(t, name, pdb)
	testSignal := models.TestSignal{
		ID:                 ksuid.New().String(),
		DBCCodesID:         dbCode.ID,
		DeviceDefinitionID: "1",
		UserDeviceID:       "1",
		Approved:           true,
		AutopiUnitID:       "1",
	}
	err := testSignal.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err, "database error")
	return testSignal
}
