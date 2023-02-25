package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponCreateUpdateBase 创建/修改优惠券基类
type CouponCreateUpdateBase struct {
	Name       string
	Price      int
	GoodsIds   string
	CategoryId uint
}

// CouponCreateInput 创建优惠券
type CouponCreateInput struct {
	CouponCreateUpdateBase
}

// CouponCreateOutput 创建优惠券返回结果
type CouponCreateOutput struct {
	CouponId int `json:"coupon_id"`
}
type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"优惠券" summary:"删除优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的优惠券" dc:"优惠券id"`
}
type CouponDeleteRes struct{}

// CouponUpdateInput 修改优惠券
type CouponUpdateInput struct {
	CouponCreateUpdateBase
	Id uint
}

// CouponGetListInput 获取内容列表
type CouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)

}

// CouponGetListOutput 查询列表结果
type CouponGetListOutput struct {
	List  []CouponGetListOutputItem `json:"list" description:"列表"`
	Page  uint                      `json:"page" description:"分页码"`
	Size  uint                      `json:"size" description:"分页数量"`
	Total int64                     `json:"total" description:"数据总数"`
}

// CouponSearchInput 搜索列表
type CouponSearchInput struct {
	Key      string // 关键字
	CouponId uint   // 栏目ID
	Page     uint   // 分页号码
	Size     uint   // 分页数量，最大50
	Sort     int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponSearchOutput 搜索列表结果
type CouponSearchOutput struct {
	List  []CouponSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int           `json:"stats"` // 搜索统计
	Page  uint                     `json:"page"`  // 分页码
	Size  uint                     `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
}

type CouponGetListOutputItem struct {
	//Coupon *CouponListItem `json:"coupon"`
	Id         uint        `json:"id"` // 自增ID
	Price      int         `json:"price"`
	GoodsIds   string      `json:"goods_ids"`
	Name       string      `json:"link"`
	CategoryId uint        `json:"category_id"`
	CreatedAt  *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"` // 修改时间
}

type CouponSearchOutputItem struct {
	CouponGetListOutputItem
}

//// CouponListItem 主要用于列表展示
//type CouponListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
