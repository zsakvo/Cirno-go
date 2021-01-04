package ciweimao

import (
	"log"

	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
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
	paras := req.Param{
		"book_id": bid,
	}
	res := util.Get("/book/get_division_list", paras, nil)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.DivisionStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Data.DivisionList
}

func getChapters(did string) []structure.ChapterList {
	paras := req.Param{
		"division_id": did,
	}
	res := util.Get("/chapter/get_updated_chapter_by_division_id", paras, nil)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ChapterStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Data.ChapterList
}
