package article

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/model/entity"
	"shop-v2/internal/service"
)

type sArticle struct{}

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

func (s *sArticle) Create(ctx context.Context, in model.ArticleCreateInput) (out model.ArticleCreateOutput, err error) {
	lastInsertID, err := dao.ArticleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ArticleCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sArticle) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.ArticleInfo.Ctx(ctx).Where(g.Map{
		dao.ArticleInfo.Columns().Id: id,
	}).Delete()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sArticle) Update(ctx context.Context, in model.ArticleUpdateInput) error {
	_, err := dao.ArticleInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.ArticleInfo.Columns().Id).
		Where(dao.ArticleInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sArticle) GetList(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error) {
	var (
		m = dao.ArticleInfo.Ctx(ctx)
	)
	out = &model.ArticleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.ArticleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	temp, err := m.Count()
	if err != nil {
		return out, err
	}
	out.Total = int(temp)
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
