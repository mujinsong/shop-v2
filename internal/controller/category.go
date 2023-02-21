package controller

import (
	"context"

	"shop-v2/api/backend"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Category 分类管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *backend.CategoryReq) (res *backend.CategoryRes, err error) {
	out, err := service.Category().Create(ctx, model.CategoryCreateInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Sort:     req.Sort,
			Level:    req.Level,
			ParentId: req.ParentId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryRes{CategoryId: out.CategoryId}, nil
}
func (a *cCategory) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	return
}

func (a *cCategory) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, model.CategoryUpdateInput{
		Id: req.Id,
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Sort:     req.Sort,
			Level:    req.Level,
			ParentId: req.ParentId,
		},
	})
	return &backend.CategoryUpdateRes{Id: req.Id}, nil
}

func (a *cCategory) List(ctx context.Context, req *backend.CategoryGetListCommonReq) (res *backend.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
