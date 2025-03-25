package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

type conf struct {
	Driver        string `mapstructure:"DB_DRIVER"`
	Host          string `mapstructure:"DB_HOST"`
	Port          string `mapstructure:"DB_PORT"`
	User          string `mapstructure:"DB_USER"`
	Password      string `mapstructure:"DB_PASSWORD"`
	Name          string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("goexpert_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	conf := conf{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	conf.TokenAuth = jwtauth.New("HS256", []byte(conf.JWTSecret), nil)

	return &conf, err

}
