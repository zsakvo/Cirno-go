package structure

type LoginStruct struct {
	Code int64     `json:"code"`
	Data LoginData `json:"data"`
}

type LoginData struct {
	LoginToken string           `json:"login_token"`
	UserCode   string           `json:"user_code"`
	ReaderInfo ReaderInfo       `json:"reader_info"`
	PropInfo   map[string]int64 `json:"prop_info"`
}

type ReaderInfo struct {
	ReaderID       string        `json:"reader_id"`
	Account        string        `json:"account"`
	IsBind         string        `json:"is_bind"`
	IsBindQq       int64         `json:"is_bind_qq"`
	IsBindWeixin   int64         `json:"is_bind_weixin"`
	IsBindHuawei   int64         `json:"is_bind_huawei"`
	IsBindApple    int64         `json:"is_bind_apple"`
	PhoneNum       string        `json:"phone_num"`
	MobileVal      string        `json:"mobileVal"`
	Email          string        `json:"email"`
	License        string        `json:"license"`
	ReaderName     string        `json:"reader_name"`
	AvatarURL      string        `json:"avatar_url"`
	AvatarThumbURL string        `json:"avatar_thumb_url"`
	BaseStatus     string        `json:"base_status"`
	ExpLV          string        `json:"exp_lv"`
	ExpValue       string        `json:"exp_value"`
	Gender         string        `json:"gender"`
	VipLV          string        `json:"vip_lv"`
	VipValue       string        `json:"vip_value"`
	IsAuthor       string        `json:"is_author"`
	IsUploader     int64         `json:"is_uploader"`
	BookAge        int64         `json:"book_age"`
	CategoryPrefer []interface{} `json:"category_prefer"`
	UsedDecoration []interface{} `json:"used_decoration"`
	Rank           string        `json:"rank"`
	Ctime          string        `json:"ctime"`
}
