package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台登录"`
}
type LoginRes struct {
	g.Meta `mime:"text/html" example:"string"`
}