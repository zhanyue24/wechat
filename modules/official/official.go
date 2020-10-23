package official

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/ddliu/go-httpclient"
	"github.com/zhanyue24/wechat/config"
)

type Official struct {
	Config config.OfficialCfg
}

type GetUserInfoReturn struct {
	OpenId     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Unionid    string `json:"unionid"`
}

// https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID
func (m *Official) GetUserInfo(authInfo GetAccessTokenReturn) (ret GetUserInfoReturn, err error) {
	params := map[string]string{
		"access_token": authInfo.AccessToken,
		"openid":       authInfo.OpenId,
		"lang":         "zh-CN",
	}

	result, err := m.http("userinfo", params)
	if err != nil {
		return
	}

	json.Unmarshal(result, &ret)

	return
}

type GetAccessTokenReturn struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionid"`
}

func (m *Official) GetAccessToken(code string) (ret GetAccessTokenReturn, err error) {
	grantType := "authorization_code"

	params := map[string]string{
		"grant_type": grantType,
		"appid":      m.Config.AppId,
		"secret":     m.Config.AppSecret,
		"code":       code,
	}

	result, err := m.http("oauth2/access_token", params)
	if err != nil {
		return
	}

	json.Unmarshal(result, &ret)

	return
}

func (m *Official) GetAuthUrl(scope, state string) (u string, err error) {

	redirect := url.QueryEscape(m.Config.Redirect)

	u = fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect", m.Config.AppId, redirect, scope, state)

	return
}

type HttpErr struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type HttpReturn struct {
	HttpErr
}

func (m *Official) http(path string, params map[string]string) (ret []byte, err error) {

	baseUrl := "https://api.weixin.qq.com/sns/"
	url := fmt.Sprintf("%s%s", baseUrl, path)

	res, err := httpclient.Get(url, params)
	defer res.Body.Close()

	ret, err = res.ReadAll()

	var jsonData HttpErr

	json.Unmarshal(ret, &jsonData)

	if jsonData.ErrCode > 0 {
		err = errors.New(fmt.Sprintf("[%d] %s", jsonData.ErrCode, jsonData.ErrMsg))
		//logrus.WithError(err).WithFields(logrus.Fields{
		//	"url":    url,
		//	"params": params,
		//}).Warning("return error")
		return
	}

	return
}
