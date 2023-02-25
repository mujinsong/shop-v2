package address

import (
	"context"
	"shop-v2/api/backend"
	"shop-v2/internal/consts"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/model/entity"
	"shop-v2/internal/service"
)

type sAddress struct{}

func init() {
	service.RegisterAddress(&sAddress{})
}

func (*sAddress) Add(ctx context.Context, in model.AddAddressInput) (out *model.AddAddressOutput, err error) {
	id, err := dao.AddressInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.AddAddressOutput{Id: int(id)}, err
}

func (*sAddress) Update(ctx context.Context, in model.UpdateAddressInput) (err error) {
	if _, err = dao.AddressInfo.Ctx(ctx).Data(in).FieldsEx(in.Id).Where(dao.AddressInfo.Columns().Id, in.Id).Update(); err != nil {
		return err
	}
	return nil
}

func (*sAddress) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.AddressInfo.Ctx(ctx).Where(dao.AddressInfo.Columns().Id, id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (*sAddress) Page(ctx context.Context, in model.PageAddressInput) (out *model.PageAddressOutput, err error) {
	var m = dao.AddressInfo.Ctx(ctx)
	out = &model.PageAddressOutput{
		CommonPaginationRes: backend.CommonPaginationRes{
			Page: in.Page,
			Size: in.Size,
			List: []entity.AddressInfo{},
		},
	}
	listModel := m.Page(in.Page, in.Size)

	temp, err := listModel.Count()
	if err != nil {
		return out, err
	}
	out.Total = int(temp)
	if out.Total == 0 {
		return out, nil
	}
	var list []entity.AddressInfo
	if err = listModel.ScanList(&list, "list"); err != nil {
		return out, err
	}
	if len(list) == 0 {
		return out, err
	}
	out.List = list

	return
}

// 客户端获取省市县区地址
func (*sAddress) GetCityList(ctx context.Context) (out *model.CityAddressListOutput, err error) {
	out = &model.CityAddressListOutput{}
	err = dao.AddressInfo.Ctx(ctx).Where(dao.AddressInfo.Columns().Pid, consts.ProvincePid).WithAll().Scan(&out.List)
	if err != nil {
		return out, err
	}
	return
}
