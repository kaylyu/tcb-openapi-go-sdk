package query

import "go.mongodb.org/mongo-driver/bson"

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
