package queries

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/volatiletech/null/v8"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"

	"github.com/DIMO-Network/shared/pkg/db"
	pgrpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetTestSignalAllQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetTestSignalAllQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetTestSignalAllQueryHandler {
	return GetTestSignalAllQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetTestSignalAllQueryRequest struct {
	Input string
}

func (h GetTestSignalAllQueryHandler) Handle(ctx context.Context, _ *GetTestSignalAllQueryRequest) (*pgrpc.GetTestSignalListResponse, error) {

	all, err := models.TestSignals().All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to get test signals"),
		}
	}

	result := &pgrpc.GetTestSignalListResponse{}

	for _, item := range all {
		result.Items = append(result.Items, &pgrpc.GetTestSignalResponse{
			Id:                 item.ID,
			UserDeviceId:       item.UserDeviceID,
			DeviceDefinitionId: item.DeviceDefinitionID,
			DbcCodesId:         item.DBCCodesID,
			Value:              item.Value,
			AutopiUnitId:       item.AutopiUnitID,
			Approved:           item.Approved,
			Signals:            string(JSONOrDefault(item.Signals)),
		})
	}

	return result, nil
}

func JSONOrDefault(j null.JSON) json.RawMessage {
	if !j.Valid || len(j.JSON) == 0 {
		return []byte(`{}`)
	}
	return j.JSON
}
