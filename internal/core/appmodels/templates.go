package appmodels

type UserDeviceAutoPIUnit struct {
	UserDeviceID       string
	DeviceDefinitionID string
	DeviceStyleID      string
}

// DeviceConfigResponse response for what templates to use
type DeviceConfigResponse struct {
	// PidURL including the version for the template
	PidURL string `json:"pidUrl"`
	// DeviceSettingURL including the version for the settings
	DeviceSettingURL string `json:"deviceSettingUrl"`
	// DbcURL including the version for the dbc file, usually same as pidurl template version
	DbcURL string `json:"dbcURL,omitempty"`
}
