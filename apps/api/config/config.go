package config

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Port        int    `mapstructure:"PORT"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	HmacKey     string `mapstructure:"HMAC_KEY"`
	HmacSalt    string `mapstructure:"HMAC_SALT"`
}

func BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config

	BindEnvs(cfg)

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return &cfg, err
	}

	return &cfg, nil
}
