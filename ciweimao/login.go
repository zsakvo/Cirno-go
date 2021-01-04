package ciweimao

import (
	"log"

	"../config"
	"../snipaste"
	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

func Login(c structure.ConfigStruct) {
	param := req.Param{
		"app_version":  c.App.AppVersion,
		"device_token": c.App.DeviceToken,
		"login_name":   c.App.UserName,
		"passwd":       c.App.Password,
	}
	r, _ := req.Get("https://app.hbooker.com/signup/login", param, req.Header{"User-Agent": c.App.UserAgent})
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.LoginStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	} else {
		config.Write(result.Data.LoginToken, result.Data.ReaderInfo.Account)
	}
}
