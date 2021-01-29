package database_test

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk"
	"github.com/kaylyu/tcb-openapi-go-sdk/config"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"gopkg.in/mgo.v2/bson"
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
				"jobs":   []string{"技工", "学者"},
			}),
			util.JsonEncode(map[string]interface{}{
				"app_id": "yy",
				"phone":  "1234",
				"jobs":   []string{"车工", "电工"},
			}),
		},
	}))
	//查询数据，有两种查询模式
	//1、精确查找
	fmt.Println(tcb.GetDatabase().HttpPostJson("/api/v2/envs/tcb-go-xxx/databases/test/documents:find?limit=100&skip=&fields=&sort=&transactionId=", bson.M{
		"query": util.JsonEncode(bson.M{
			"app_id": "xx",
		}),
	}))
	//2、查询指令方式查找，参考文档：https://docs.cloudbase.net/database/query.html#cha-xun-zhi-ling
	//参考js https://github.com/TencentCloudBase/tcb-js-sdk-database/blob/master/src/commands/query.ts
	fmt.Println(tcb.GetDatabase().HttpPostJson("/api/v2/envs/tcb-go-xxx/databases/test/documents:find?limit=100&skip=&fields=&sort=&transactionId=", bson.M{
		"query": util.JsonEncode(bson.M{
			//"app_id": bson.M{"$eq":"xx"},//精准匹配
			"app_id": bson.M{"$regex": "x"}, //模糊匹配
			//"jobs": bson.M{"$in":[]string{"车工"}},//过滤数组
			"jobs": bson.M{"$regex": "车"}, //模糊匹配
		}),
	}))

}
