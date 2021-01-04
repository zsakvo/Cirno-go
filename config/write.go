package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func Write(name, password, token, account string) {
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	filePath := expandedDir + "/Cirno/config.yaml"
	if isExist(filePath) {
		os.Remove(filePath)
	}
	file, _ := os.Create(filePath)
	defer file.Close()
	viper.SetConfigName("config")
	viper.AddConfigPath(expandedDir + "/Cirno")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	viper.Set("app.login_token", token)
	viper.Set("app.account", account)
	viper.Set("app.user_name", name)
	viper.Set("app.password", password)
	err1 := viper.WriteConfig()
	if err1 != nil {
		fmt.Printf("writeout config file error: %s\n", err1)
		os.Exit(1)
	}
	fmt.Println("登陆成功，可以开始使用了")
	os.Exit(0)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
