package functions_test

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk"
	"github.com/kaylyu/tcb-openapi-go-sdk/config"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"github.com/spf13/viper"
	"testing"
	"time"
)

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

func TestFunction(t *testing.T) {
	fmt.Println(client.GetFunction().Invoke("test", map[string]interface{}{
		"data": map[string]interface{}{
			"path":       "/ping",
			"httpMethod": "GET",
			"body":       "",
		},
	}))
}
