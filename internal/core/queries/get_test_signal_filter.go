package queries

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/common"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"

	"github.com/DIMO-Network/shared/db"
	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetTestSignalFilterQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetTestSignalFilterQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetTestSignalFilterQueryHandler {
	return GetTestSignalFilterQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetTestSignalFilterQueryRequest struct {
	DeviceDefinitionID string
	DBCCodeID          string
	UserDeviceID       string
}

func (h GetTestSignalFilterQueryHandler) Handle(ctx context.Context, query *GetTestSignalFilterQueryRequest) (*p_grpc.GetTestSignalListResponse, error) {

	var queryMods []qm.QueryMod

	if len(query.DeviceDefinitionID) > 1 {
		queryMods = append(queryMods, models.TestSignalWhere.DeviceDefinitionID.EQ(string(query.DeviceDefinitionID)))
	}

	if len(query.UserDeviceID) > 1 {
		queryMods = append(queryMods, models.TestSignalWhere.UserDeviceID.EQ(string(query.UserDeviceID)))
	}

	if len(query.DBCCodeID) > 1 {
		queryMods = append(queryMods, models.TestSignalWhere.DBCCodesID.EQ(string(query.DBCCodeID)))
	}

	queryMods = append(queryMods,
		qm.Load(models.TestSignalRels.DBCCode),
		qm.OrderBy("device_definition_id ASC, autopi_unit_id ASC, created_at DESC"))

	all, err := models.TestSignals(queryMods...).All(ctx, h.DBS().Reader)

	if err != nil {
		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to get test signals"),
		}
	}

	result := &p_grpc.GetTestSignalListResponse{}

	for _, item := range all {
		result.Items = append(result.Items, &p_grpc.GetTestSignalResponse{
			Id:                 item.ID,
			UserDeviceId:       item.UserDeviceID,
			DeviceDefinitionId: item.DeviceDefinitionID,
			DbcCodesId:         item.DBCCodesID,
			Value:              item.Value,
			AutopiUnitId:       item.AutopiUnitID,
			Approved:           item.Approved,
			Signals:            string(common.JSONOrDefault(item.Signals)),
			CreatedAt:          timestamppb.New(item.CreatedAt),
		})

	}

	return result, nil
}
