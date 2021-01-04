package ciweimao

import (
	"fmt"

	"../util"
	"github.com/imroc/req"
)

func Demo() {
	util.InitReq()
	// var res = `{"code":100000,"data":{"login_token":"5f90deae3d935a7d667e46e5404267cb","user_code":"0b0ad490df4feef2b5d96d21996a9d7f","reader_info":{"reader_id":"37248","account":"\u4e66\u5ba2824724878","is_bind":"1","is_bind_qq":0,"is_bind_weixin":0,"is_bind_huawei":0,"is_bind_apple":0,"phone_num":"13234039720","mobileVal":"1","email":"528558128@qq.com","license":"","reader_name":"\u65e0\u8da3\u6c34\u74f6","avatar_url":"https:\/\/c2.kuangxiangit.com\/novel\/img\/37248\/avatar\/5c80162bd98403ccccfbceddfdd2f43e.jpg","avatar_thumb_url":"https:\/\/c2.kuangxiangit.com\/novel\/img\/37248\/avatar\/thumb_5c80162bd98403ccccfbceddfdd2f43e.jpg","base_status":"1","exp_lv":"4","exp_value":"297","gender":"1","vip_lv":"0","vip_value":"0","is_author":"1","is_uploader":0,"book_age":6,"category_prefer":[],"used_decoration":[],"rank":"0","ctime":"2015-11-19 19:01:26"},"prop_info":{"rest_gift_hlb":0,"rest_hlb":2,"rest_yp":0,"rest_recommend":0,"rest_total_blade":0,"rest_month_blade":0,"rest_total_100":0,"rest_total_588":0,"rest_total_1688":0,"rest_total_5000":0,"rest_total_10000":0,"rest_total_100000":0}}}`
	// var json = jsoniter.ConfigCompatibleWithStandardLibrary
	// var result structure.LoginStruct
	// json.Unmarshal([]byte(res), &result)
	// err := json.Unmarshal([]byte(res), &result)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	para := req.Param{
		"xxxx": "xxxx",
	}
	xx := util.Get("/xxxx", para)
	fmt.Println(string(xx))
	// fmt.Println(result.Data.LoginToken)
	// fmt.Println(result.Data.ReaderInfo.Account)
	// config.Write(result.Data.LoginToken, result.Data.ReaderInfo.Account)
}
