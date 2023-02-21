package backend

import "github.com/gogf/gf/v2/frame/g"

type CategoryReq struct {
	g.Meta `path:"/category/add" tags:"Category" method:"post" summary:"You first category api"`
	CommonAddUpdate
}

type CommonAddUpdate struct {
	ParentId uint   `json:"parent_id" dc:"父级ID"`
	Name     string `json:"name" v:"required#名称不能为空"`
	PicUrl   string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Level    uint8  `json:"level"  dc:"等级 默认为1"`
	Sort     uint8  `json:"sort"     dc:"排序"`
}

type CategoryRes struct {
	CategoryId int `json:"category_id"`
}
type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"商品分类" summary:"删除商品分类接口"`
	Id     uint `v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
}
type CategoryDeleteRes struct{}

type CategoryUpdateReq struct {
	g.Meta `path:"/category/update/{Id}" method:"post" tags:"商品分类" summary:"修改商品分类接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品分类" dc:"商品分类Id"`
	CommonAddUpdate
}
type CategoryUpdateRes struct {
	Id uint `json:"id"`
}

type CategoryGetListCommonReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"商品分类" summary:"修改商品分类列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CategoryGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int64       `json:"total" description:"数据总数"`
}
