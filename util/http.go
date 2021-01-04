package util

import (
	"time"

	"../config"
	"../structure"
	"github.com/imroc/req"
)

var cfg structure.ConfigStruct

func InitReq() {
	cfg = config.Load()
}

func Get(url string, paras req.Param) []byte {
	var param req.Param
	if url == "/signup/login" {
		param = req.Param{
			"app_version":  cfg.App.AppVersion,
			"device_token": cfg.App.DeviceToken,
		}
	} else {
		param = req.Param{
			"account":      cfg.App.Account,
			"device_token": cfg.App.DeviceToken,
			"app_version":  cfg.App.AppVersion,
			"login_token":  cfg.App.LoginToken,
		}
	}
	for k, v := range paras {
		param[k] = v
	}
	client := req.New()
	client.SetTimeout(20 * time.Second)
	r, _ := client.Get("https://app.hbooker.com"+url, param, req.Header{"User-Agent": cfg.App.UserAgent})
	res := Decode(r.String(), cfg.App.DefaultKey)
	return res
}
