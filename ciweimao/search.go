package ciweimao

import (
	"fmt"
	"log"

	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

func Search(bookName string, page int) {
	paras := req.Param{
		"count":          10,
		"page":           page,
		"category_index": 0,
		"key":            bookName,
	}
	res := util.Get("/bookcity/get_filter_search_book_list", paras)
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
