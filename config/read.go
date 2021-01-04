package config

import (
	"fmt"
	"log"
	"os"

	"../structure"
	"github.com/spf13/viper"
)

func Load() structure.ConfigStruct {
	var config structure.ConfigStruct
	if isExist("./config.yaml") {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("config file error: %s\n", err)
			os.Exit(1)
		}
		err = viper.Unmarshal(&config) // 将配置解析到 config 变量
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	} else {
		println("参数文件不存在，已为您初始化，请填写账户密码后登入")
		file, err := os.Create("./config.yaml")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetDefault("app.user_name", "")
		viper.SetDefault("app.password", "")
		viper.SetDefault("app.app_version", "2.7.017")
		viper.SetDefault("app.device_token", "ciweimao_client")
		viper.SetDefault("app.user_agent", "Android com.kuangxiangciweimao.novel")
		viper.SetDefault("app.default_key", "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn")
		err = viper.WriteConfig()
		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(1)
	}
	return config
	// name := viper.Get("app.user_name")
	// viper.Set("app.login_token", "233")
	// viper.WriteConfig()
	// fmt.Println(config.App.UserName)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
