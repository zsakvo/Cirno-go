package ciweimao

import (
	"fmt"
	"log"

	"../snipaste"
	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

var appConfig structure.ConfigStruct
var chapterId string

func GetContent(cid string, cfg structure.ConfigStruct) structure.ChapterInfo {
	chapterId = cid
	appConfig = cfg
	key := getKey()
	return getDecrypt(key)
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
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Data.Command
}

func getDecrypt(key string) structure.ChapterInfo {
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
	if chapterId == "102839480" {
		fmt.Println(string(res))
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ContentStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	result.Data.ChapterInfo.TxtContent = string(util.Decode(result.Data.ChapterInfo.TxtContent, key))
	return result.Data.ChapterInfo
}
