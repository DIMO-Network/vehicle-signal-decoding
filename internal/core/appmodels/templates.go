package appmodels

type UserDeviceAutoPIUnit struct {
	UserDeviceID       string
	DeviceDefinitionID string
	DeviceStyleID      string
}

// DeviceConfigResponse response for what templates to use, mobile app dependency: userGetVehicleDecoding.ts
type DeviceConfigResponse struct {
	// PidURL including the version for the template
	PidURL string `json:"pidUrl"`
	// DeviceSettingURL including the version for the settings
	DeviceSettingURL string `json:"deviceSettingUrl"`
	// DbcURL including the version for the dbc file, usually same as pidurl template version
	DbcURL string `json:"dbcUrl,omitempty"`
}

// SettingsData used for the template device power settings mostly
type SettingsData struct {
	SafetyCutOutVoltage             float64 `json:"safety_cut_out_voltage"`               //nolint
	SleepTimerEventDrivenPeriodSecs float64 `json:"sleep_timer_event_driven_period_secs"` //nolint
	WakeTriggerVoltageLevel         float64 `json:"wake_trigger_voltage_level"`           //nolint
	MinVoltageOBDLoggers            float64 `json:"min_voltage_obd_loggers"`              //nolint
	LocationFrequencySecs           float64 `json:"location_frequency_secs"`              //nolint
}
