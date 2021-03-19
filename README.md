# kaylyu/tcb-openapi-go-sdk

[Cloudbase Open API](https://docs.cloudbase.net/api-reference/openapi/introduction.html#liao-jie-qing-qiu-jie-gou) development sdk written in Golang

### 快速开始
```go
import github.com/kaylyu/tcb-openapi-go-sdk
```
### 示例
```go
tcb := tcb.NewTcb(&config.Config{
    EnvId:     "",                              //云开发环境ID
    Timeout:   time.Duration(15) * time.Second, //请求超时
    LogPrefix: "tcb",                           //日志头
    Debug:     false,                           //接口是否开启DEBUG
    SecretId:"",                                //腾讯云密钥SecretId
    SecretKey:"",                               //腾讯云密钥SecretKey
})

//根据文档ID查找对应云数据库表中记录
fmt.Println(client.GetDatabase().GetDocument("users", "124444", 10, 0, nil, nil))

//调用云函数
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
- 数据库 [示例](https://github.com/kaylyu/tcb-openapi-go-sdk/blob/master/component/database/database_test.go)

### License
MIT