package frontend

import "github.com/gogf/gf/v2/frame/g"

type RotationGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int64       `json:"total" description:"数据总数"`
}

type RotationGetListCommonReq struct {
	g.Meta `path:"/frontend/rotation/list" method:"get" tags:"轮播图" summary:"修改轮播图列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
