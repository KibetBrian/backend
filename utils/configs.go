package utils
import (
	"github.com/spf13/viper"
)

type Config struct{
	DbSource string `mapstructure:"DB_SOURCE"`
	DbDriver string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string)(config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil{
		return
	}
	viper.Unmarshal(&config)
	return
}