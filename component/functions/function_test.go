package functions_test

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk"
	"github.com/kaylyu/tcb-openapi-go-sdk/config"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"testing"
	"time"
)

func TestFunction(t *testing.T) {
	tcb := tcb.NewTcb(&config.Config{
		EnvId:     "",
		Timeout:   time.Duration(15) * time.Second,
		LogPrefix: "tcb",
		Debug:     false,
		StsConfig: sts.Config{
			SecretId:        "",
			SecretKey:       "",
			Region:          "ap-guangzhou",
			Name:            "tcb",
			Policy:          `{"version":"2.0","statement":[{"effect":"allow","action":["tcb:*","scf:invocations"],"resource":["*"]}]}`,
			DurationSeconds: 7200,
			Debug:           true,
		},
		RedisConfig: gredis.Config{
			Host: "127.0.0.1",
			Port: 6379,
			Db:   1,
		},
	})

	fmt.Println(tcb.GetFunction().Invoke("test", map[string]interface{}{
		"data": map[string]interface{}{
			"path":       "/ping",
			"httpMethod": "GET",
			"body":       "",
		},
	}))
}
