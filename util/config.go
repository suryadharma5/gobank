package util

import "github.com/spf13/viper"

// Stores all config of the application
// values are read by viper
// map all the env variables in here
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// Load config reads configuration from file or env variables
func LoadConfig(path string) (conf Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") // sesuaikan dengan nama .env file kita
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	return
}
