package ciweimao

import (
	"fmt"

	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

func Search(bookName string, page int) {
	var err error
	var res []byte
	paras := req.Param{
		"count":          10,
		"page":           page,
		"category_index": 0,
		"key":            bookName,
	}
	res, err = util.Get("/bookcity/get_filter_search_book_list", paras)
	util.PanicErr(err)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.SearchStruct
	err = json.Unmarshal(res, &result)
	util.PanicErr(err)
	bookList := result.Data.BookList
	for i, book := range bookList {
		fmt.Println(i, "-", book.BookName, "-", book.BookID)
	}
}
