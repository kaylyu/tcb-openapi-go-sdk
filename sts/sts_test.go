package sts_test

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestCam(t *testing.T) {
	client := sts.NewStsClient(&sts.Config{
		SecretId:        "",
		SecretKey:       "",
		Region:          "ap-guangzhou",
		Name:            "tcb",
		Policy:          `{"version":"2.0","statement":[{"effect":"allow","action":["tcb:*"],"resource":["*"]}]}`,
		DurationSeconds: 7200,
	}, gredis.New(gredis.Config{
		Host: "127.0.0.1",
		Port: 6379,
		Db:   0,
	}), &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &util.CustomerFormatter{Prefix: "TestSign"},
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: true,
	})
	rsp, err := client.GetFederationToken()
	fmt.Println(util.JsonEncode(rsp), err)
}
