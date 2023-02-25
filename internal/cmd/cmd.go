package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop-v2/internal/controller"
	"shop-v2/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "Go-shop",
		Usage: "learning",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//不需要登录的路由组绑定
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
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,
						controller.Upload,
						controller.Category,   //商品分类
						controller.Coupon,     //商品优惠券
						controller.UserCoupon, //用户优惠券
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
