package permission

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/model/entity"
	"shop-v2/internal/service"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	//插入数据返回id
	lastInsertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	// 删除内容
	_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
		dao.PermissionInfo.Columns().Id: id,
	}).Unscoped().Delete()
	return err
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	//更新操作
	_, err := dao.PermissionInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.PermissionInfo.Columns().Id).
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	return err

}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	var (
		m = dao.PermissionInfo.Ctx(ctx)
	)
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.PermissionInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Permission
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
