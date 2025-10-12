package structure

type ConfigStruct struct {
	App   App   `mapstructure:"app"`
	Web   Web   `mapstructure:"web"`
	Extra Extra `mapstructure:"extra"`
}

type App struct {
	UserName    string `mapstructure:"user_name"`
	Password    string `mapstructure:"password"`
	Account     string `mapstructure:"account"`
	LoginToken  string `mapstructure:"login_token"`
	DeviceToken string `mapstructure:"device_token"`
	AppVersion  string `mapstructure:"app_version"`
	UserAgent   string `mapstructure:"user_agent"`
	DefaultKey  string `mapstructure:"default_key"`
	HostUrl     string `mapstructure:"host_url"`
	//入参签名key
	SignatureKey string `mapstructure:"signature_key"`
	//http://127.0.0.1:8888
	ProxyUrl string `mapstructure:"proxy_url"`
}

type Web struct {
	Port int `mapstructure:"port"`
}

type Extra struct {
	Coroutines  int  `mapstructure:"coroutines"`
	Cpic        bool `mapstructure:"cpic"`
	CacheNoPaid bool `mapstructure:"cache_no_paid"`
}
