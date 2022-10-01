package backendApi

import "github.com/gogf/gf/v2/frame/g"

type ChannelIndexApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目分类列表"`
}
type ChannelIndexApiRes struct {
}

type ChannelStatusApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目启用禁用"`
	Id     int `p:"id" name:"id" brief:"栏目ID" des:"栏目ID"  arg:"true" v:"required#请选择栏目ID"`
}
type ChannelStatusApiRes struct {
}

type ChannelDeleteApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目删除"`
	Id     int `p:"id" name:"id" brief:"栏目ID" des:"栏目ID"  arg:"true" v:"required#请选择栏目ID"`
}
type ChannelDeleteApiRes struct {
}

type ChannelAddApiReq struct {
	g.Meta           `tags:"Backend" method:"post" summary:"栏目添加、编辑"`
	Pid              int    `p:"pid" name:"pid" brief:"父级分类ID" des:"父级分类ID" arg:"true" d:"0" v:"required#请选择父级分类ID"`
	Name             string `p:"name" name:"name" brief:"分类名称" des:"分类名称"  arg:"true" v:"required#分类名称"`
	Thumb            string `p:"thumb" name:"thumb" brief:"栏目缩略图" des:"栏目缩略图"  arg:"true" v:""`
	Sort             int    `p:"sort" name:"sort" brief:"排序" des:"排序" arg:"true" d:"100" v:"required#请输入排序"`
	Status           int    `p:"status" name:"status" brief:"状态" des:"状态" arg:"true" d:"0" v:"required|in:0,1#请填写状态|状态不合法"`
	Description      string `p:"description" name:"description" brief:"栏目简介" des:"栏目简介"  arg:"true" v:""`
	Type             int    `p:"type" name:"type" brief:"栏目类型" des:"栏目类型" arg:"true" d:"0" v:"required|in:1,2,3#请选择栏目类型|栏目类型不合法"`
	LinkUrl          string `p:"link_url" name:"link_url" brief:"链接地址" des:"链接地址"  arg:"true" v:""`
	LinkTrigger      int    `p:"link_trigger" name:"link_trigger" brief:"打开方式" des:"打开方式" arg:"true" d:"0" v:"required|in:0,1#请选择打开方式|打开方式不合法"`
	Model            string `p:"model" name:"model" brief:"模型" des:"模型"  arg:"true" v:"required#请选择模型"`
	ListController   string `p:"list_controller" name:"list_controller" brief:"列表控制器" des:"列表控制器"  arg:"true" v:""`
	DetailController string `p:"detail_controller" name:"detail_controller" brief:"详情控制器" des:"详情控制器"  arg:"true" v:""`
	ListTemplate     string `p:"list_template" name:"list_template" brief:"列表模板" des:"列表模板"  arg:"true" v:""`
	DetailTemplate   string `p:"detail_template" name:"detail_template" brief:"详情模板" des:"详情模板"  arg:"true" v:""`
}
type ChannelAddApiRes struct {
}

type ChannelEditApiReq struct {
	Id int `p:"id" name:"id" brief:"分类ID" des:"分类ID" arg:"true" d:"0" v:"required|min:1#请选择父级分类ID|分类ID不能为0"`
	ChannelAddApiReq
}
type ChannelEditApiRes struct {
}
