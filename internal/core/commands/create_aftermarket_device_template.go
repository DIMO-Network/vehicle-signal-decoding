package commands

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CreateAftermarketDeviceTemplateCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type CreateAftermarketDeviceTemplateCommand struct {
	AftermarketDeviceEthereumAddress []byte
	TemplateName                     string
}

func NewCreateAftermarketDeviceTemplateCommandHandler(dbs func() *db.ReaderWriter) *CreateAftermarketDeviceTemplateCommandHandler {
	return &CreateAftermarketDeviceTemplateCommandHandler{DBS: dbs}
}

func (h *CreateAftermarketDeviceTemplateCommandHandler) Execute(ctx context.Context, cmd CreateAftermarketDeviceTemplateCommand) error {
	aftermarketDeviceTemplate := &models.AftermarketDeviceToTemplate{
		AftermarketDeviceEthereumAddress: cmd.AftermarketDeviceEthereumAddress,
		TemplateName:                     cmd.TemplateName,
	}

	err := aftermarketDeviceTemplate.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}
