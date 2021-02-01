package query

import "gopkg.in/mgo.v2/bson"

//判断字段是否存在
func (q *Query) Type(key string, value bool) *Query {
	q.query[key] = bson.M{"$type": value}
	return q
}

//判断字段是否存在
func (q *Query) Exists(key string, value bool) *Query {
	q.query[key] = bson.M{"$exists": value}
	return q
}
