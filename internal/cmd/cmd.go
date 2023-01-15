package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/api/backend"
	"shop-v2/internal/consts"
	"shop-v2/internal/controller"
	"shop-v2/internal/dao"
	"shop-v2/internal/model/entity"
	"shop-v2/internal/service"
	"shop-v2/utility"
	"shop-v2/utility/response"
	"strconv"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			//loginFunc := Login
			// 启动gtoken
			gfAdminToken := &gtoken.GfToken{
				ServerName:       "shop_v2",
				CacheMode:        2, //gredis
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  loginFunc,
				LoginAfterFunc:   loginAfterFunc,
				LogoutPath:       "/backend/user/logout",
				AuthPaths:        g.SliceStr{"/backend/admin/info"},
				AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				MultiLogin:       true,
				AuthAfterFunc:    authAfterFunc,
			}
			//todo 抽取方法
			//err = gfAdminToken.Start()
			//if err != nil {
			//	return err
			//}
			// 认证接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//gtoken 中间件绑定
				//err := gfToken.Middleware(ctx, group)
				//if err != nil {
				//	panic(err)
				//}
				group.Bind(
					controller.Hello,
					controller.Rotation,     //轮播图
					controller.Position,     //手工位
					controller.Admin.Create, //管理员
					controller.Admin.Update, //管理员
					controller.Admin.Delete, //管理员
					controller.Admin.List,   //管理员
					controller.Login,        //登陆
					controller.Role,         //角色
					controller.Data,         //数据大屏相关
					controller.Permission,   //权限

				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth)//for jwt
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,
						controller.Upload,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)

//todo 迁移合适位置
func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()
	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	//验证帐号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}

	// 唯一标识，扩展参数user data
	//g.Dump("admininfo", adminInfo)
	return consts.GtokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

//todo 迁移合适位置
//自定义登陆之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	//g.Dump("respData:", respData)
	g.Dump("在这了")
	if !respData.Success() {
		//g.Dump("在这了")
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		//adminId := respData.GetString("userKey")
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GtokenAdminPrefix)
		//g.Dump("ID:", adminId)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		//g.Dump(adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    5, //10 * 24 * 60 * 60, //单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		//g.Dump("data:", data)
		response.JsonExit(r, 0, "", data)
	}
	return
}
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var adminInfo entity.AdminInfo
	//g.Dump("res:", respData)
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		g.Dump("验证出问题了", err)
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}

//func LoginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
//	//g.Dump("respData:", respData)
//	if !respData.Success() {
//		respData.Code = 0
//		r.Response.WriteJson(respData)
//	} else {
//		respData.Code = 1
//		//获得登录用户id
//		userKey := respData.GetString("userKey")
//		adminId := gstr.StrEx(userKey, consts.GtokenAdminPrefix)
//		//g.Dump("adminId:", adminId)
//		//根据id获得登录用户其他信息
//		adminInfo := entity.AdminInfo{}
//		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
//		if err != nil {
//			return
//		}
//		//通过角色查询权限
//		//先通过角色查询权限id
//		var rolePermissionInfos []entity.RolePermissionInfo
//		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
//		if err != nil {
//			return
//		}
//		permissionIds := g.Slice{}
//		for _, info := range rolePermissionInfos {
//			permissionIds = append(permissionIds, info.PermissionId)
//		}
//
//		var permissions []entity.PermissionInfo
//		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
//		if err != nil {
//			return
//		}
//		data := &backend.LoginRes{
//			Type:        "Bearer",
//			Token:       respData.GetString("token"),
//			ExpireIn:    10 * 24 * 60 * 60, //单位秒,todo 根据实际情况修改
//			IsAdmin:     adminInfo.IsAdmin,
//			RoleIds:     adminInfo.RoleIds,
//			Permissions: permissions,
//		}
//		response.JsonExit(r, 0, "", data) //todo 替换成相同的方法
//	}
//	return
//}
