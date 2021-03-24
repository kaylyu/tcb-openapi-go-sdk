package query

import "go.mongodb.org/mongo-driver/bson"

//等于
func (q *Query) Eq(key string, value interface{}) *Query {
	q.query[key] = bson.M{"$eq": value}
	return q
}

//不等于
func (q *Query) Neq(key string, value interface{}) *Query {
	q.query[key] = bson.M{"$neq": value}
	return q
}

//小于
func (q *Query) Lt(key string, value interface{}) *Query {
	q.query[key] = bson.M{"$lt": value}
	return q
}

//小于或等于
func (q *Query) Lte(key string, value interface{}) *Query {
	q.query[key] = bson.M{"$lte": value}
	return q
}

//大于
func (q *Query) Gt(key string, value interface{}) *Query {
	q.query[key] = bson.M{"$gt": value}
	return q
}

//大于或等于
func (q *Query) Gte(key string, value interface{}) *Query {
	q.query[key] = bson.M{"$gte": value}
	return q
}

//匹配所有
func (q *Query) All(key string, values []interface{}) *Query {
	q.query[key] = bson.M{"$in": values}
	return q
}

//字段值在给定数组中
func (q *Query) In(key string, values []interface{}) *Query {
	q.query[key] = bson.M{"$in": values}
	return q
}

//字段值不在给定数组中
func (q *Query) Nin(key string, values []interface{}) *Query {
	q.query[key] = bson.M{"$nin": values}
	return q
}
