package helpers

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config yaml file: %w", err))
	}

	viper.SetConfigName("mysql")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error mysql yaml file: %w", err))
	}
}
