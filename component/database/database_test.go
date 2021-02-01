package database_test

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk"
	"github.com/kaylyu/tcb-openapi-go-sdk/component/database/query"
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
	fmt.Println(client.GetDatabase().GetDocument("users", "ba5028c16017631a000000186b9e2dd7", 10, 0, nil, nil))
}

//单文档更新
func TestUpdateDocument(t *testing.T) {
	fmt.Println(client.GetDatabase().UpdateDocument("users", "ba5028c16017631a000000186b9e2dd7", bson.M{"app_id": "ffff"}, ""))
}

//单文档替换更新
func TestSetDocument(t *testing.T) {
	fmt.Println(client.GetDatabase().SetDocument("users", "12333333333", bson.M{"app_id": "ffff"}, ""))
}

//单文档插入
func TestInsertDocument(t *testing.T) {
	fmt.Println(client.GetDatabase().InsertDocument("users", "124444", bson.M{
		"app_id":      "aaaa",
		"phone":       "123",
		"jobs":        []string{"技工", "学者"},
		"_createTime": util.Millisecond(),
		"_updateTime": util.Millisecond(),
	}))
}

//单文档删除
func TestDeleteDocument(t *testing.T) {
	fmt.Println(client.GetDatabase().DeleteDocument("users", "12333333333", ""))
}

//批量插入文档
func TestInsertDocuments(t *testing.T) {
	fmt.Println(client.GetDatabase().InsertDocuments("users", []bson.M{
		bson.M{
			"app_id":      "xxxxx1",
			"phone":       "123",
			"jobs":        []string{"技工", "学者"},
			"_createTime": util.Millisecond(),
			"_updateTime": util.Millisecond(),
		},
		bson.M{
			"app_id":      "yyyyyy2",
			"phone":       "1234",
			"jobs":        []string{"车工", "电工"},
			"_createTime": util.Millisecond(),
			"_updateTime": util.Millisecond(),
		},
	},
	))
}

func TestFind(t *testing.T) {
	qa := query.NewQuery()
	qa.Regex("app_id", "f")
	qa.Regex("phone", "1")
	fmt.Println(client.GetDatabase().Find("users", qa, 10, 0, bson.M{}, bson.M{}))
}

func TestCount(t *testing.T) {
	qa := query.NewQuery()
	qa.Regex("app_id", "f")
	qa.Regex("phone", "1")
	fmt.Println(client.GetDatabase().Count("users", qa))
}

func TestUpdateOne(t *testing.T) {
	qa := query.NewQuery()
	qa.Regex("app_id", "f")
	qa.Regex("phone", "1")

	fmt.Println(client.GetDatabase().UpdateOne("users", qa, bson.M{
		"app_id": "12kkk",
	}))
}

func TestUpdateMany(t *testing.T) {
	qa := query.NewQuery()
	qa.Regex("phone", "1")

	fmt.Println(client.GetDatabase().UpdateMany("users", qa, bson.M{
		"app_id": "66666",
	}))
}

func TestDeleteOne(t *testing.T) {
	qa := query.NewQuery()
	qa.Regex("phone", "1")

	fmt.Println(client.GetDatabase().DeleteOne("users", qa))
}

func TestDeleteMany(t *testing.T) {
	qa := query.NewQuery()
	qa.Regex("phone", "1")

	fmt.Println(client.GetDatabase().DeleteMany("users", qa))
}

func TestTransaction(t *testing.T) {
	fmt.Println(client.GetDatabase().Transaction("users"))
}

//测试AND查询条件
//更多请查看 https://docs.mongodb.com/manual/reference/operator/query-comparison/
func TestAnd(t *testing.T) {
	//{"query":"{\"$and\":[{\"app_id\":{\"$regex\":\"1\"}},{\"phone\":{\"$regex\":\"2\"}}]}"}
	qa := query.NewQuery()
	qa.Magic("$and", []bson.M{bson.M{"app_id": bson.M{"$regex": "1"}}, bson.M{"phone": bson.M{"$regex": "2"}}})
	fmt.Println(client.GetDatabase().Find("users", qa, 10, 0, bson.M{}, bson.M{}))
}
