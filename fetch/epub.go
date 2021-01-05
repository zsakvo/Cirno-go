package fetch

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"../ciweimao"
	"../structure"
	"../util"
	"github.com/cheggaaa/pb"
	"github.com/imroc/req"
	"github.com/mitchellh/go-homedir"
)

var epubBar *pb.ProgressBar
var tmpPath string
var oebpsPath string

func initTemp(name, author, cover string, chapters []structure.ChapterList) {
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	tmpPath = filepath.Join(expandedDir, "Cirno", "download", "tmp", name)
	oebpsPath = filepath.Join(tmpPath, "OEBPS")
	coverElement := coverHeader + "\n" + "<img src=\"cover.jpg\" alt=\"" + name + "\" />" + coverFooter
	writeOut(mimetype, tmpPath, "mimetype")
	writeOut(container, filepath.Join(tmpPath, "META-INF"), "container.xml")
	writeOut(coverElement, oebpsPath, "cover.html")
	writeOut(getCover(cover), oebpsPath, "cover.jpg")
	writeOut(css, oebpsPath, "style.css")
	writeOut(genBookToc(chapters), oebpsPath, "book-toc.html")
	writeOut(genContentOpf(name, author, chapters), oebpsPath, "content.opf")
	writeOut(genTocNcx(name, author, chapters), oebpsPath, "toc.ncx")
}

func DownloadEpub(bid string) {
	epubDetail := ciweimao.GetDetail(bid)
	fmt.Println(epubDetail.BookName, "/", epubDetail.AuthorName)
	epubChapters := ciweimao.GetCatalog(bid)
	fmt.Println("正在获取数据……")
	initTemp(epubDetail.BookName, epubDetail.AuthorName, epubDetail.Cover, epubChapters)
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
	for {
		select {
		case epub := <-epubc:
			epubContainer[epub.Cid] = epub.Text
			epubBar.Increment()
		case e := <-epubErrc:
			go getChapterEpub([]structure.ChapterList{e}, epubc, epubErrc)
		}
		if len(epubContainer) == len(epubChapters) {
			close(epubc)
			close(epubErrc)
			break
		}
	}
	///
	///
	epubBar.Finish()
	fmt.Println("正在写出文件……")
	// writeEpub(epubDetail.BookName, epubContainer, epubChapters)
	epubPath := filepath.Join(tmpPath, "..", "..", epubDetail.BookName+".epub")
	util.CompressEpub(tmpPath, epubPath)
	// fmt.Println(tmpPath)
	fmt.Println("下载成功！")
}

func writeEpub(bookName string, epubContainer map[string]string, epubChapters []structure.ChapterList) {
}

func getChapterEpub(chapters []structure.ChapterList, epubc chan chapterStruct, epubErrc chan structure.ChapterList) {
	for _, chapter := range chapters {
		// text := ""
		chapterInfo, err := ciweimao.GetContent(chapter.ChapterID)
		if err != nil {
			epubErrc <- chapter
		} else {
			content := chapterInfo.TxtContent
			titleElement := "<h2 id=\"title\" class=\"titlel2std\">" + chapterInfo.ChapterTitle + "</h2>"
			content = strings.Replace(content, "　　", "<p class=\"a\">　　", -1)
			content = strings.Replace(content, "\n", "</p>", -1)
			content = strings.Replace(content, "</p></p>", "</p>", -1)
			content = strings.Replace(content, "&", "&#38;", -1)
			// fmt.Println(i)
			// text += chapterInfo.ChapterTitle
			// text += "\n\n"
			// text += chapterInfo.TxtContent
			// text += chapterInfo.AuthorSay
			// text += "\n\n\n"
			txtstr := chapterStruct{contentHeader + "\n" + titleElement + "\n" + content + "\n" + contentFooter, chapter.ChapterID}
			fileName := "chapter" + chapter.ChapterID + ".html"
			writeOut(contentHeader+"\n"+titleElement+"\n"+content+"\n"+contentFooter, oebpsPath, fileName)
			epubc <- txtstr
		}
	}
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func genBookToc(chapters []structure.ChapterList) string {
	var str string
	for _, chapter := range chapters {
		str += "<dt class=\"tocl2\"><a href=\"chapter" + chapter.ChapterID + ".html\">" + chapter.ChapterTitle + "</a></dt>"

	}
	return bookTocHeader + str + bookTocFooter
}

func genContentOpf(name, author string, chapters []structure.ChapterList) string {
	var manifestStr string
	var spineStr string
	for i, chapter := range chapters {
		manifestStr += "<item id=\"chapter" + strconv.Itoa(i) + "\" href=\"chapter" + chapter.ChapterID + ".html\" media-type=\"application/xhtml+xml\"/>"
		spineStr += "<itemref idref=\"chapter" + strconv.Itoa(i) + "\" linear=\"yes\"/>"

	}
	contentOpfHeader = strings.Replace(contentOpfHeader, "bookTitle", name, 1)
	contentOpfHeader = strings.Replace(contentOpfHeader, "bookAuthor", author, 1)
	return contentOpfHeader + contentOpfManifestHeader + manifestStr + contentOpfManifestFooter + contentOpfNcxtocHeader + spineStr + contentOpfNcxtocFooter + contentOpfFooter
}

func genTocNcx(name, author string, chapters []structure.ChapterList) string {
	docTitle := "<docTitle><text>" + name + "</text></docTitle>"
	docAuthor := "<docAuthor><text>" + author + "</text></docAuthor>"
	navMap := "<navMap> <navPoint id=\"cover\" playOrder=\"1\"> <navLabel><text>封面</text></navLabel> <content src=\"cover.html\"/> </navPoint> <navPoint id=\"htmltoc\" playOrder=\"2\"> <navLabel><text>目录</text></navLabel> <content src=\"book-toc.html\"/> </navPoint>\""
	var str string
	for i, chapter := range chapters {
		str += "<navPoint id=\"chapter" + strconv.Itoa(i) + "\" playOrder=\"" + strconv.Itoa(3+i) + "\"> <navLabel><text>" + chapter.ChapterTitle + "</text></navLabel> <content src=\"chapter" + chapter.ChapterID + ".html\"/> </navPoint>"

	}
	return tocNcxHeader + docTitle + docAuthor + navMap + str + tocNcxFooter
}

func getCover(url string) string {
	client := req.New()
	client.SetTimeout(20 * time.Second)
	r, _ := client.Get(url)
	return r.String()
}
