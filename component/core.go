package component

import (
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
	"github.com/kaylyu/tcb-openapi-go-sdk/http"
	"github.com/kaylyu/tcb-openapi-go-sdk/sts"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"net/url"
)

type Core struct {
	ctx    *context.Context
	client *http.Client
}

func NewCore(ctx *context.Context, sts *sts.Sts) *Core {
	return &Core{ctx, http.NewHttpClient(ctx, sts)}
}

//GET
func (c *Core) HttpGetJson(path string, params url.Values, headers ...map[string]string) (response string, err error) {
	//附加数据
	response, err = c.client.HttpGetJson(path, params, headers...)
	if err != nil {
		return
	}
	return
}

//POST
func (c *Core) HttpPostJson(path string, request interface{}, headers ...map[string]string) (response string, err error) {
	//附加数据
	response, err = c.client.HttpPostJson(path, util.JsonEncode(request), headers...)
	if err != nil {
		return
	}
	return
}
