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

func GetDetail(bid string, c structure.ConfigStruct) structure.BookInfo {
	fmt.Println((bid))
	param := req.Param{
		"book_id":      bid,
		"account":      c.App.Account,
		"device_token": c.App.DeviceToken,
		"app_version":  c.App.AppVersion,
		"login_token":  c.App.LoginToken,
	}
	r, _ := req.Get("https://app.hbooker.com/book/get_info_by_id", param, req.Header{"User-Agent": c.App.UserAgent})
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.DetailStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Data.BookInfo
}
