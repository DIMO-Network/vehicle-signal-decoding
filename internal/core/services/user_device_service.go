package services

import (
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/models"
)

type UserDeviceService interface {
	GetUserDeviceServiceByAutoPIUnitID(id string) (*models.UserDeviceAutoPIUnit, error)
}

type userDeviceService struct {
}

func NewUserDeviceService() UserDeviceService {
	return &userDeviceService{}
}

func (a *userDeviceService) GetUserDeviceServiceByAutoPIUnitID(id string) (*models.UserDeviceAutoPIUnit, error) {
	return &models.UserDeviceAutoPIUnit{UserDeviceID: "", DeviceDefinitionID: "", DeviceStyleID: ""}, nil
}
