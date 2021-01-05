package fetch

import (
	"fmt"
	"strings"
	"text/template"

	"../ciweimao"
	"../structure"
	"github.com/cheggaaa/pb"
)

var epubBar *pb.ProgressBar

func DownloadEpub(bid string) {
	epubDetail := ciweimao.GetDetail(bid)
	fmt.Println(epubDetail.BookName, "/", epubDetail.AuthorName)
	epubName := epubDetail.BookName
	epubChapters := ciweimao.GetCatalog(bid)
	epubTotalCount := len(epubChapters)
	epubBar = pb.StartNew(epubTotalCount)
	epubContainer := make(map[string]string)
	epubc := make(chan chapterStruct, 409600)
	epubErrc := make(chan structure.ChapterList, 102400)
	epubChaptersArr := splitArray(epubChapters, 16)
	///
	for _, cs := range epubChaptersArr {
		go getChapterEpub(cs, epubc, epubErrc)
	}
	///
	///
	epubBar.Finish()
	fmt.Println("正在写出文件……")
	writeEpub(epubName, epubContainer, epubChapters)
	fmt.Println("下载成功！")
}

func writeEpub(bookName string, epubContainer map[string]string, epubChapters []structure.ChapterList) {
}

func getChapterEpub(chapters []structure.ChapterList, epubc chan chapterStruct, epubErrc chan structure.ChapterList) {
	for _, chapter := range chapters {
		text := ""
		chapterInfo, err := ciweimao.GetContent(chapter.ChapterID)
		if err != nil {
			epubErrc <- chapter
		} else {
			funcMap := template.FuncMap{
				"replace": replace,
			}
			// fmt.Println(i)
			text += chapterInfo.ChapterTitle
			text += "\n\n"
			text += chapterInfo.TxtContent
			text += chapterInfo.AuthorSay
			text += "\n\n\n"
			txtstr := chapterStruct{text, chapter.ChapterID}
			epubc <- txtstr
		}
	}
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}
