package ciweimao

import (
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
	"github.com/zsakvo/Cirno-go/structure"
	"github.com/zsakvo/Cirno-go/util"
)

func GetCatalog(bid string) []structure.ChapterList {
	var chapterList []structure.ChapterList
	divisions := getDivision(bid)
	for _, division := range divisions {
		chapters := getChapters(division.DivisionID)
		for _, chapter := range chapters {
			if chapter.IsValid == "1" {
				chapterList = append(chapterList, chapter)
			}
		}
	}
	return chapterList
}

func getDivision(bid string) []structure.DivisionList {
	var err error
	var res []byte
	paras := req.Param{
		"book_id": bid,
	}
	res, err = util.Get("/book/get_division_list", paras)
	util.PanicErr(err)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.DivisionStruct
	err = json.Unmarshal(res, &result)
	util.PanicErr(err)
	return result.Data.DivisionList
}

func getChapters(did string) []structure.ChapterList {
	var err error
	var res []byte
	paras := req.Param{
		"division_id": did,
	}
	res, err = util.Get("/chapter/get_updated_chapter_by_division_id", paras)
	util.PanicErr(err)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ChapterStruct
	err = json.Unmarshal(res, &result)
	util.PanicErr(err)
	return result.Data.ChapterList
}
