package functions

import (
	"fmt"
	"github.com/kaylyu/tcb-openapi-go-sdk/component"
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
)

//云函数
//https://docs.cloudbase.net/api-reference/openapi/function.html
type Function struct {
	context *context.Context
	core    *component.Core
}

/*
创建实例
*/
func NewFunction(context *context.Context, core *component.Core) *Function {
	return &Function{context, core}
}

//触发
func (f *Function) Invoke(functionName string, data interface{}) (out interface{}, err error) {

	//准备请求路径
	uri := fmt.Sprintf("/api/v2/envs/%s/functions/%s:invoke", f.context.Config.EnvId, functionName)

	//请求
	return f.core.HttpPostJson(uri, data)

}
