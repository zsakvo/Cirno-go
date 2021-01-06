package ciweimao

import (
	"fmt"

	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
	"github.com/zsakvo/Cirno-go/structure"
	"github.com/zsakvo/Cirno-go/util"
)

func GetContent(cid string) (structure.ChapterInfo, error) {
	var err error
	var key string
	key, err = getKey(cid)
	if err != nil {
		fmt.Println(err)
		return structure.ChapterInfo{}, err
	}
	return getDecrypt(key, cid)
}

func getKey(chapterId string) (string, error) {
	var err error
	var res []byte
	paras := req.Param{
		"chapter_id": chapterId,
	}
	res, err = util.Get("/chapter/get_chapter_cmd", paras)
	if err != nil {
		return "", err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.KeyStruct
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}
	return result.Data.Command, nil
}

func getDecrypt(key, chapterId string) (structure.ChapterInfo, error) {
	var err error
	var res []byte
	paras := req.Param{
		"chapter_id":      chapterId,
		"chapter_command": key,
	}
	res, err = util.Get("/chapter/get_cpt_ifm", paras)
	if err != nil {
		fmt.Println(err)
		return structure.ChapterInfo{}, err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ContentStruct
	err = json.Unmarshal(res, &result)
	if err != nil {
		return structure.ChapterInfo{}, err
	}
	var bytes []byte
	bytes, err = util.Decode(result.Data.ChapterInfo.TxtContent, key)
	if err != nil {
		return structure.ChapterInfo{}, err
	}
	result.Data.ChapterInfo.TxtContent = string(bytes)
	return result.Data.ChapterInfo, nil
}
