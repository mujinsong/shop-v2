package backend

import "github.com/gogf/gf/v2/frame/g"

type UserCouponReq struct {
	g.Meta `path:"/user/coupon/add" tags:"UserCoupon" method:"post" summary:"添加用户优惠券"`
	UserCouponCommonAddUpdate
}

type UserCouponCommonAddUpdate struct {
	UserId   uint  `json:"user_id" v:"require#用户ID必填" dc:"用户ID"`
	CouponId uint  `json:"id" v:"require#优惠券ID必填" `
	Status   uint8 `json:"status" dc:"状态"`
}

type UserCouponRes struct {
	Id uint `json:"id"`
}
type UserCouponDeleteReq struct {
	g.Meta `path:"/user/coupon/delete" method:"delete" tags:"用户优惠券" summary:"删除用户优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}
type UserCouponDeleteRes struct {
	Id uint `json:"id"`
}

type UserCouponUpdateReq struct {
	g.Meta `path:"/user/coupon/update/" method:"post" tags:"商品分类" summary:"商品分类接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品分类" dc:"商品分类Id"`
	UserCouponCommonAddUpdate
}
type UserCouponUpdateRes struct {
	Id uint `json:"id"`
}

type UserCouponGetListCommonReq struct {
	g.Meta `path:"/user/coupon/list" method:"get" tags:"商品分类" summary:"商品分类列表"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  uint        `json:"page" description:"分页码"`
	Size  uint        `json:"size" description:"分页数量"`
	Total int64       `json:"total" description:"数据总数"`
}
