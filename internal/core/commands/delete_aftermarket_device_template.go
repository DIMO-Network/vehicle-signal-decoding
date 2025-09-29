package commands

import (
	"context"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

type DeleteAftermarketDeviceTemplateCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type DeleteAftermarketDeviceTemplateCommand struct {
	AftermarketDeviceEthereumAddress []byte
	TemplateName                     string
}

func NewDeleteAftermarketDeviceTemplateCommandHandler(dbs func() *db.ReaderWriter) *DeleteAftermarketDeviceTemplateCommandHandler {
	return &DeleteAftermarketDeviceTemplateCommandHandler{DBS: dbs}
}

func (h *DeleteAftermarketDeviceTemplateCommandHandler) Execute(ctx context.Context, cmd DeleteAftermarketDeviceTemplateCommand) error {
	_, err := models.AftermarketDeviceToTemplates(
		models.AftermarketDeviceToTemplateWhere.AftermarketDeviceEthereumAddress.EQ(cmd.AftermarketDeviceEthereumAddress),
	).DeleteAll(ctx, h.DBS().Writer)

	if err != nil {
		return err
	}

	return nil
}
