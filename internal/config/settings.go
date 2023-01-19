package config

import (
	"github.com/DIMO-Network/shared/db"
)

type Settings struct {
	Environment      string      `yaml:"ENVIRONMENT"`
	Port             string      `yaml:"PORT"`
	LogLevel         string      `yaml:"LOG_LEVEL"`
	DB               db.Settings `yaml:"DB"`
	ServiceName      string      `yaml:"SERVICE_NAME"`
	ServiceVersion   string      `yaml:"SERVICE_VERSION"`
	GRPCPort         string      `yaml:"GRPC_PORT"`
	TraceMonitorView string      `yaml:"TRACE_MONITOR_VIEW"`
}
