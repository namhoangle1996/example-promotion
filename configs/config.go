package configs

import "github.com/spf13/viper"

type Config struct {
	ENV         string         `json:"env" mapstructure:"env"`
	MaxPoolSize int            `json:"max_pool_size" mapstructure:"max_pool_size"`
	MongoURI    string         `json:"mongo_uri" mapstructure:"mongo_uri"`
	Services    ServicesConfig `json:"services" mapstructure:"services"`
}

type ServicesConfig struct {
	ServiceCore      string `json:"service_core" mapstructure:"service_core"`
	ServiceUser      string `json:"service_user" mapstructure:"service_user"`
	ServicePromotion string `json:"service_promotion" mapstructure:"service_promotion"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigType("json")
	viper.SetConfigName("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	result := &Config{}
	err = viper.Unmarshal(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LoadTestConfig load config for running tests
func LoadTestConfig(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	viper.SetConfigName("config_test.json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	result := &Config{}
	err = viper.Unmarshal(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
