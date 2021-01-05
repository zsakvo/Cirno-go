package fetch

import (
	"fmt"
	"log"
	"os"

	"../ciweimao"
	"../structure"
	"github.com/cheggaaa/pb"
	"github.com/mitchellh/go-homedir"
)

var errList []structure.ChapterList
var bar *pb.ProgressBar
var chapterInfos []structure.ChapterInfo

func DownloadText(bid string) {
	detail := ciweimao.GetDetail(bid)
	name := detail.BookName
	chapters := ciweimao.GetCatalog(bid)
	totalCount := len(chapters)
	fmt.Println("开始下载", "《"+name+"》")
	bar = pb.StartNew(totalCount)
	// for _, chapter := range chapters {
	// 	chapterInfo, err := ciweimao.GetContent(chapter.ChapterID)
	// 	if err != nil {
	// 		errList = append(errList, chapter)
	// 	} else {
	// 		chapterInfos = append(chapterInfos, chapterInfo)
	// 		bar.Increment()
	// 	}
	// }
	var txts []string
	go getChapterText(chapters, errList, txts)
	for len(errList) > 0 {
		dealErr()
	}
	writeText(name, chapterInfos)
	bar.Finish()
	fmt.Println("下载成功！")
}

func dealErr() {
	for _, chapter := range errList {
		chapterInfo, err := ciweimao.GetContent(chapter.ChapterID)
		if err != nil {
			errList = append(errList, chapter)
		} else {
			chapterInfos = append(chapterInfos, chapterInfo)
			bar.Increment()
		}
	}
}

func writeText(bookName string, chapterInfos []structure.ChapterInfo) {
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	bookText := initText(chapterInfos)
	fileName := expandedDir + "/Cirno/download/" + bookName + ".txt"
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

func getChapterText(chapters, errs []structure.ChapterList, txts []string) {
	for _, chapter := range chapters {
		text := ""
		chapterInfo, err := ciweimao.GetContent(chapter.ChapterID)
		if err != nil {
			errs = append(errs, chapter)
		} else {
			text += chapterInfo.ChapterTitle
			text += "\n\n"
			text += chapterInfo.TxtContent
			text += chapterInfo.AuthorSay
			text += "\n\n\n"
			txts = append(txts, text)
			bar.Increment()
		}
	}
}
