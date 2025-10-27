package commands

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/aarondl/sqlboiler/v4/boil"
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

	exists, err := models.AftermarketDeviceToTemplates(
		models.AftermarketDeviceToTemplateWhere.AftermarketDeviceEthereumAddress.EQ(cmd.AftermarketDeviceEthereumAddress),
	).Exists(ctx, h.DBS().Reader)

	if err != nil {
		return &exceptions.InternalError{
			Err: errors.Wrapf(err, "error checking if already a template registered for %s", cmd.AftermarketDeviceEthereumAddress),
		}
	}

	if exists {
		return &exceptions.ConflictError{
			Err: errors.Errorf("Already a template registered for: %s", cmd.AftermarketDeviceEthereumAddress),
		}
	}

	aftermarketDeviceTemplate := &models.AftermarketDeviceToTemplate{
		AftermarketDeviceEthereumAddress: cmd.AftermarketDeviceEthereumAddress,
		TemplateName:                     cmd.TemplateName,
	}

	err = aftermarketDeviceTemplate.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}
