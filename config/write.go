package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Write(token, account string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	viper.Set("app.login_token", token)
	viper.Set("app.account", account)
	err1 := viper.WriteConfig()
	if err1 != nil {
		fmt.Printf("writeout config file error: %s\n", err1)
		os.Exit(1)
	}
	fmt.Println("配置写出成功！")
}
