# kaylyu/tcb-openapi-go-sdk

[Cloudbase Open API](https://docs.cloudbase.net/api-reference/openapi/introduction.html#liao-jie-qing-qiu-jie-gou) development sdk written in Golang

### 快速开始
```go
import github.com/kaylyu/tcb-openapi-go-sdk
```

### 注意事项
- ⚠需要提前开通云开发服务并创建环境，否则无法使用
- ⚠使用子账户模式，请先通过主账户授权开通 QcloudTCBFullAccess(云开发全读写访问), QcloudAccessForTCBRole(云开发对云资源的访问权限) [子账户权限设置指引](https://cloud.tencent.com/document/product/598/36256)

### 示例
- 注意，data中参数形式需符合对应请求云函数的参数接收规则，本例中的接收参数规则为腾讯云网关API触发模式，参数可参考文档 [API 网关触发器概述](https://cloud.tencent.com/document/product/583/12513)

```go
tcb := tcb.NewTcb(&config.Config{
    EnvId:     "",
    Timeout:   time.Duration(15) * time.Second,
    LogPrefix: "tcb",
    Debug:     false,
    StsConfig: sts.Config{ //参考文档：https://cloud.tencent.com/document/product/598/33416
        SecretId:        "",
        SecretKey:       "",
        Region:          "ap-guangzhou",
        Name:            "tcb",
        Policy:          `{"version":"2.0","statement":[{"effect":"allow","action":["tcb:*","scf:invocations"],"resource":["*"]}]}`,//参考文档：https://cloud.tencent.com/document/product/598/10603
        DurationSeconds: 7200,
        Debug:           true,
    },
    RedisConfig: gredis.Config{ //用于存储STS临时信息，可不传，每次都从远程服务器获取
        Host: "127.0.0.1",
        Port: 6379,
        Db:   1,
    },
})

fmt.Println(tcb.GetFunction().Invoke("test_func", map[string]interface{}{
    "data": map[string]interface{}{
        "path":       "/ping",
        "httpMethod": "GET",
        "body":       "",
    },
}))

```

### 支持
- 云函数
- 文件存储

### 部分支持
- 数据库 [示例](https://github.com/kaylyu/tcb-openapi-go-sdk/blob/master/component/database/database_test.go)

### License
MIT