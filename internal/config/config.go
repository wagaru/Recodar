package config

import "github.com/spf13/viper"

type Config struct {
	MongoURI   string `mapstructure:"MONGO_URI"`
	ServerPort uint16 `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path, configName, configType string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
