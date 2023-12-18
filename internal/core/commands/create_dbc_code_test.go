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

const (
	dbName               = "vehicle_signal_decoding_api"
	migrationsDirRelPath = "../../infrastructure/db/migrations"
)

type CreateDbcCodeTestSuite struct {
	suite.Suite
	*require.Assertions

	ctrl      *gomock.Controller
	pdb       db.Store
	container testcontainers.Container
	ctx       context.Context
	handler   CreateDBCCodeCommandHandler
}

func TestCreateDBCCodeCommandHandler(t *testing.T) {
	suite.Run(t, new(CreateDbcCodeTestSuite))
}

func (s *CreateDbcCodeTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.ctx = context.Background()
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)

	s.handler = NewCreateDBCCodeCommandHandler(s.pdb.DBS)
}

func (s *CreateDbcCodeTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, dbName, s.T())
	s.ctrl.Finish()
}

func (s *CreateDbcCodeTestSuite) Test_CreateDbcCode() {
	type tableTestCases struct {
		description string
		command     *CreateDBCCodeCommandRequest
		expected    string
		isError     bool
	}

	const (
		dbcName = "dbc name"
	)

	for _, scenario := range []tableTestCases{
		{
			description: "Create dbc code success",
			command: &CreateDBCCodeCommandRequest{
				Name: dbcName,
			},
			expected: dbcName,
			isError:  false,
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
