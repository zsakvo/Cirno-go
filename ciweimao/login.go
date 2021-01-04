package ciweimao

import (
	"log"

	"../config"
	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

func Login(c structure.ConfigStruct) {
	paras := req.Param{
		"login_name": c.App.UserName,
		"passwd":     c.App.Password,
	}
	res := util.Get("/signup/login", paras)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.LoginStruct
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Fatalln(err)
	} else {
		config.Write(result.Data.LoginToken, result.Data.ReaderInfo.Account)
	}
}
