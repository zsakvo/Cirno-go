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

var bar *pb.ProgressBar

type textStruct struct {
	Text string
	Cid  string
}

func DownloadText(bid string) {
	detail := ciweimao.GetDetail(bid)
	name := detail.BookName
	chapters := ciweimao.GetCatalog(bid)
	totalCount := len(chapters)
	fmt.Println("开始下载", "《"+name+"》")
	bar = pb.StartNew(totalCount)
	txtContainer := make(map[string]string)
	errs := []structure.ChapterList{}
	txt := make(chan textStruct, 409600)
	err := make(chan structure.ChapterList, 102400)
	chaptersArr := splitArray(chapters, 4)
	for _, cs := range chaptersArr {
		go getChapterText(cs, txt, err)
	}
	for {
		select {
		case t := <-txt:
			txtContainer[t.Cid] = t.Text
			bar.Increment()
		case e := <-err:
			go getChapterText([]structure.ChapterList{e}, txt, err)
		}
		if len(txtContainer)+len(errs) == len(chapters) {
			close(txt)
			close(err)
			break
		}
	}
	writeText(name, txtContainer, chapters)
	bar.Finish()
	fmt.Println("下载成功！")
}

func writeText(bookName string, txtContainer map[string]string, chapters []structure.ChapterList) {
	bookText := ""
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	for _, chapter := range chapters {
		bookText += txtContainer[chapter.ChapterID]
	}
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
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func getChapterText(chapters []structure.ChapterList, txt chan textStruct, errc chan structure.ChapterList) {
	for _, chapter := range chapters {
		text := ""
		chapterInfo, err := ciweimao.GetContent(chapter.ChapterID)
		if err != nil {
			// fmt.Println(err)
			errc <- chapter
			// errs = append(errs, chapter)
		} else {
			// fmt.Println(i)
			text += chapterInfo.ChapterTitle
			text += "\n\n"
			text += chapterInfo.TxtContent
			text += chapterInfo.AuthorSay
			text += "\n\n\n"
			txtstr := textStruct{text, chapter.ChapterID}
			txt <- txtstr
		}
	}
}

func splitArray(arr []structure.ChapterList, num int64) [][]structure.ChapterList {
	var segmens = make([][]structure.ChapterList, 0)
	max := int64(len(arr))
	if max < num {
		segmens = append(segmens, arr)
		return segmens
	}
	quantity := max / num
	end := int64(0)
	for i := int64(1); i <= num; i++ {
		qu := i * quantity
		if i != num {
			segmens = append(segmens, arr[i-1+end:qu])
		} else {
			segmens = append(segmens, arr[i-1+end:])
		}
		end = qu - i
	}
	return segmens
}
