package conf

import (
	"github.com/spf13/viper"
	"sign/tool/http"
	"sign/tool/mysql"
	"sign/tool/redis"
)

var Conf *Config

type App struct {
	ServerName string `toml:"serverName"`
	Debug      bool   `toml:"debug"`
	JwtSecret  string `toml:"jwtSecret"`
}

type Grpc struct {
	Port string `toml:"port"`
}

type Config struct {
	App   *App
	Mysql *mysql.Instance
	Http  *http.Instance
	Grpc  *Grpc
	Redis *redis.Instance
}

func Init() error {
	v := viper.New()
	v.AddConfigPath("conf")
	v.SetConfigType("toml")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&Conf)
	return err
}
