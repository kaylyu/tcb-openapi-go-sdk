package database

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/kaylyu/tcb-openapi-go-sdk/component"
	"github.com/kaylyu/tcb-openapi-go-sdk/component/database/query"
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"go.mongodb.org/mongo-driver/bson"
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
func (d *Database) GetDocument(table, docId string, limit, skip uint64, fields bson.M, sort bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents/%s", d.context.Config.EnvId, table, docId)
	params := make(url.Values)
	params.Set("limit", gconv.String(limit))
	params.Set("skip", gconv.String(skip))
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

//单文档替换更新
//https://docs.cloudbase.net/api-reference/openapi/database.html#setdocument
func (d *Database) SetDocument(table, docId string, data bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents/%s?transactionId=%s", d.context.Config.EnvId, table, docId, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"data": util.JsonEncode(data),
	})
}

//单文档插入
//https://docs.cloudbase.net/api-reference/openapi/database.html#insertdocument
func (d *Database) InsertDocument(table, docId string, data bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents/%s?transactionId=%s", d.context.Config.EnvId, table, docId, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"data": util.JsonEncode(data),
	})
}

//单文档删除
//https://docs.cloudbase.net/api-reference/openapi/database.html#deletedocument
func (d *Database) DeleteDocument(table, docId string, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents/%s?transactionId=%s", d.context.Config.EnvId, table, docId, d.appendTransActionId(transactionId...))

	return d.core.HttpDeleteJson(path, bson.M{})
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

//批量查询文档
//https://docs.cloudbase.net/api-reference/openapi/database.html#find
func (d *Database) Find(table string, query *query.Query, limit, skip uint64, fields bson.M, sort bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents:find?limit=%d&skip=%d&fields=%s&sort=%s&transactionId=%s",
		d.context.Config.EnvId, table, limit, skip, util.JsonEncode(fields), util.JsonEncode(sort), d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"query": query.ToString(),
	})
}

//统计查询文档总数
//https://docs.cloudbase.net/api-reference/openapi/database.html#count
func (d *Database) Count(table string, query *query.Query, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents:count?transactionId=%s", d.context.Config.EnvId, table, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"query": query.ToString(),
	})
}

//批量查询并更新单文档
//https://docs.cloudbase.net/api-reference/openapi/database.html#updateone
func (d *Database) UpdateOne(table string, query *query.Query, data bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents:updateOne?transactionId=%s", d.context.Config.EnvId, table, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"query": query.ToString(),
		"data":  util.JsonEncode(data),
	})
}

//批量查询并批量更新
//https://docs.cloudbase.net/api-reference/openapi/database.html#updatemany
func (d *Database) UpdateMany(table string, query *query.Query, data bson.M, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents:updateMany?transactionId=%s", d.context.Config.EnvId, table, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"query": query.ToString(),
		"data":  util.JsonEncode(data),
	})
}

//批量查询并删除单文档
//https://docs.cloudbase.net/api-reference/openapi/database.html#deleteone
func (d *Database) DeleteOne(table string, query *query.Query, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents:deleteOne?transactionId=%s", d.context.Config.EnvId, table, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"query": query.ToString(),
	})
}

//批量查询并批量删除
//https://docs.cloudbase.net/api-reference/openapi/database.html#deletemany
func (d *Database) DeleteMany(table string, query *query.Query, transactionId ...string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/documents:deleteMany?transactionId=%s", d.context.Config.EnvId, table, d.appendTransActionId(transactionId...))

	return d.core.HttpPostJson(path, bson.M{
		"query": query.ToString(),
	})
}

//开始事务
//https://docs.cloudbase.net/api-reference/openapi/database.html#transaction
func (d *Database) Transaction(table string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/transaction", d.context.Config.EnvId, table)

	return d.core.HttpPostJson(path, bson.M{})
}

//提交事务
//https://docs.cloudbase.net/api-reference/openapi/database.html#transaction
func (d *Database) CommitTransaction(table string, transactionId string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/%s:commit", d.context.Config.EnvId, table, transactionId)

	return d.core.HttpPostJson(path, bson.M{})
}

//提交事务
//https://docs.cloudbase.net/api-reference/openapi/database.html#transaction
func (d *Database) RollbackTransaction(table string, transactionId string) (body string, err error) {
	//请求参数
	path := fmt.Sprintf("/api/v2/envs/%s/databases/%s/%s:rollback", d.context.Config.EnvId, table, transactionId)

	return d.core.HttpPostJson(path, bson.M{})
}

//可直接使用Get请求
func (d *Database) HttpGetJson(path string, params url.Values, headers ...map[string]string) (body string, err error) {
	//请求
	return d.core.HttpGetJson(path, params, headers...)

}

//可直接使用Post请求
func (d *Database) HttpPostJson(path string, request interface{}, headers ...map[string]string) (body string, err error) {
	//请求
	return d.core.HttpPostJson(path, request, headers...)

}

//可直接使用Patch请求
func (d *Database) HttpPatchJson(path string, request interface{}, headers ...map[string]string) (body string, err error) {
	//请求
	return d.core.HttpPatchJson(path, request, headers...)

}

//可直接使用Delete请求
func (d *Database) HttpDeleteJson(path string, request interface{}, headers ...map[string]string) (body string, err error) {
	//请求
	return d.core.HttpDeleteJson(path, request, headers...)

}
