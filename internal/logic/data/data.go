package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
	"shop-v2/utility"
)

type sData struct {
}

//注册服务
func init() {
	service.RegisterData(New())
}
func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: int(todayOrederCount(ctx)), //todo
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, nil
}

//查询今天的订单总数
func todayOrederCount(ctx context.Context) (count int64) {
	count, err := dao.OrderInfo.Ctx(ctx).WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
