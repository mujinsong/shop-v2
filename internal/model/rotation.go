package model

import "github.com/gogf/gf/v2/frame/g"

// RotationCreateUpdateBase 创建/修改轮播图基类
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationCreateInput 创建轮播图
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建轮播图返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}
type RotationDeleteReq struct {
	g.Meta `path:"/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}
type RotationDeleteRes struct{}

// RotationUpdateInput 修改轮播图
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}
