package database_test

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk"
	"github.com/kaylyu/tcb-openapi-go-sdk/config"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

//
var client *tcb.Tcb

//初始化
func init() {
	envId := viper.GetString("env_id")
	//小程序handle
	client = tcb.NewTcb(&config.Config{
		EnvId:     envId,
		Timeout:   time.Duration(15) * time.Second,
		LogPrefix: viper.GetString("sts_name"),
		Debug:     viper.GetBool("tcb_open_api_debug"),
		StsConfig: sts.Config{
			SecretId:        viper.GetString("sts_app_id"),
			SecretKey:       viper.GetString("sts_secret"),
			Region:          viper.GetString("sts_region"),
			Name:            viper.GetString("sts_name"),
			Policy:          viper.GetString("sts_policy"),
			DurationSeconds: viper.GetUint64("sts_duration_seconds"),
			Debug:           viper.GetBool("sts_debug"),
		},
		RedisConfig: gredis.Config{
			Host: viper.GetString("redis_host"),
			Port: viper.GetInt("redis_port"),
			Db:   viper.GetInt("redis_db"),
			Pass: viper.GetString("redis_pwd"),
		},
	})
}

//单文档插入
func TestGetDocument(t *testing.T) {
	//插入记录
	fmt.Println(client.GetDatabase().GetDocument("users", "1d7219966013df210000000c00ba0e33", "10", "", nil, nil))
}

//单文档更新
func TestUpdateDocument(t *testing.T) {
	//插入记录
	fmt.Println(client.GetDatabase().UpdateDocument("users", "1d7219966013df210000000d04060cf8", bson.M{"app_id": "kkkkkkk"}, ""))
}

//单文档插入
func TestInsertDocument(t *testing.T) {
	//插入记录
	fmt.Println(client.GetDatabase().InsertDocument("users", "124444", bson.M{
		"app_id":      "aaaa",
		"phone":       "123",
		"jobs":        []string{"技工", "学者"},
		"_createTime": util.Unix(),
		"_updateTime": util.Unix(),
	}))
}

//批量插入文档
func TestInsertDocuments(t *testing.T) {
	fmt.Println(client.GetDatabase().InsertDocuments("users", []bson.M{
		bson.M{
			"app_id":      "xxxxx1",
			"phone":       "123",
			"jobs":        []string{"技工", "学者"},
			"_createTime": util.Unix(),
			"_updateTime": util.Unix(),
		},
		bson.M{
			"app_id":      "yyyyyy2",
			"phone":       "1234",
			"jobs":        []string{"车工", "电工"},
			"_createTime": util.Unix(),
			"_updateTime": util.Unix(),
		},
	},
	))
}

func TestFunction(t *testing.T) {

	//查询数据，有两种查询模式
	//1、精确查找
	//fmt.Println(client.GetDatabase().HttpPostJson(fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents:find?limit=100&skip=&fields=&sort=%s&transactionId=", envId, "users", util.JsonEncode(bson.M{"_id":-1})), bson.M{
	//	"query": util.JsonEncode(bson.M{
	//		"app_id": "xx",
	//	}),
	//}))
	//2、查询指令方式查找，参考文档：https://docs.cloudbase.net/database/query.html#cha-xun-zhi-ling
	//参考js https://github.com/TencentCloudBase/tcb-js-sdk-database/blob/master/src/commands/query.ts
	//fmt.Println(client.GetDatabase().HttpPostJson("/api/v2/envs/tcb-go-xxx/databases/test/documents:find?limit=100&skip=&fields=&sort=&transactionId=", bson.M{
	//	"query": util.JsonEncode(bson.M{
	//		//"app_id": bson.M{"$eq":"xx"},//精准匹配
	//		"app_id": bson.M{"$regex": "x"}, //模糊匹配
	//		//"jobs": bson.M{"$in":[]string{"车工"}},//过滤数组
	//		//"jobs": bson.M{"$regex": "车"}, //模糊匹配
	//	}),
	//}))
}
