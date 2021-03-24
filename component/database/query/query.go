package query

import (
	"github.com/kaylyu/tcb-openapi-go-sdk/util"
	"go.mongodb.org/mongo-driver/bson"
)

//查询条件
//https://docs.mongodb.com/manual/reference/operator/query-comparison/
type Query struct {
	query bson.M
}

func NewQuery() *Query {
	return &Query{
		query: make(bson.M),
	}
}

//魔法方法
/*
如
Magic("$and", []bson.M{bson.M{"app_id": bson.M{"$regex": "1"}},bson.M{"phone": bson.M{"$regex": "2"}}})
*/
func (q *Query) Magic(key string, value interface{}) *Query {
	q.query[key] = value
	return q
}

//输出字符串
func (q *Query) ToString() string {
	return util.JsonEncode(q.query)
}
