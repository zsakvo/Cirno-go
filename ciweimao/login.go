package ciweimao

import (
	"fmt"

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
	r, _ := req.Get("https://app.hbooker.com/signup/login", param)
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.LoginStruct
	json.Unmarshal(res, &result)
	fmt.Println(result.Data.LoginToken)
	fmt.Println(result.Data.ReaderInfo.Account)
	config.Write(result.Data.LoginToken, result.Data.ReaderInfo.Account)
}
