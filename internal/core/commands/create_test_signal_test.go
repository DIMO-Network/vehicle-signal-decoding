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
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, dbName, s.T())
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

	dbCode := setupCreateDbcCode(s.T(), "db_code_name", s.pdb)

	for _, scenario := range []tableTestCases{
		{
			description: "Create test signal success",
			command: &CreateTestSignalCommandRequest{
				DBCCodesID:         dbCode.ID,
				DeviceDefinitionID: deviceDefinitionID,
				UserDeviceID:       userDeviceID,
				AutoPIUnitID:       autoPIUnitID,
				Value:              value,
				Approved:           approved,
			},
			expected: dbCode.ID,
			isError:  false,
		},
		{
			description: "Create test signal with bad request dbc_code",
			command: &CreateTestSignalCommandRequest{
				DBCCodesID:         "",
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
				assert.Equal(t, scenario.expected, dbCode.ID)
			}

		})
	}
}

func setupCreateDbcCode(t *testing.T, name string, pdb db.Store) models.DBCCode {
	dbcCode := models.DBCCode{
		ID:   ksuid.New().String(),
		Name: name,
	}
	err := dbcCode.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err, "database error")
	return dbcCode
}
