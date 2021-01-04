package ciweimao

import (
	"log"

	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

func GetDetail(bid string) structure.BookInfo {
	paras := req.Param{
		"book_id": bid,
	}
	res := util.Get("/book/get_info_by_id", paras, nil)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.DetailStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Data.BookInfo
}
