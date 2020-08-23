package openPlatform

import (
	"github.com/zhanyue24/wechat/config"
)

type OfficialAccount struct {
	Config config.OfficialAccountCfg
}

type JssdkConfigReturn struct {
	Debug     bool   `json:"debug"`
	AppId     string `json:"appId"`
	TimeStamp int    `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
	JsApiList string `json:"jsApiList"`
}

func (m *OfficialAccount) JssdkConfig() (ret JssdkConfigReturn, err error) {

	ret.AppId = m.Config.AppId
	ret.Debug = m.

	return
}
