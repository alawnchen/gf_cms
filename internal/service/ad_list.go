// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model/entity"
)

type (
	IAdList interface {
		Add(ctx context.Context, req *backendApi.AdListAddReq) (out interface{}, err error)
		Edit(ctx context.Context, req *backendApi.AdListEditReq) (out interface{}, err error)
		Delete(ctx context.Context, req *backendApi.AdListDeleteReq) (out interface{}, err error)
		BatchStatus(ctx context.Context, req *backendApi.AdListBatchStatusReq) (out interface{}, err error)
		Sort(ctx context.Context, req *backendApi.AdListSortReq) (out interface{}, err error)
		GetAdInfoById(ctx context.Context, id int) (out interface{}, err error)
		PcHomeListByChannelId(ctx context.Context, channelId int) (out []*entity.CmsAd, err error)
	}
)

var (
	localAdList IAdList
)

func AdList() IAdList {
	if localAdList == nil {
		panic("implement not found for interface IAdList, forgot register?")
	}
	return localAdList
}

func RegisterAdList(i IAdList) {
	localAdList = i
}
