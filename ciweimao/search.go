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

func Search(bookName string, page int, config structure.ConfigStruct) {
	param := req.Param{
		"count":          10,
		"page":           page,
		"category_index": 0,
		"key":            bookName,
		"account":        config.App.Account,
		"device_token":   config.App.DeviceToken,
		"app_version":    config.App.AppVersion,
		"login_token":    config.App.LoginToken,
	}
	r, _ := req.Get("https://app.hbooker.com/bookcity/get_filter_search_book_list", param)
	res := util.Decode(r.String(), snipaste.InitEncryptKey)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.SearchStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	bookList := result.Data.BookList
	for i, book := range bookList {
		fmt.Println(i, "-", book.BookName, "-", book.BookID)
	}
}
