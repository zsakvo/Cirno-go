package fetch

import (
	"fmt"
	"log"
	"os"

	"../ciweimao"
	"../structure"
	"github.com/cheggaaa/pb"
)

func DownloadText(bid string, config structure.ConfigStruct) {
	var chapterInfos []structure.ChapterInfo
	detail := ciweimao.GetDetail(bid)
	name := detail.BookName
	chapters := ciweimao.GetCatalog(bid)
	totalCount := len(chapters)
	fmt.Println("开始下载", name)
	bar := pb.StartNew(totalCount)
	for _, chapter := range chapters {
		chapterInfos = append(chapterInfos, ciweimao.GetContent(chapter.ChapterID))
		bar.Increment()
	}
	writeText(name, chapterInfos)
	bar.Finish()
	fmt.Println("下载成功！")
}

func writeText(bookName string, chapterInfos []structure.ChapterInfo) {
	bookText := initText(chapterInfos)
	fileName := "./download/" + bookName + ".txt"
	if isExist(fileName) {
		os.Remove(fileName)
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = file.WriteString(bookText)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
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
