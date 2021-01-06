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
	bookType := "txt"
	app := &cli.App{
		Name:  "cirno",
		Usage: "download e-books from hbooker.com",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "type",
				Value:       "txt",
				Usage:       "set books type",
				Aliases:     []string{"t"},
				Destination: &bookType,
			},
		},
		Action: func(c *cli.Context) error {
			var args = c.Args()
			if args.Get(0) == "login" {
				ciweimao.Login()
			} else {
				if !canExec {
					fmt.Println("Please login first.")
					os.Exit(0)
				}
				switch args.Get(0) {
				case "login":
					ciweimao.Login()
				case "search":
					ciweimao.Search(args.Get(1), 0)
				case "download":
					switch bookType {
					case "txt":
						fetch.DownloadText(args.Get(1))
					case "epub":
						fetch.DownloadEpub(args.Get(1))
					default:
						fmt.Println("invlid type.")
						os.Exit(0)
					}
				default:
					fmt.Println("invlid args.")
					os.Exit(0)
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
