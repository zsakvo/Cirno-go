package ciweimao

import (
	"fmt"
	"os"
	"strconv"

	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
	"github.com/zsakvo/Cirno-go/structure"
	"github.com/zsakvo/Cirno-go/util"
)

func Search(bookName string, page int, bookType string) {
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
		fmt.Println(i, "-", book.BookName)
	}
	enterCmd(bookName, page, bookList, bookType)
}

func enterCmd(bookName string, page int, bookList []structure.BookList, bookType string) {
	var input string
	fmt.Printf("enter the number to download, n to next page, and p to previous page: ")
	fmt.Scanln(&input)
	fmt.Println("")
	switch input {
	case "n":
		if len(bookList) < 10 {
			fmt.Println("already the last page.")
			enterCmd(bookName, page, bookList, bookType)
		}
		fmt.Printf("\x1bc")
		fmt.Printf("search result:")
		fmt.Println("")
		Search(bookName, page+1, bookType)
	case "p":
		if page-1 < 0 {
			fmt.Println("already the first page.")
			enterCmd(bookName, page, bookList, bookType)
		}
		fmt.Printf("\x1bc")
		fmt.Printf("search result:")
		fmt.Println("")
		Search(bookName, page-1, bookType)
	default:
		bidNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid num code")
			enterCmd(bookName, page, bookList, bookType)
		}
		switch bookType {
		case "txt":
			fmt.Printf("\x1bc")
			fmt.Printf("downloading…")
			fmt.Println("")
			DownloadText(bookList[bidNum].BookID)
		case "epub":
			fmt.Printf("\x1bc")
			fmt.Printf("downloading…")
			fmt.Println("")
			DownloadEpub(bookList[bidNum].BookID)
		default:
			fmt.Println("invlid type.")
			os.Exit(0)
		}
	}
}
