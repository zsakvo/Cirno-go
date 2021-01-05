package ciweimao

import (
	"io/ioutil"
	"os"
	"time"

	"../util"
	"github.com/imroc/req"
)

func Demo() {
	jpg := "https://c1.kuangxiangit.com/uploads/allimg/c170225/25-02-17151236-6938-100024573.jpg"
	client := req.New()
	client.SetTimeout(20 * time.Second)
	r, _ := client.Get(jpg)
	println(r.String())
	writeOut(r.String(), "./", "cover.jpg")
}

func writeOut(content, dirPath, fileName string) error {
	if !isExist(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		util.PanicErr(err)
	}
	outPath := dirPath + fileName
	d := []byte(content)
	err := ioutil.WriteFile(outPath, d, 0644)
	return err
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
