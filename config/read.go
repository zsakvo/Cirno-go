package config

import (
	"fmt"
	"log"
	"os"

	"../structure"
	"github.com/spf13/viper"
)

func Load() structure.ConfigStruct {
	viper.SetConfigName("config") //把 json 文件换成 yaml 文件，只需要配置文件名 (不带后缀)即可
	viper.AddConfigPath(".")      //添加配置文件所在的路径
	viper.SetConfigType("yaml")   //设置配置文件类型
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	var config structure.ConfigStruct
	err1 := viper.Unmarshal(&config) // 将配置解析到 config 变量
	if err1 != nil {
		log.Fatalf("unable to decode into struct, %v", err1)
	}
	// name := viper.Get("app.user_name")
	// viper.Set("app.login_token", "233")
	// viper.WriteConfig()
	// fmt.Println(config.App.UserName)
	return config
}
