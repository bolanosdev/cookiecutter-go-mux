package config

import "time"

type PasetoConfig struct {
	TOKEN_SYMETRIC_KEY    string        `mapstructure:"token_symetric_key"`
	TOKEN_RENEW_DURATION  time.Duration `mapstructure:"token_renew_duration"`
	TOKEN_EXPIRE_DURATION time.Duration `mapstructure:"token_expire_duration"`
}
