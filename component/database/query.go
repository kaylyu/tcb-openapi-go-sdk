package database

import (
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"gopkg.in/mgo.v2/bson"
)

//查询条件
type Query struct {
	query bson.M
}

func NewQuery() *Query {
	return &Query{query: make(bson.M)}
}

//等于
func (q *Query) Eq(key string, value string) *Query {
	q.query[key] = bson.M{"$eq": value}
	return q
}

//输出字符串
func (q *Query) ToString() string {
	return util.JsonEncode(q.query)
}
