package main

import (
	"log"
	"os"

	"./ciweimao"
	"./config"
	"./structure"
	"github.com/urfave/cli"
)

// func init(){

// }

var AppConfig structure.ConfigStruct

func main() {
	config := config.Load()
	AppConfig = config
	app := &cli.App{
		Action: func(c *cli.Context) error {
			var args = c.Args()
			switch args.Get(0) {
			case "login":
				ciweimao.Login(config)
				break
			case "search":
				ciweimao.Search(args.Get(1), 0, config)
				break
			case "download":
				ciweimao.GetCatalog(args.Get(1), config)
				break
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
