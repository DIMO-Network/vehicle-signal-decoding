package models

type DeviceConfigResponse struct {
	PidURL           string `json:"pidUrl"`
	DeviceSettingURL string `json:"deviceSettingUrl"`
	DbcURL           string `json:"dbcURL,omitempty"`
	Version          string `json:"version"`
}

type UserDeviceAutoPIUnit struct {
	UserDeviceID       string
	DeviceDefinitionID string
	DeviceStyleID      string
}
