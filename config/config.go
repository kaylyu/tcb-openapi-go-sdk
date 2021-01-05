package config

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"time"
)

type Config struct {
	EnvId       string        //TCB 环境 ID
	Timeout     time.Duration //请求超时设置
	LogPrefix   string        //日志前缀
	Debug       bool          //debug
	StsConfig   sts.Config    //cam config
	RedisConfig gredis.Config //redis config
}
