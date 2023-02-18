package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/role/add" method:"post" desc:"添加角色" tags:"role"`
	Name   string `json:"name" v:"required#名称必填" dc:"名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"post" desc:"修改角色" tags:"role"`
	Id     uint   `json:"id" v:"required#id必填" desc:"id"`
	Name   string `json:"name" v:"required#名称必填" dc:"名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleUpdateRes struct {
	Id uint `json:"id"`
}
type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"角色" summary:"删除角色接口"`
	Id     uint `v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RoleDeleteRes struct{}

type RoleGetListCommonReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色" summary:"修改角色列表"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int64       `json:"total" description:"数据总数"`
}
type AddPermissionReq struct {
	g.Meta       `path:"/role/add/permission" method:"post" tags:"角色" summary:"角色添加权限接口"`
	RoleId       uint `json:"role_id" desc:"角色id"`
	PermissionId uint `json:"permission_id" desc:"权限id"`
}

type AddPermissionRes struct {
	Id uint `json:"id"`
}

type DeletePermissionReq struct {
	g.Meta       `path:"/role/delete/permission" method:"delete" tags:"角色" summary:"角色删除权限接口"`
	RoleId       uint `json:"role_id" desc:"角色id"`
	PermissionId uint `json:"permission_id" desc:"权限id"`
}

type DeletePermissionRes struct{}
