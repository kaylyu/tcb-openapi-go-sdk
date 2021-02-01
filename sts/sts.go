package sts

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	v20180813 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
)

//https://cloud.tencent.com/document/product/598/33416
type Sts struct {
	*Config
	redis    *gredis.Redis
	cacheKey string
	logger   *logrus.Logger
}

//初始化实列
func NewStsClient(c *Config, redis *gredis.Redis, logger *logrus.Logger) *Sts {
	cacheKey := fmt.Sprintf("sts_federation_token_%s", c.SecretId)
	return &Sts{c, redis, cacheKey, logger}
}

//自定义调用方英文名称
func (c *Sts) SetName(name string) *Sts {
	c.Name = name
	return c
}

//授予该临时证书权限的CAM策略
func (c *Sts) SetPolicy(policy string) *Sts {
	c.Policy = policy
	return c
}

//刷新token
func (c *Sts) RefreshFederationToken() (response *v20180813.GetFederationTokenResponse, err error) {
	if c.Debug {
		fmt.Println("read direct from server")
	}
	//密钥
	credential := common.NewCredential(c.SecretId, c.SecretKey)
	client, err := v20180813.NewClient(credential, c.Region, profile.NewClientProfile())
	if err != nil {
		return
	}
	req := v20180813.NewGetFederationTokenRequest()
	req.Name = common.StringPtr(c.Name)                       //设置名称
	req.Policy = common.StringPtr(c.Policy)                   //设置授予该临时证书权限的CAM策略
	req.DurationSeconds = common.Uint64Ptr(c.DurationSeconds) //主动设置有效期
	//获取
	response, err = client.GetFederationToken(req)
	//添加到缓存中
	if c.redis != nil {
		//忽略缓存保存得异常
		expire := c.DurationSeconds
		if expire > 300 { //有效期偏移300s
			expire -= 300
		}
		_, _ = c.redis.Do("SETEX", c.cacheKey, expire, util.JsonEncode(response))
	}
	return
}

//获取临时token
func (c *Sts) GetFederationToken() (response *v20180813.GetFederationTokenResponse, err error) {
	//判断是否设置缓存
	if c.redis == nil {
		//直接获取并返回
		return c.RefreshFederationToken()
	}

	//先读取缓存中得token
	res, err := c.redis.Do("GET", c.cacheKey)
	//校验
	if err != nil || res == nil {
		//获取缓存失败，重新获取数据
		response, err = c.RefreshFederationToken()
		//直接返回
		return
	}
	if c.Debug {
		c.logger.Debug("read from redis cache...")
	}
	response = new(v20180813.GetFederationTokenResponse)
	if c.Debug {
		c.logger.Debugf("res cache:%v\n", string(res.([]uint8)))
	}
	//获取缓存中的数据为空
	if response.Response == nil {
		//获取缓存失败，重新获取数据
		response, err = c.RefreshFederationToken()
		//直接返回
		return
	}
	//解析
	err = util.JsonDecode(string(res.([]uint8)), response)
	return

}
