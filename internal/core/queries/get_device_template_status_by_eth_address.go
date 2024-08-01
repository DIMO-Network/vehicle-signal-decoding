package queries

import (
	"context"
	"database/sql"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetDeviceTemplateStatusByEthAddressQuery struct {
	EthAddress common.Address
}

type GetDeviceTemplateStatusByEthAddressQueryHandler struct {
	DBS func() *db.ReaderWriter
}

func NewGetDeviceTemplateStatusByEthAddressQuery(dbs func() *db.ReaderWriter) *GetDeviceTemplateStatusByEthAddressQueryHandler {
	return &GetDeviceTemplateStatusByEthAddressQueryHandler{DBS: dbs}
}

func (h *GetDeviceTemplateStatusByEthAddressQueryHandler) Handle(ctx context.Context, q GetDeviceTemplateStatusByEthAddressQuery) (*grpc.GetDeviceTemplateStatusResponse, error) {
	deviceTemplateStatus, err := models.DeviceTemplateStatuses(
		models.DeviceTemplateStatusWhere.DeviceEthAddr.EQ(q.EthAddress.Bytes())).
		One(ctx, h.DBS().Reader)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	response := &grpc.GetDeviceTemplateStatusResponse{
		TemplateDbcUrl:      deviceTemplateStatus.TemplateSettingsURL.String,
		TemplatePidUrl:      deviceTemplateStatus.TemplatePidURL.String,
		FirmwareVersion:     deviceTemplateStatus.FirmwareVersion.String,
		TemplateSettingsUrl: deviceTemplateStatus.TemplateSettingsURL.String,
		UpdatedAt:           timestamppb.New(deviceTemplateStatus.UpdatedAt),
	}

	return response, nil
}
