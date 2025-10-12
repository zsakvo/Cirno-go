package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/imroc/req"
	"github.com/zsakvo/Cirno-go/config"
)

var client *req.Req

func Post(url string, paras req.Param) ([]byte, error) {
	cfg := config.Config
	var err error
	var r *req.Resp
	var res []byte
	var param req.Param
	if url == "/signup/login" {
		param = req.Param{
			"app_version":  cfg.App.AppVersion,
			"device_token": cfg.App.DeviceToken,
		}
	} else {
		param = req.Param{
			"account":      cfg.App.Account,
			"device_token": cfg.App.DeviceToken,
			"app_version":  cfg.App.AppVersion,
			"login_token":  cfg.App.LoginToken,
		}
	}
	for k, v := range paras {
		param[k] = v
	}

	//加密字段处理 start
	randStr, err := GenerateRandomString(16)
	param["rand_str"] = randStr

	signature := "CkMxWNB666"
	signatureKey := cfg.App.SignatureKey
	signatures := fmt.Sprintf("%s%s", signatureKey, signature)
	signStr := fmt.Sprintf("account=%s&app_version=%s&rand_str=%s&signatures=%s", cfg.App.Account, cfg.App.AppVersion, randStr, signatures)
	encodeUrl := encodeURI(signStr)

	sha256Data := HmacSha256(signatureKey, encodeUrl)
	sha256Str := Base64Encode(sha256Data)
	param["p"] = sha256Str
	//加密字段处理 end

	if client == nil {
		client = req.New()
	}
	client.SetTimeout(20 * time.Second)
	if len(cfg.App.ProxyUrl) > 0 {
		client.SetProxyUrl(strings.TrimSpace(cfg.App.ProxyUrl))
	}
	//slog.Info("client.Post start", "param", param)
	r, err = client.Post(cfg.App.HostUrl+url, param, req.Header{"User-Agent": cfg.App.UserAgent, "Accept-Encoding": "gzip, deflate, br", "Accept-Language": "zh-Hant-HK;q=1, zh-Hans-CN;q=0.9, en-CN;q=0.8", "Accept": "*/*"})
	//slog.Info("client.Post result", "r", r)
	if err != nil {
		return nil, err
	}
	res, err = Decode(r.String(), cfg.App.DefaultKey)
	if !strings.Contains(string(res), "100000") {
		err = errors.New(ErrTip(string(res)))
	}
	return res, err
}
