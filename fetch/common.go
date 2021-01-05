package fetch

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"../structure"
	"../util"
)

type chapterStruct struct {
	Text string
	Cid  string
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

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func writeOut(content, dirPath, fileName string) error {
	if !isExist(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		util.PanicErr(err)
	}
	outPath := filepath.Join(dirPath, fileName)
	d := []byte(content)
	err := ioutil.WriteFile(outPath, d, 0644)
	return err
}

var contentHeader = "<?xml version=\"1.0\" encoding=\"utf-8\" ?><html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"zh-CN\"> <head> <meta http-equiv=\"Content-Type\" content=\"application/xhtml+xml; charset=utf-8\" /> <meta name=\"generator\" content=\"EasyPub v1.44\" /> <title> chapter 7 - 0 </title> <link rel=\"stylesheet\" href=\"style.css\" type=\"text/css\"/> </head> <body>"
var contentFooter = "</body></html>"
var container = "<?xml version=\"1.0\"?> <container version=\"1.0\" xmlns=\"urn:oasis:names:tc:opendocument:xmlns:container\"> <rootfiles> <rootfile full-path=\"OEBPS/content.opf\" media-type=\"application/oebps-package+xml\"/> </rootfiles> </container>"
var mimetype = "application/epub+zip"
var coverHeader = "<?xml version=\"1.0\" encoding=\"utf-8\" ?> <!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\"> <html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"zh-CN\"> <head> <meta http-equiv=\"Content-Type\" content=\"application/xhtml+xml; charset=utf-8\" /> <meta name=\"generator\" content=\"EasyPub v1.44\" /> <title> Cover </title> <style type=\"text/css\"> html,body {height:100%; margin:0; padding:0;} .wedge {float:left; height:50%; margin-bottom: -337px;} .container {clear:both; height:0em; position: relative;} table, tr, th {height: 675px; width:100%; text-align:center;} </style> <link rel=\"stylesheet\" href=\"style.css\" type=\"text/css\"/> </head> <body> <div class=\"wedge\"></div> <div class=\"container\"> <table><tr><td>"
var coverFooter = "</td></tr></table></div></body></html>"
var css = "@font-face{font-family:easypub;src:url(res:///system/fonts/DroidSansFallback.ttf),url(res:///ebook/fonts/../../system/fonts/DroidSansFallback.ttf)}@page{margin-top:0;margin-bottom:0}body{font-family:easypub;padding:0;margin-left:0;margin-right:0;orphans:0;widows:0}p{font-family:easypub;font-size:120%;line-height:125%;margin-top:5px;margin-bottom:0;margin-left:0;margin-right:0;orphans:0;widows:0}.a{text-indent:0}div.centeredimage{text-align:center;display:block;margin-top:.5em;margin-bottom:.5em}img.attpic{border:1px solid #000;max-width:100%;margin:0}.booktitle{margin-top:30%;margin-bottom:0;border-style:none solid none none;border-width:50px;border-color:#4E594D;font-size:3em;line-height:120%;text-align:right}.bookauthor{margin-top:0;border-style:none solid none none;border-width:50px;border-color:#4E594D;page-break-after:always;font-size:large;line-height:120%;text-align:right}.titlel1std,.titlel1top,.titlel2std,.titlel2top,.titlel3std,.titlel3top,.titlel4std,.titletoc{margin-top:0;border-style:none double none solid;border-width:0 5px 0 20px;border-color:#586357;background-color:#C1CCC0;padding:45px 5px 5px 5px;font-size:x-large;line-height:115%;text-align:justify}.titlel1single,.titlel2single,.titlel3single{margin-top:35%;border-style:none solid none none;border-width:30px;border-color:#4E594D;padding:30px 5px 5px 5px;font-size:x-large;line-height:125%;text-align:right}.toc{margin-left:16%;padding:0;line-height:130%;text-align:justify}.toc a{text-decoration:none;color:#000}.tocl1{margin-top:.5em;margin-left:-30px;border-style:none double double solid;border-width:0 5px 2px 20px;border-color:#6B766A;line-height:135%;font-size:132%}.tocl2{margin-top:.5em;margin-left:-20px;border-style:none double none solid;border-width:0 2px 0 10px;border-color:#939E92;line-height:123%;font-size:120%}.tocl3{margin-top:.5em;margin-left:-20px;border-style:none double none solid;border-width:0 2px 0 8px;border-color:#939E92;line-height:112%;font-size:109%}.tocl4{margin-top:.5em;margin-left:-20px;border-style:none double none solid;border-width:0 2px 0 6px;border-color:#939E92;line-height:115%;font-size:110%}.subtoc{margin-left:15%;padding:0;text-align:justify}.subtoclist{margin-top:.5em;margin-left:-20px;border-style:none double none solid;border-width:0 2px 0 10px;border-color:#939E92;line-height:123%;font-size:120%}"
var tocNcxHeader = "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"no\"?> <!DOCTYPE ncx PUBLIC \"-//NISO//DTD ncx 2005-1//EN\" \"http://www.daisy.org/z3986/2005/ncx-2005-1.dtd\"> <ncx xmlns=\"http://www.daisy.org/z3986/2005/ncx/\" version=\"2005-1\"> <head> <meta name=\"cover\" content=\"cover\"/> <meta name=\"dtb:uid\" content=\"easypub-9b14e97e\" /> <meta name=\"dtb:depth\" content=\"1\"/> <meta name=\"dtb:generator\" content=\"EasyPub v1.44\"/> <meta name=\"dtb:totalPageCount\" content=\"0\"/> <meta name=\"dtb:maxPageNumber\" content=\"0\"/> </head>"
var tocNcxFooter = "</navMap></ncx>"
var contentOpfHeader = "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"no\"?> <package version=\"2.0\" xmlns=\"http://www.idpf.org/2007/opf\" unique-identifier=\"bookid\"> <metadata xmlns:dc=\"http://purl.org/dc/elements/1.1/\" xmlns:opf=\"http://www.idpf.org/2007/opf\"> <dc:identifier id=\"bookid\">easypub-9b14e97e</dc:identifier> <dc:title>bookTitle</dc:title> <dc:creator opf:role=\"aut\">bookAuthor</dc:creator> <dc:date>2019</dc:date> <dc:rights>Created with EasyPub v1.44</dc:rights> <dc:language>zh-CN</dc:language> <meta name=\"cover\" content=\"cover-image\"/> </metadata>"
var contentOpfManifestHeader = "<manifest><item id=\"ncxtoc\" href=\"toc.ncx\" media-type=\"application/x-dtbncx+xml\"/> <item id=\"htmltoc\" href=\"book-toc.html\" media-type=\"application/xhtml+xml\"/> <item id=\"css\" href=\"style.css\" media-type=\"text/css\"/> <item id=\"cover-image\" href=\"cover.jpg\" media-type=\"image/jpeg\"/> <item id=\"cover\" href=\"cover.html\" media-type=\"application/xhtml+xml\"/>"
var contentOpfManifestFooter = "</manifest>"
var contentOpfNcxtocHeader = "<spine toc=\"ncxtoc\"> <itemref idref=\"cover\" linear=\"no\"/> <itemref idref=\"htmltoc\" linear=\"yes\"/>"
var contentOpfNcxtocFooter = "</spine>"
var contentOpfFooter = "<guide> <reference href=\"cover.html\" type=\"cover\" title=\"Cover\"/> <reference href=\"book-toc.html\" type=\"toc\" title=\"Table Of Contents\"/> <reference href=\"chapter0.html\" type=\"text\" title=\"Beginning\"/> </guide> </package>"
var contentOpfManifest string
var contentOpfNcxtoc string
var bookTocHeader = "<?xml version=\"1.0\" encoding=\"utf-8\" ?> <!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\"> <html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"zh-CN\"> <head> <meta http-equiv=\"Content-Type\" content=\"application/xhtml+xml; charset=utf-8\" /> <meta name=\"generator\" content=\"EasyPub v1.44\" /> <title> Table Of Contents </title> <link rel=\"stylesheet\" href=\"style.css\" type=\"text/css\"/> </head> <body> <h2 class=\"titletoc\"> 目录 </h2> <div class=\"toc\"> <dl>"
var bookTocFooter = "</dl></div></body></html>"
