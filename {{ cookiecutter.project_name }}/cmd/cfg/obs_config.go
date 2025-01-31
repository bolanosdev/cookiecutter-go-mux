package cfg

type ObsConfig struct {
	JAEGER JaegerConfig `mapstructure:"jaeger"`
}

type JaegerConfig struct {
	DIAL_HOSTNAME string `mapstructure:"dial_hostname"`
}
