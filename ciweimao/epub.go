package ciweimao

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/imroc/req"
	"github.com/mitchellh/go-homedir"
	"github.com/zsakvo/Cirno-go/config"
	"github.com/zsakvo/Cirno-go/structure"
	"github.com/zsakvo/Cirno-go/util"
)

var epubBar *pb.ProgressBar
var tmpPath string
var bookPath string
var oebpsPath string
var cpic bool
var cacheNoPaid bool

func fixImgTag(str string) string {
	if strings.Contains(str, "<img src") && !strings.Contains(str, "</img?") {
		str += "</img>"
	}
	return str
}

func getImg(str string) string {
	if strings.Contains(str, "<img src") {
		var res = regexp.MustCompile(`<img src="(\S*)" alt`)
		var src = res.FindAllStringSubmatch(str, -1)[0][1]
		arrs := strings.Split(src, "/")
		picName := arrs[len(arrs)-1]
		err := writeOut(getCover(src), oebpsPath+"/covers", picName)
		if err != nil {
			fmt.Println(picName, "下载失败")
		}
		return `<img src="./covers/` + picName + `"/>`
	} else {
		return str
	}
}

func initTemp(name, author, cover string, chapters []structure.ChapterList) {
	var err error
	err = util.RemoveContents(tmpPath)
	util.PanicErr(err)
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	cleanDestBook(expandedDir+"/Cirno/download/", name, "epub")
	tmpPath = expandedDir + "/Cirno/download/tmp/"
	bookPath = tmpPath + name
	oebpsPath = bookPath + "/EPUB"
	err = writeOut(mimetype, bookPath, "mimetype")
	util.PanicErr(err)
	err = writeOut(containerXml, bookPath+"/META-INF", "container.xml")
	util.PanicErr(err)
	err = writeOut(getCover(cover), oebpsPath+"/covers", "cover.jpg")
	util.PanicErr(err)
	err = writeOut(epubCss, oebpsPath+"/css", "epub.css")
	util.PanicErr(err)
	err = writeOut(synthCss, oebpsPath+"/css", "synth.css")
	genBookToc(name, chapters)
	util.PanicErr(err)
	genContentOpf(name, author, chapters)
}

func DownloadEpub(bid string) {
	cpic = config.Config.Extra.Cpic
	cacheNoPaid = config.Config.Extra.CacheNoPaid
	var err error
	epubDetail := GetDetail(bid)
	fmt.Println(epubDetail.BookName, "/", epubDetail.AuthorName)
	epubChapters := GetCatalog(bid)
	var chapters []structure.ChapterList
	if !cacheNoPaid {
		for _, c := range epubChapters {
			if c.AuthAccess == "1" {
				chapters = append(chapters, c)
			}
		}
	} else {
		chapters = epubChapters
	}
	fmt.Println("fetching datas…")
	initTemp(epubDetail.BookName, epubDetail.AuthorName, epubDetail.Cover, epubChapters)
	epubTotalCount := len(chapters)
	epubBar = pb.StartNew(epubTotalCount)
	epubContainer := []int{}
	epubc := make(chan int, 1024)
	epubErrc := make(chan structure.ChapterList, 102400)
	coroutine := config.Config.Extra.Coroutines
	epubChaptersArr := splitArray(chapters, coroutine)
	for _, cs := range epubChaptersArr {
		go getChapterEpub(cs, epubc, epubErrc)
	}
	for {
		select {
		case epub := <-epubc:
			epubContainer = append(epubContainer, epub)
			epubBar.Increment()
		case e := <-epubErrc:
			go getChapterEpub([]structure.ChapterList{e}, epubc, epubErrc)
		}
		if len(epubContainer) == len(chapters) {
			break
		}
	}
	epubBar.Finish()
	fmt.Println("writing out files…")
	epubPath := filepath.Join(tmpPath, "..", epubDetail.BookName+".epub")
	err = util.CompressEpub(bookPath, epubPath)
	util.PanicErr(err)
	close(epubc)
	close(epubErrc)
	err = util.RemoveContents(tmpPath)
	util.PanicErr(err)
	fmt.Println("download success!")
	os.Exit(0)
}

func getChapterEpub(chapters []structure.ChapterList, epubc chan int, epubErrc chan structure.ChapterList) {
	for _, chapter := range chapters {
		chapterInfo, err := GetContent(chapter.ChapterID)
		if err != nil {
			epubErrc <- chapter
		} else {
			content := chapterInfo.TxtContent
			content += chapterInfo.AuthorSay
			content = strings.Replace(content, "&", "&#38;", -1)
			contentParas := strings.Split(content, "\n")
			contentS := contentStruct{chapterInfo.ChapterTitle, contentParas}
			fileName := oebpsPath + "/" + chapter.ChapterID + ".xhtml"
			var tmpl *template.Template
			if cpic {
				tmpl, _ = template.New("content").Funcs(template.FuncMap{"fixImg": getImg}).Parse(chapterTpl)
			} else {
				tmpl, _ = template.New("content").Funcs(template.FuncMap{"fixImg": fixImgTag}).Parse(chapterTpl)
			}
			file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
			err := tmpl.Execute(file, contentS)
			if err != nil {
				epubErrc <- chapter
			}
			epubc <- 0
		}
	}
}

func genBookToc(name string, chapters []structure.ChapterList) {
	tocS := bookTocStruct{name, chapters}
	fileName := oebpsPath + "/book-toc.xhtml"
	tmpl, _ := template.New("toc").Parse(bookTocTpl)
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	err := tmpl.Execute(file, tocS)
	util.PanicErr(err)
}

func genContentOpf(name, author string, chapters []structure.ChapterList) {
	opfS := opfStruct{name, author, chapters}
	fileName := oebpsPath + "/package.opf"
	tmpl, _ := template.New("opf").Parse(opfTpl)
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	err := tmpl.Execute(file, opfS)
	util.PanicErr(err)
}

func getCover(url string) string {
	client := req.New()
	client.SetTimeout(20 * time.Second)
	r, _ := client.Get(url)
	return r.String()
}
