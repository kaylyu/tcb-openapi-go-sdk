package database_test

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk"
	"github.com/kaylyu/tcb-openapi-go-sdk/config"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"testing"
	"time"
)

func TestFunction(t *testing.T) {
	tcb := tcb.NewTcb(&config.Config{
		EnvId:     "",
		Timeout:   time.Duration(15) * time.Second,
		LogPrefix: "tcb",
		Debug:     true,
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

	//插入记录
	fmt.Println(tcb.GetDatabase().HttpPostJson("/api/v2/envs/tcb-go-xxx/databases/test/documents", map[string]interface{}{
		"data": []string{
			util.JsonEncode(map[string]interface{}{
				"app_id": "xx",
				"phone":  "123",
			}),
			util.JsonEncode(map[string]interface{}{
				"app_id": "yy",
				"phone":  "1234",
			}),
		},
	}))
}
