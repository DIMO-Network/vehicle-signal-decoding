package queries

import (
	"context"
	"testing"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
	dbtesthelper "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"go.uber.org/mock/gomock"
)

const (
	dbName               = "vehicle_signal_decoding_api"
	migrationsDirRelPath = "../../infrastructure/db/migrations"
)

type GetDeviceTemplateStatusByEthAddressTestSuite struct {
	suite.Suite
	*require.Assertions

	ctrl      *gomock.Controller
	pdb       db.Store
	container testcontainers.Container
	ctx       context.Context
	handler   *GetDeviceTemplateStatusByEthAddressQueryHandler
}

func TestGetDeviceTemplateStatusByEthAddressQueryHandler(t *testing.T) {
	suite.Run(t, new(GetDeviceTemplateStatusByEthAddressTestSuite))
}

func (s *GetDeviceTemplateStatusByEthAddressTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.ctx = context.Background()
	s.pdb, s.container = dbtesthelper.StartContainerDatabase(s.ctx, dbName, s.T(), migrationsDirRelPath)

	s.handler = NewGetDeviceTemplateStatusByEthAddressQuery(s.pdb.DBS)
}

func (s *GetDeviceTemplateStatusByEthAddressTestSuite) TearDownTest() {
	dbtesthelper.TruncateTables(s.pdb.DBS().Writer.DB, dbName, s.T())
	s.ctrl.Finish()
}

func (s *GetDeviceTemplateStatusByEthAddressTestSuite) Test_GetDeviceTemplateStatus() {

	ethAddr := "DC1eE274BCA98b421293f3737D1b9E4563c60cb3"
	address := common2.HexToAddress(ethAddr)

	PidURL := "http://localhost:3000/v1/device-config/pids/some-template-emptydbc@v1.0.0"
	DeviceSettingURL := "http://localhost:3000/v1/device-config/settings/default-hev-emptydbc@v1.0.0"
	DbcURL := ""

	setupCreateDeviceTemplateStatus(s.T(), address, PidURL, DeviceSettingURL, DbcURL, s.pdb)

	ethAddr2 := "98D78d711C0ec544F6fb5d54fcf6559CF41546a9"
	address2 := common2.HexToAddress(ethAddr2)
	setupCreateDeviceTemplateStatusWithEmptyValues(s.T(), address2, s.pdb)

	ethAddr3 := "7219ED580Bf8894646033333cCa7045B5CEa58e3"
	address3 := common2.HexToAddress(ethAddr3)

	type tableTestCases struct {
		description string
		query       *GetDeviceTemplateStatusByEthAddressQuery
		expected    *grpc.GetDeviceTemplateStatusResponse
		isError     bool
	}

	for _, scenario := range []tableTestCases{
		{
			description: "Get device template status",
			query: &GetDeviceTemplateStatusByEthAddressQuery{
				EthAddress: address,
			},
			expected: &grpc.GetDeviceTemplateStatusResponse{
				TemplatePidUrl:      PidURL,
				TemplateSettingsUrl: DeviceSettingURL,
				TemplateDbcUrl:      DbcURL,
			},
			isError: false,
		},
		{
			description: "Get device template status with empty values",
			query: &GetDeviceTemplateStatusByEthAddressQuery{
				EthAddress: address2,
			},
			expected: &grpc.GetDeviceTemplateStatusResponse{},
			isError:  false,
		},
		{
			description: "Device template status not exists",
			query: &GetDeviceTemplateStatusByEthAddressQuery{
				EthAddress: address3,
			},
			expected: &grpc.GetDeviceTemplateStatusResponse{},
			isError:  false,
		},
	} {
		s.T().Run(scenario.description, func(t *testing.T) {
			result, err := s.handler.Handle(s.ctx, *scenario.query)
			if scenario.isError {
				s.Nil(result)
				s.Error(err)
			} else {
				assert.Equal(t, scenario.expected.TemplatePidUrl, result.TemplatePidUrl)
				assert.Equal(t, scenario.expected.TemplateSettingsUrl, result.TemplateSettingsUrl)
				assert.Equal(t, scenario.expected.TemplateDbcUrl, result.TemplateDbcUrl)
			}

		})
	}
}

func setupCreateDeviceTemplateStatus(t *testing.T, address common2.Address, pid, setting, dbc string, pdb db.Store) {
	deviceTemplateStatus := models.DeviceTemplateStatus{
		DeviceEthAddr:       address.Bytes(),
		TemplatePidURL:      null.StringFrom(pid),
		TemplateSettingsURL: null.StringFrom(setting),
		TemplateDBCURL:      null.StringFrom(dbc),
	}
	err := deviceTemplateStatus.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err, "database error")
}

func setupCreateDeviceTemplateStatusWithEmptyValues(t *testing.T, address common2.Address, pdb db.Store) {
	deviceTemplateStatus := models.DeviceTemplateStatus{
		DeviceEthAddr: address.Bytes(),
	}
	err := deviceTemplateStatus.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err, "database error")
}
