package config

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/zsakvo/Cirno-go/structure"
)

func Load() structure.ConfigStruct {
	var config structure.ConfigStruct
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	viper.SetConfigName("config")
	viper.AddConfigPath(expandedDir + "/Cirno")
	viper.SetConfigType("yaml")
	viper.SetDefault("app.app_version", "2.7.017")
	viper.SetDefault("app.device_token", "ciweimao_client")
	viper.SetDefault("app.user_agent", "Android com.kuangxiangciweimao.novel")
	viper.SetDefault("app.default_key", "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	err = viper.Unmarshal(&config) // 将配置解析到 config 变量
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return config
	// name := viper.Get("app.user_name")
	// viper.Set("app.login_token", "233")
	// viper.WriteConfig()
	// fmt.Println(config.App.UserName)
}

func GetTmp() structure.ConfigStruct {
	return structure.ConfigStruct{
		App: structure.App{
			AppVersion:  "2.7.017",
			DeviceToken: "ciweimao_client",
			UserAgent:   "Android com.kuangxiangciweimao.novel",
			DefaultKey:  "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn",
		},
	}
}

// func isExist(path string) bool {
// 	_, err := os.Stat(path)
// 	if err != nil {
// 		return os.IsExist(err)
// 	}
// 	return true
// }
