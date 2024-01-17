package services

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"database/sql"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

//go:generate mockgen -source device_template_service.go -destination mocks/device_template_service_mock.go
type DeviceTemplateService interface {
	AssociateTemplate(ctx context.Context, vin, templateDbcURL, templatePidURL, templateSettingURL, version string) (*models.DeviceTemplate, error)
}

type userDeviceTemplateService struct {
	db *sql.DB
}

func NewUserDeviceTemplateService(database *sql.DB) DeviceTemplateService {
	return &userDeviceTemplateService{
		db: database,
	}
}

func (u userDeviceTemplateService) AssociateTemplate(ctx context.Context, vin, templateDbcURL, templatePidURL, templateSettingURL, version string) (*models.DeviceTemplate, error) {

	deviceTemplate, err := models.DeviceTemplates(models.DeviceTemplateWhere.Vin.EQ(vin)).
		One(ctx, u.db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if deviceTemplate != nil {
		deviceTemplate.Version = version
		if deviceTemplate.TemplateDBCURL != templateDbcURL {
			deviceTemplate.IsTemplateUpdated = false
		}
		if deviceTemplate.TemplatePidURL != templatePidURL {
			deviceTemplate.IsTemplateUpdated = false
		}
		if deviceTemplate.TemplateSettingURL != templateSettingURL {
			deviceTemplate.IsTemplateUpdated = false
		}

		if _, err = deviceTemplate.Update(ctx, u.db, boil.Infer()); err != nil {
			return nil, err
		}
	}

	if deviceTemplate == nil {
		deviceTemplate = &models.DeviceTemplate{
			Vin:                vin,
			TemplateDBCURL:     templateDbcURL,
			TemplatePidURL:     templatePidURL,
			TemplateSettingURL: templateSettingURL,
			IsTemplateUpdated:  true,
		}

		if err = deviceTemplate.Insert(ctx, u.db, boil.Infer()); err != nil {
			return nil, err
		}
	}

	return deviceTemplate, nil
}
