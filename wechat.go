package wechat

import (
	"github.com/sirupsen/logrus"
	"github.com/zhanyue24/wechat/config"
	"github.com/zhanyue24/wechat/modules/official"
	"github.com/zhanyue24/wechat/modules/openPlatform"
)

func OpenPlatform(cfg config.OpenPlatformCfg) *openPlatform.OpenPlatform {

	logrus.WithField("data", cfg).Info("openplatform load config")

	return &openPlatform.OpenPlatform{
		Config: cfg,
	}
}

func Official(cfg config.OfficialCfg) *official.Official {

	logrus.WithField("data", cfg).Info("official load config")

	return &official.Official{
		Config: cfg,
	}
}
