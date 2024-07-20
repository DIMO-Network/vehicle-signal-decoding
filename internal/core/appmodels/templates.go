package appmodels

type UserDeviceAutoPIUnit struct {
	UserDeviceID       string
	DeviceDefinitionID string
	DeviceStyleID      string
}

// SettingsData used for the template device power settings mostly
type SettingsData struct {
	SafetyCutOutVoltage                      float64 `json:"safety_cut_out_voltage"`                        //nolint
	SleepTimerEventDrivenPeriodSecs          float64 `json:"sleep_timer_event_driven_period_secs"`          //nolint
	WakeTriggerVoltageLevel                  float64 `json:"wake_trigger_voltage_level"`                    //nolint
	MinVoltageOBDLoggers                     float64 `json:"min_voltage_obd_loggers"`                       //nolint
	LocationFrequencySecs                    float64 `json:"location_frequency_secs"`                       //nolint
	SleepTimerInactivityAfterSleepSecs       float64 `json:"sleep_timer_inactivity_after_sleep_secs"`       //nolint
	SleepTimerInactivityFallbackIntervalSecs float64 `json:"sleep_timer_inactivity_fallback_interval_secs"` //nolint
	SleepTimerEventDrivenIntervalSecs        float64 `json:"sleep_timer_event_driven_interval_secs"`        //nolint
	BatteryCriticalLevelVoltage              float64 `json:"battery_critical_level_voltage"`                //nolint
}
