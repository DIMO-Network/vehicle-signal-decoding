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
		dbcName = "dbc name"
	)

	for _, scenario := range []tableTestCases{
		{
			description: "Create dbc code success",
			command: &UpdateDBCCodeCommandRequest{
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
				assert.Equal(t, scenario.expected, result.ID)
			}

		})
	}
}
