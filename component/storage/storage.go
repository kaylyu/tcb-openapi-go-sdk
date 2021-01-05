package storage

import (
	"fmt"
	"github.com/kaylyu/tcb-openapi-go-sdk/component"
	"github.com/kaylyu/tcb-openapi-go-sdk/context"
)

//https://docs.cloudbase.net/api-reference/openapi/storage.html
type Storage struct {
	context *context.Context
	core    *component.Core
}

/*
创建实例
*/
func NewStorage(context *context.Context, core *component.Core) *Storage {
	return &Storage{context, core}
}

//获取文件上传属性
func (s *Storage) GetUploadMetaData(data interface{}) (out interface{}, err error) {
	//准备请求路径
	uri := fmt.Sprintf("/api/v2/envs/%s/storages:getUploadMetaData", s.context.Config.EnvId)

	//请求
	return s.core.HttpPostJson(uri, data)
}

//获取文件下载链接
func (s *Storage) BatchGetTempUrls(data interface{}) (out interface{}, err error) {
	//准备请求路径
	uri := fmt.Sprintf("/api/v2/envs/%s/storages:batchGetTempUrls", s.context.Config.EnvId)

	//请求
	return s.core.HttpPostJson(uri, data)

}

//批量删除文件
func (s *Storage) BatchDelete(data interface{}) (out interface{}, err error) {
	//准备请求路径
	uri := fmt.Sprintf("/api/v2/envs/%s/storages:batchDelete", s.context.Config.EnvId)

	//请求
	return s.core.HttpPostJson(uri, data)
}
