package util

import "github.com/spf13/viper"

type ConfigStruct struct {
	HueBridgeIp        string `mapstructure:"HUE_BRIDGE_IP"`
	HueUsername        string `mapstructure:"HUE_USERNAME"`
	HueLightNumber     string `mapstructure:"HUE_LIGHT_NUMBER"`
	ColorCycleInterval int    `mapstructure:"COLOR_CYCLE_INTERVAL"`
	ColorBrightness    int    `mapstructure:"COLOR_BRI"`
	ColorSaturation    int    `mapstructure:"COLOR_SAT"`
}

var Config ConfigStruct

func LoadConfig(path string) (config ConfigStruct, err error) {
	// config file
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// environment variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	Config = config
	return
}
