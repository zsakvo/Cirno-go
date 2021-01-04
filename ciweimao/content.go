package ciweimao

import (
	"log"

	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

var chapterId string

func GetContent(cid string) structure.ChapterInfo {
	chapterId = cid
	key := getKey()
	return getDecrypt(key)
}

func getKey() string {
	paras := req.Param{
		"chapter_id": chapterId,
	}
	res := util.Get("/chapter/get_chapter_cmd", paras, nil)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.KeyStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Data.Command
}

func getDecrypt(key string) structure.ChapterInfo {
	paras := req.Param{
		"chapter_id":      chapterId,
		"chapter_command": key,
	}
	res := util.Get("/chapter/get_cpt_ifm", paras, nil)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ContentStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	result.Data.ChapterInfo.TxtContent = string(util.Decode(result.Data.ChapterInfo.TxtContent, key))
	return result.Data.ChapterInfo
}
