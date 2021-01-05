package database

import (
	"github.com/kaylyu/tcb-openapi-go-sdk/component"
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
	"net/url"
)

//数据库
//https://docs.cloudbase.net/api-reference/openapi/database.html
type Database struct {
	context *context.Context
	core    *component.Core
}

/*
创建实例
*/
func NewDatabase(context *context.Context, core *component.Core) *Database {
	return &Database{context, core}
}

//TODO
func (d *Database) HttpGetJson(path string, params url.Values, headers ...map[string]string) (body string, err error) {
	//请求
	return d.core.HttpGetJson(path, params, headers...)

}

//TODO
func (d *Database) HttpPostJson(path string, request interface{}, headers ...map[string]string) (body string, err error) {
	//请求
	return d.core.HttpPostJson(path, request, headers...)

}
