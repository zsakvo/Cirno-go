package fetch

import (
	"fmt"
	"log"
	"os"

	"../ciweimao"
	"../structure"
)

func DownloadText(bid string, config structure.ConfigStruct) {
	detail := ciweimao.GetDetail(bid, config)
	chapterInfos := ciweimao.GetCatalog(bid, config)
	name := detail.BookName
	author := detail.AuthorName
	fmt.Println(name, author)
	writeText(name+"-"+author, chapterInfos)
}

func writeText(bookName string, chapterInfos []structure.ChapterInfo) {
	bookText := initText(chapterInfos)
	if isExist("./download/" + bookName + ".txt") {
	} else {
		file, err := os.Create("./download/" + bookName + ".txt")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = file.WriteString(bookText)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
	}
	// text := ""
	// text += chapterInfo.ChapterTitle
	// text += "\n\n"
	// text += chapterInfo.TxtContent
	// text += chapterInfo.AuthorSay
	// text += "\n\n\n"
	// f, err := os.Create("./demo.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	_, err = f.Write([]byte(text))
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
	// defer f.Close()
}

func initText(chapterInfos []structure.ChapterInfo) string {
	text := ""
	for _, chapterInfo := range chapterInfos {
		text += chapterInfo.ChapterTitle
		text += "\n\n"
		text += chapterInfo.TxtContent
		text += chapterInfo.AuthorSay
		text += "\n\n\n"
	}
	return text
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
