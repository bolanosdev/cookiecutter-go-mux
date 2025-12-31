package config

type ObsConfig struct {
	SENTRY_DSN      string       `mapstructure:"sentry_dsn"`
	JAEGER          JaegerConfig `mapstructure:"jaeger"`
	SENSITIVE_PATHS []string     `mapstructure:"sensitive_paths"`
	IGNORED_PATHS   []string     `mapstructure:"ignored_paths"`
	DEBUGGER_KEY    string       `mapstructure:"debugger_key"`
}

type JaegerConfig struct {
	DIAL_HOSTNAME string `mapstructure:"dial_hostname"`
}
