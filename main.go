package main

import (
	"log"
	"os"

	"./ciweimao"
	"./fetch"
	"./util"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func init() {
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	util.InitReq()
}

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			var args = c.Args()
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
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
