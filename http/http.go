package http

import (
	"crypto/tls"
	"fmt"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"github.com/kaylyu/tcb-openapi-go-sdk/util/loghttp"
	"net/http"
	"net/url"
	"time"
)

/*
Cloudbase Open API 服务地址
https://docs.cloudbase.net/api-reference/openapi/introduction.html
*/
var TcbAPI = "https://ap-guangzhou.tcb-api.tencentcloudapi.com"

//
type Client struct {
	ctx                 *context.Context
	sts                 *sts.Sts
	version             string //version
	authorizationHeader string //X-CloudBase-Authorization
	sessionTokenHeader  string //X-CloudBase-SessionToken
	timeStampHeader     string //X-CloudBase-TimeStamp
}

//获取实例
func NewHttpClient(ctx *context.Context, sts *sts.Sts) (client *Client) {
	return &Client{
		ctx,
		sts,
		"1.0",
		"X-CloudBase-Authorization",
		"X-CloudBase-SessionToken",
		"X-CloudBase-TimeStamp",
	}
}

//version
func (c *Client) SetVersion(version string) {
	c.version = version
}

//authorizationHeader
func (c *Client) SetAuthorizationHeader(authorizationHeader string) {
	c.authorizationHeader = authorizationHeader
}

//sessionTokenHeader
func (c *Client) SetSessionTokenHeader(sessionTokenHeader string) {
	c.sessionTokenHeader = sessionTokenHeader
}

//timeStampHeader
func (c *Client) SetTimeStampHeader(timeStampHeader string) {
	c.timeStampHeader = timeStampHeader
}

//请求
func (c *Client) request(method string, reqUrl string, reqBody string, headers ...map[string]string) (body string, err error) {
	uri, err := url.Parse(TcbAPI + reqUrl)
	if err != nil {
		return
	}
	//默认超时10S
	timeout := c.ctx.Config.Timeout
	if timeout == 0 {
		timeout = time.Duration(10) * time.Second
	}
	req := ghttp.NewClient()
	transport := &loghttp.Transport{
		Transport: &http.Transport{
			// No validation for https certification of the server in default.
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DisableKeepAlives:   true,
			MaxIdleConns:        0,
			MaxIdleConnsPerHost: 1000,
		},
	}
	if c.ctx.Config.Debug && c.ctx.Logger != nil {
		transport.Logger = c.ctx.Logger
	}
	req.SetTimeOut(timeout)
	req.Transport = transport

	//获取CAM临时TOKEN
	rsp, err := c.sts.GetFederationToken()
	if err != nil {
		return
	}
	//临时证书
	credentials := rsp.Response.Credentials

	//签名参数
	signIn := &signIn{
		host:        "api.tcloudbase.com",
		uri:         "//api.tcloudbase.com/",
		method:      "POST",
		contentType: "application/json; charset=utf-8",
		queryString: uri.RawQuery,
		payload:     []byte(""),
		now:         time.Now(),
		secretId:    *credentials.TmpSecretId,
		secretKey:   *credentials.TmpSecretKey,
		debug:       c.ctx.Config.Debug,
	}
	//签名
	_, authorization, err := c.signature(signIn)
	if err != nil {
		return
	}
	req.SetHeader("Accept", "*/*")
	req.SetHeader("Accept-Charset", "utf-8")
	req.SetHeader("Content-Type", signIn.contentType)
	if len(c.version) > 0 {
		req.SetHeader(c.authorizationHeader, fmt.Sprintf("%s %s", c.version, authorization))
	} else {
		req.SetHeader(c.authorizationHeader, authorization)
	}
	req.SetHeader(c.sessionTokenHeader, *credentials.Token)
	req.SetHeader(c.timeStampHeader, gconv.String(signIn.now.Unix()))
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.SetHeader(k, v)
		}
	}

	body = req.RequestContent(method, uri.String(), reqBody)

	return
}

//GET
func (c *Client) HttpGetJson(path string, params url.Values, headers ...map[string]string) (body string, err error) {
	//组装
	reqUrl := fmt.Sprintf("%s?%s", path, gurl.BuildQuery(params))

	//请求
	body, err = c.request("GET", reqUrl, "", headers...)

	return
}

//POST
func (c *Client) HttpPostJson(path string, params string, headers ...map[string]string) (body string, err error) {
	//请求
	body, err = c.request("POST", path, params, headers...)

	return
}

//PATCH
func (c *Client) HttpPatchJson(path string, params string, headers ...map[string]string) (body string, err error) {
	//请求
	body, err = c.request("PATCH", path, params, headers...)

	return
}
