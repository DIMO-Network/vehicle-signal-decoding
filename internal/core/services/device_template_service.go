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
	AssociateTemplate(ctx context.Context, vin, templateDbcURL, templatePidURL, templateSettingURL, version string) error
}

type userDeviceTemplateService struct {
	db *sql.DB
}

func NewUserDeviceTemplateService(database *sql.DB) DeviceTemplateService {
	return &userDeviceTemplateService{
		db: database,
	}
}

func (u userDeviceTemplateService) AssociateTemplate(ctx context.Context, vin, templateDbcURL, templatePidURL, templateSettingURL, version string) error {

	userDeviceTemplate, err := models.DeviceTemplates(models.DeviceTemplateWhere.Vin.EQ(vin)).
		One(ctx, u.db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if userDeviceTemplate != nil {
		userDeviceTemplate.Version = version
		if userDeviceTemplate.TemplateDBCURL != templateDbcURL {
			userDeviceTemplate.IsTemplateUpdated = false
		}
		if userDeviceTemplate.TemplatePidURL != templatePidURL {
			userDeviceTemplate.IsTemplateUpdated = false
		}
		if userDeviceTemplate.TemplateSettingURL != templateSettingURL {
			userDeviceTemplate.IsTemplateUpdated = false
		}

		if _, err = userDeviceTemplate.Update(ctx, u.db, boil.Infer()); err != nil {
			return err
		}
	}

	if userDeviceTemplate == nil {
		userDeviceTemplate = &models.DeviceTemplate{
			Vin:                vin,
			TemplateDBCURL:     templateDbcURL,
			TemplatePidURL:     templatePidURL,
			TemplateSettingURL: templateSettingURL,
			IsTemplateUpdated:  true,
		}

		if err = userDeviceTemplate.Insert(ctx, u.db, boil.Infer()); err != nil {
			return err
		}
	}

	return nil
}
