package util

import "github.com/spf13/viper"

// stores all configurations of the application
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // path to look for the config file in
	viper.SetConfigName("app") // name of config file without extension
	viper.SetConfigType("env") // extension of config file
	viper.AutomaticEnv()       // read in environment variables that match

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
