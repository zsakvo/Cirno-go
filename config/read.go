package config

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/zsakvo/Cirno-go/structure"
)

var Config structure.ConfigStruct

func InitConfig(hasConfig bool) {
	if hasConfig {
		dir, _ := homedir.Dir()
		expandedDir, _ := homedir.Expand(dir)
		viper.SetConfigName("config")
		viper.AddConfigPath(expandedDir + "/Cirno")
		viper.SetConfigType("yaml")
		viper.SetDefault("app.app_version", "2.7.017")
		viper.SetDefault("app.device_token", "ciweimao_client")
		viper.SetDefault("app.user_agent", "Android com.kuangxiangciweimao.novel")
		viper.SetDefault("app.default_key", "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn")
		viper.SetDefault("app.host_url", "https://app.hbooker.com")
		viper.SetDefault("extra.coroutines", 3)
		viper.SetDefault("extra.cpic", false)
		viper.SetDefault("extra.cache_no_paid", false)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("config file error: %s\n", err)
			os.Exit(1)
		}
		err = viper.Unmarshal(&Config) // 将配置解析到 config 变量
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	} else {
		Config = structure.ConfigStruct{
			App: structure.App{
				AppVersion:  "2.7.017",
				DeviceToken: "ciweimao_client",
				UserAgent:   "Android com.kuangxiangciweimao.novel",
				DefaultKey:  "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn",
				HostUrl:     "http://app.hbooker.com",
			},
		}
	}
}
