package config

type OpenPlatformCfg struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Token     string `json:"token"`
	AesKey    string `json:"aes_key"`
}

type OfficialCfg struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Redirect  string `json:"redirect"`
	Token     string `json:"token"`
	AesKey    string `json:"aes_key"`
}
