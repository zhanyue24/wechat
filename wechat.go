package wechat

import (
	"github.com/sirupsen/logrus"
	"github.com/zhanyue24/wechat/config"
	"github.com/zhanyue24/wechat/modules/openPlatform"
)

var Wechat = new(WechatStru)

type WechatStru struct {
}

func OpenPlatform(cfg config.OpenPlatformCfg) *openPlatform.OpenPlatform {

	logrus.WithField("data", cfg).Info("load config")

	return &openPlatform.OpenPlatform{
		Config: cfg,
	}
}
