package config

import "time"

type AppConfig struct {
	Version string `yaml:"version"`
	Ping_interval time.Duration `yaml:"ping_interval"`
	Log_level LogLevel `yaml:"log_level"`
	Max_concurrent_pings uint8 `yaml:"max_concurrent_pings"`
}

type ResourcesList struct {
	Name string `yaml:"name"`
	Address string `yaml:"address"`
	Enabled bool `yaml:"enabled"`
	Timeout int `yaml:"timeout"`
	Retries int `yaml:"retries"`
}

type AlertsConfig struct {
	Enabled bool `yaml:"enabled"`
}

type LogsConfig struct {
	Enabled bool `yaml:"enabled"`
}

type GrafanaConfig struct {
	DashboardEnabled bool `yaml:"dashboard_enabled"`
	AlertsEnabled bool `yaml:"alerts_enabled"`
	Thresholds GrafanaThresholdConfig `yaml:"thresholds"`
}

type GrafanaThresholdConfig struct {
	LatencyWarning int `yaml:"latency_warning"`
	LatencyCritical int `yaml:"latency_critical"`
	PacketLossWarning int `yaml:"packet_loss_warning"`
	PacketLossCritical int `yaml:"packet_loss_critical"`
}

type LogLevel string

const (
	Trace LogLevel = "trace"
	Debug LogLevel = "debug"
	Info LogLevel = "info"
	Warn LogLevel = "warn"
	Error LogLevel = "error"
)