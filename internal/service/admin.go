// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type IAdmin interface {
	LoginVerify(ctx context.Context, in model.AdminLoginInput) (admin *entity.CmsAdmin, err error)
	GetUserByUserNamePassword(ctx context.Context, in model.AdminLoginInput) g.Map
	GetRoleIdsByAccountId(accountId string) []gdb.Value
	BackendAdminGetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error)
	BackendApiAdminAdd(ctx context.Context, in *backendApi.AdminAddReq) (out interface{}, err error)
	BackendApiAdminEdit(ctx context.Context, in *backendApi.AdminEditReq) (out interface{}, err error)
	BackendApiAdminStatus(ctx context.Context, in *backendApi.AdminStatusReq) (out interface{}, err error)
	BackendApiAdminDelete(ctx context.Context, in *backendApi.AdminDeleteReq) (out interface{}, err error)
	BackendApiAdminDeleteBatch(ctx context.Context, in *backendApi.AdminDeleteBatchReq) (out interface{}, err error)
}

var localAdmin IAdmin

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
