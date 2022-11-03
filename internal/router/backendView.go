package router

import (
	"gf_cms/internal/controller/backend"
	"gf_cms/internal/logic/middleware"
	"gf_cms/internal/logic/util"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//后台view路由分组
func backendViewHandle(s *ghttp.Server) {
	var backendGroup = util.Util().BackendGroup()
	s.Group(backendGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
		)
		group.ALLMap(g.Map{
			"/admin/login": backend.Admin.Login,
		})
	})
	s.Group(backendGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			middleware.Middleware().BackendAuthSession,
			middleware.Middleware().BackendCheckPolicy,
		)
		group.ALLMap(g.Map{
			/*后台首页*/
			"/": backend.Index.Index,
			/*后台欢迎页*/
			"/welcome/index": backend.Welcome.Index,
			/*后台设置*/
			"/setting/index": backend.Setting.Index,
			/*管理员列表*/
			"/admin/index": backend.Admin.Index, //管理员列表
			"/admin/add":   backend.Admin.Add,   //添加
			"/admin/edit":  backend.Admin.Edit,  //编辑
			/*角色*/
			"/role/index": backend.Role.Index, //角色列表
			"/role/add":   backend.Role.Add,   //添加角色
			"/role/edit":  backend.Role.Edit,  //编辑角色
			/*栏目分类*/
			"/channel/index": backend.Channel.Index, //栏目分类列表
			"/channel/add":   backend.Channel.Add,   //添加栏目
			"/channel/edit":  backend.Channel.Edit,  //编辑栏目
			/*模型数据*/
			"/channel_model/index": backend.ChannelModel.Index, //列表
			/*文章*/
			"/article/move": backend.Article.Move, //移动文章
			"/article/add":  backend.Article.Add,  //新增文章
			"/article/edit": backend.Article.Edit, //编辑文章
			/**回收站**/
			"/recycle_bin/index": backend.RecycleBin.Index, //回收站列表
		})
	})
}
