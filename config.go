package config

import (
	"fmt"
	"strings"

	"github.com/flamefatex/log"
	"github.com/spf13/viper"
)

var defaultConfig *viper.Viper = viper.New()

//	从config.yml初始化配置
func Init(serviceName string) {
	defaultConfig = readViperConfig(serviceName)
	if Config().GetBool("config.enable_log") {
		log.Infof("Config All Settings From File: %v", defaultConfig.AllSettings())
	}
}

func Config() Provider {
	return defaultConfig
}

func SetTestConfig(v *viper.Viper) {
	defaultConfig = v
}

func readViperConfig(serviceName string) *viper.Viper {
	v := viper.New()

	// env
	envServiceName := strings.ToUpper(strings.ReplaceAll(serviceName, "-", "_"))
	v.SetEnvPrefix(envServiceName)
	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// config file
	v.SetConfigName("config")                    // name of config file (without extension)
	v.AddConfigPath("/etc/" + serviceName + "/") // path to look for the config file in
	v.AddConfigPath("$HOME/." + serviceName)     // call multiple times to add many search paths
	v.AddConfigPath(".")                         // optionally look for config in the working directory
	err := v.ReadInConfig()                      // Find and read the config file
	if err != nil {                              // Handle errors reading the config file
		err = fmt.Errorf("ReadInConfig err: %w", err)
		panic(err)
	}

	// global defaults
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	return v
}
