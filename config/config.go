package config

type OpenPlatformCfg struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Token     string `json:"token"`
	AesKey    string `json:"aes_key"`
}

type OfficialAccountCfg struct {
	AppId  string `json:"app_id"`
	Debug  bool   `json:"debug"`
	Token  string `json:"token"`
	AesKey string `json:"aes_key"`
}
