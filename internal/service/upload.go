// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
)

type IUpload interface {
	SingleUploadFile(ctx context.Context, in model.FileUploadInput, dir string) (out *backendApi.UploadFileRes, err error)
}

var localUpload IUpload

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}