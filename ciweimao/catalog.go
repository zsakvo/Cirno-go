package ciweimao

import (
	"fmt"

	"../snipaste"
	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

func GetCatalog(bid string, config structure.ConfigStruct) []structure.ChapterList {
	var chapterList []structure.ChapterList
	divisions := getDivision(bid, config)
	for _, division := range divisions {
		chapters := getChapters(division.DivisionID, config)
		for _, chapter := range chapters {
			fmt.Println(chapter)
			chapterList = append(chapterList, chapter)
		}
	}
	fmt.Println(len(chapterList))
	return chapterList
}

func getDivision(bid string, config structure.ConfigStruct) []structure.DivisionList {
	param := req.Param{
		"book_id":      bid,
		"account":      config.App.Account,
		"device_token": config.App.DeviceToken,
		"app_version":  config.App.AppVersion,
		"login_token":  config.App.LoginToken,
	}
	r, _ := req.Get("https://app.hbooker.com/book/get_division_list", param)
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.DivisionStruct
	json.Unmarshal(res, &result)
	return result.Data.DivisionList
}

func getChapters(did string, config structure.ConfigStruct) []structure.ChapterList {
	param := req.Param{
		"division_id":  did,
		"account":      config.App.Account,
		"device_token": config.App.DeviceToken,
		"app_version":  config.App.AppVersion,
		"login_token":  config.App.LoginToken,
	}
	r, _ := req.Get("https://app.hbooker.com/chapter/get_updated_chapter_by_division_id", param)
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ChapterStruct
	json.Unmarshal(res, &result)
	return result.Data.ChapterList
}
