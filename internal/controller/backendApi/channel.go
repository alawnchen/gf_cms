package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

var (
	Channel = cChannel{}
)

type cChannel struct{}

// Index 获取后台栏目分类接口数据
func (c *cChannel) Index(ctx context.Context, req *backendApi.ChannelIndexApiReq) (res *backendApi.ChannelIndexApiRes, err error) {
	data, err := service.Channel().BackendIndex(ctx)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, 0, "获取成功", data)
	return
}

// Status 启用禁用
func (c *cChannel) Status(ctx context.Context, req *backendApi.ChannelStatusApiReq) (res *backendApi.ChannelStatusApiRes, err error) {
	_, err = service.Channel().BackendApiStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// Delete 删除
func (c *cChannel) Delete(ctx context.Context, req *backendApi.ChannelDeleteApiReq) (res *backendApi.ChannelDeleteApiRes, err error) {
	_, err = service.Channel().BackendApiDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}