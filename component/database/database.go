package database

import (
	"fmt"
	"github.com/kaylyu/tcb-openapi-go-sdk/component"
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"gopkg.in/mgo.v2/bson"
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

//处理事务
func (d *Database) appendTransActionId(transactionId ...string) string {
	if len(transactionId) == 0 {
		return ""
	}
	return transactionId[0]
}

//单文档查询
//https://docs.cloudbase.net/api-reference/openapi/database.html#getdocument
func (d *Database) GetDocument(table, docId, limit, skip string, fields bson.M, sort bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents/%s", d.context.Config.EnvId, table, docId)
	params := make(url.Values)
	params.Set("limit", limit)
	params.Set("skip", skip)
	params.Set("fields", util.JsonEncode(fields))
	params.Set("sort", util.JsonEncode(sort))
	params.Set("transactionId", d.appendTransActionId(transactionId...))

	return d.core.HttpGetJson(path, params)
}

//单文档更新
//https://docs.cloudbase.net/api-reference/openapi/database.html#updatedocument
func (d *Database) UpdateDocument(table, docId string, data bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents/%s", d.context.Config.EnvId, table, docId)

	return d.core.HttpPatchJson(path, bson.M{
		"data": util.JsonEncode(data),
	})
}

//单文档插入
//https://docs.cloudbase.net/api-reference/openapi/database.html#insertdocument
func (d *Database) InsertDocument(table, docId string, data bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents/{%s}?transactionId=%s", d.context.Config.EnvId, table, docId, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"data": util.JsonEncode(data),
	})
}

//批量插入文档
//https://docs.cloudbase.net/api-reference/openapi/database.html#insertdocuments
func (d *Database) InsertDocuments(table string, data []bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents?transactionId=%s", d.context.Config.EnvId, table, d.appendTransActionId(transactionId...))

	//待插入
	inserts := make([]string, 0)
	for _, item := range data {
		inserts = append(inserts, util.JsonEncode(item))
	}
	return d.core.HttpPostJson(path, bson.M{
		"data": inserts,
	})
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
