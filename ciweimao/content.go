package ciweimao

import (
	"../snipaste"
	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

var appConfig structure.ConfigStruct
var chapterId string

func GetContent(cid string, cfg structure.ConfigStruct) {
	chapterId = cid
	appConfig = cfg
	key := getKey()
	eTxt := getEncryptText(key)
	dTxt := string(util.Decode(eTxt, key))
	println(dTxt)
}

func getKey() string {
	param := req.Param{
		"chapter_id":   chapterId,
		"account":      appConfig.App.Account,
		"device_token": appConfig.App.DeviceToken,
		"app_version":  appConfig.App.AppVersion,
		"login_token":  appConfig.App.LoginToken,
	}
	r, _ := req.Get("https://app.hbooker.com/chapter/get_chapter_cmd", param)
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.KeyStruct
	json.Unmarshal(res, &result)
	println(result.Data.Command)
	return result.Data.Command
}

func getEncryptText(key string) string {
	param := req.Param{
		"chapter_id":      chapterId,
		"chapter_command": key,
		"account":         appConfig.App.Account,
		"device_token":    appConfig.App.DeviceToken,
		"app_version":     appConfig.App.AppVersion,
		"login_token":     appConfig.App.LoginToken,
	}
	r, _ := req.Get("https://app.hbooker.com/chapter/get_cpt_ifm", param)
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ContentStruct
	json.Unmarshal(res, &result)
	// print(result.Data.ChapterInfo.TxtContent)
	return result.Data.ChapterInfo.TxtContent
}
