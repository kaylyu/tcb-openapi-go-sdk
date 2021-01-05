package context

import (
	"github.com/kaylyu/tcb-openapi-go-sdk/config"
	"github.com/sirupsen/logrus"
)

//上下文
type Context struct {
	Config *config.Config
	Logger *logrus.Logger
}

/*
SetLogger 日志记录 默认输出到 os.Stdout
可以新建 logger 输出到指定文件
如果不想开启日志，可以输出到 /dev/null log.SetOutput(ioutil.Discard)
*/
func (c *Context) SetLogger(logger *logrus.Logger) {
	if logger != nil {
		c.Logger = logger
	}
}
