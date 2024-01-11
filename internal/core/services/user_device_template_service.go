package services

import (
	"context"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/models"

	"database/sql"
)

//go:generate mockgen -source user_device_template_service.go -destination mocks/user_device_template_service_mock.go
type UserDeviceTemplateService interface {
	AssociateTemplate(ctx context.Context, ethAddr string) (models.AssociatedUserDeviceTemplate, error)
}

type userDeviceTemplateService struct {
	db *sql.DB
}

func NewUserDeviceTemplateService(database *sql.DB) UserDeviceTemplateService {
	return &userDeviceTemplateService{
		db: database,
	}
}

func (u userDeviceTemplateService) AssociateTemplate(ctx context.Context, ethAddr string) (models.AssociatedUserDeviceTemplate, error) {
	return models.AssociatedUserDeviceTemplate{RequiresUpdateVersion: false}, nil
}
