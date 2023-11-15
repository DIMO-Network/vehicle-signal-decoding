package commands

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	dbtesthelper "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"

	"github.com/DIMO-Network/shared/db"
	"github.com/testcontainers/testcontainers-go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type UpdateDbcCodeTestSuite struct {
	suite.Suite
	*require.Assertions

	ctrl      *gomock.Controller
	pdb       db.Store
	container testcontainers.Container
	ctx       context.Context
	handler   UpdateDBCCodeCommandHandler
}

func TestUpdateDbcCodeCommandHandler(t *testing.T) {
	suite.Run(t, new(UpdateDbcCodeTestSuite))
}

func (s *UpdateDbcCodeTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.ctx = context.Background()
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)

	s.handler = NewUpdateDBCCodeCommandHandler(s.pdb.DBS)
}

func (s *UpdateDbcCodeTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, s.T())
	s.ctrl.Finish()
}

func (s *UpdateDbcCodeTestSuite) Test_UpdateDbcCode() {
	type tableTestCases struct {
		description string
		command     *UpdateDBCCodeCommandRequest
		expected    string
		isError     bool
	}

	const (
		dbcName       = "db_code_name"
		dbcNameUpdate = "db_code_name update"
	)

	dbCode := setupCreateDbcCode(s.T(), dbcName, s.pdb)

	for _, scenario := range []tableTestCases{
		{
			description: "Update dbc code success and does not allow updating name",
			command: &UpdateDBCCodeCommandRequest{
				ID:               dbCode.ID,
				Name:             dbcNameUpdate,
				RecordingEnabled: dbCode.RecordingEnabled,
				Trigger:          dbCode.Trigger,
				MaxSampleSize:    int32(dbCode.MaxSampleSize),
				Header:           dbCode.Header.Int,
				DBCContents:      dbCode.DBCContents.String,
			},
			expected: dbcName, // does not allow updating name
			isError:  false,
		},
		{
			description: "Update dbc code with not found dbc_code",
			command: &UpdateDBCCodeCommandRequest{
				Name:             dbcNameUpdate,
				RecordingEnabled: dbCode.RecordingEnabled,
				Trigger:          dbCode.Trigger,
				MaxSampleSize:    int32(dbCode.MaxSampleSize),
				Header:           dbCode.Header.Int,
				DBCContents:      dbCode.DBCContents.String,
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
				assert.Equal(t, scenario.expected, result.Name)
			}

		})
	}
}
