package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	v *viper.Viper
)

func InitConfig() error {
	v = viper.New()
	v.SetConfigName("app")
	v.AddConfigPath("./conf")
	v.SetConfigType("toml")

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetConfig(key string) interface{} {
	return v.Get(key)
}

func GetString(key string) string {
	return v.GetString(key)
}

func GetBool(key string) bool {
	return v.GetBool(key)
}

func GetInt(key string) int {
	return v.GetInt(key)
}

func GetInt32(key string) int32 {
	return v.GetInt32(key)
}

func GetInt64(key string) int64 {
	return v.GetInt64(key)
}

func GetUint(key string) uint {
	return v.GetUint(key)
}

func GetUint32(key string) uint32 {
	return v.GetUint32(key)
}

func GetUint64(key string) uint64 {
	return v.GetUint64(key)
}

func GetFloat64(key string) float64 {
	return v.GetFloat64(key)
}

func GetTime(key string) time.Time {
	return v.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return v.GetDuration(key)
}

func GetStringSlice(key string) []string {
	return v.GetStringSlice(key)
}

func GetStringMap(key string) map[string]interface{} {
	return v.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return v.GetStringMapString(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return v.GetStringMapStringSlice(key)
}

func IsSet(key string) bool {
	return v.IsSet(key)
}
