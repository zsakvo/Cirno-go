package main

import (
	"fmt"
	"log"
	"os"

	"./ciweimao"
	"./fetch"
	"./util"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var canExec bool

func init() {
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	if !util.IsExist(expandedDir + "/Cirno/download") {
		err := os.MkdirAll(expandedDir+"/Cirno/download", os.ModePerm)
		util.PanicErr(err)
	}
	if util.IsExist(expandedDir + "/Cirno/config.yaml") {
		canExec = true
		util.InitReq()
	} else {
		canExec = false
	}
}

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			var args = c.Args()
			if args.Get(0) == "login" {
				ciweimao.Login()
			} else {
				if !canExec {
					fmt.Println("请先使用 login 命令登陆")
					os.Exit(0)
				}
				switch args.Get(0) {
				case "demo":
					ciweimao.Demo()
				case "login":
					ciweimao.Login()
				case "search":
					ciweimao.Search(args.Get(1), 0)
				case "download":
					fetch.DownloadText(args.Get(1))
				}
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
