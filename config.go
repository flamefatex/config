package config

import (
	"strings"

	"github.com/flamefatex/log"
	"github.com/spf13/viper"
)

var defaultConfig *viper.Viper

//	从config.yml初始化配置
func Init(serviceName string) {
	defaultConfig = readViperConfig(serviceName)
	log.L().Infof("Config All Settings From File: %v", defaultConfig.AllSettings())
}

func Config() Provider {
	return defaultConfig
}

func SetTestConfig(v *viper.Viper) {
	defaultConfig = v
}

func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func readViperConfig(serviceName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(strings.ToUpper(serviceName))
	v.AutomaticEnv()

	v.SetConfigName("config")                    // name of config file (without extension)
	v.AddConfigPath("/etc/" + serviceName + "/") // path to look for the config file in
	v.AddConfigPath("$HOME/." + serviceName)     // call multiple times to add many search paths
	v.AddConfigPath(".")                         // optionally look for config in the working directory
	err := v.ReadInConfig()                      // Find and read the config file
	if err != nil {                              // Handle errors reading the config file
		log.L().Fatalf("ReadInConfig err: %s", err)
		return v
	}

	// global defaults
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	return v
}
