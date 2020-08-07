package wechat

import "github.com/sirupsen/logrus"

var Wechat = new(WechatStru)

type WechatStru struct {
}

func New() *WechatStru {
	return &WechatStru{}
}

func (m *WechatStru) OpenPlatform() {
	logrus.Info("xxxxx")
}
