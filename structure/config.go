package structure

type ConfigStruct struct {
	App App `mapstructure:"app"`
	Web Web `mapstructure:"web"`
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
}

type Web struct {
	Port interface{} `mapstructure:"port"`
}
