package controller

import (
	"context"
	"shop-v2/api/backend"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Login 登录管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	res = &backend.LoginDoRes{}

	err = service.Login().Login(ctx, model.UserLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	// 识别并跳转到登录前页面
	res.Info = service.Session().GetUser(ctx)

	return
}
