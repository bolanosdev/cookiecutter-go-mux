package cfg

import "time"

type PasetoConfig struct {
	TOKEN_SYMETRIC_KEY string        `mapstructure:"token_symetric_key"`
	TOKEN_DURATION     time.Duration `mapstructure:"token_duration"`
}
