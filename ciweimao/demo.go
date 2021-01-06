package ciweimao

import (
	"os"
	"text/template"

	"../util"
)

func Demo() {
	type Cdemo struct {
		Name string
		ID   string
	}
	type OpfStruct struct {
		Name     string
		Author   string
		Chapters []Cdemo
	}
	sweaters := OpfStruct{"不可思议的幻想乡的旅途", "Kazami⑨Yuuka", []Cdemo{Cdemo{"第1章", "001"}, Cdemo{"第2章", "002"}, Cdemo{"第3章", "003"}}}
	tmpl, err := template.New("test").Parse(opfTpl)
	util.PanicErr(err)
	err = tmpl.Execute(os.Stdout, sweaters)
	util.PanicErr(err)
}

var opfTpl = `<?xml version="1.0" encoding="utf-8" standalone="no"?>
<package xmlns="http://www.idpf.org/2007/opf" xmlns:dc="http://purl.org/dc/elements/1.1/"
  xmlns:dcterms="http://purl.org/dc/terms/" version="3.0" xml:lang="en"
  unique-identifier="pub-identifier">
  <metadata>
    <dc:title id="pub-title">{{.Name}}</dc:title>
    <dc:language id="pub-language">zh</dc:language>
    <dc:creator>Cirno</dc:creator>
    <dc:contributor>{{.Author}}</dc:contributor> 
    <dc:rights>Copyright © 2015 Hangzhou Fantasy Technology NetworkCo.,Ltd.</dc:rights>
    <meta name="cover" content="cover"/>
  </metadata>
  <manifest>
    <item id="htmltoc" properties="nav" media-type="application/xhtml+xml" href="book-toc.xhtml"/>
    <item media-type="text/css" id="epub-css" href="css/epub.css"/>
    <item media-type="text/css" id="epub-tss-css" href="css/synth.css"/> 
    <item id="cover" properties="cover-image" href="covers/cover.jpg" media-type="image/jpeg"/>
{{ range .Chapters }}    <item id="id-{{.ID}}" href="{{.ID}}.xhtml" media-type="application/xhtml+xml"/>
{{ end }}    </manifest>
  <spine>
    <itemref idref="cover" linear="no"/>
    <itemref idref="htmltoc" linear="yes"/> 
{{ range .Chapters }}    <itemref idref="id-{{.ID}}"/>
{{ end }}    </spine>
</package>
`
