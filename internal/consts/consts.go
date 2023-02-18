package consts

const (
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GtokenAdminPrefix        = "Admin:"
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	CodeMissingParameterMsg  = "缺少参数"
	CacheModeRedis           = 2
	BackendServerName        = "shop_v2"
	MultiLogin               = true
	ErrLoginFaulMsg          = "登录失败，帐号或密码错误"
	GTokenExpireIn           = 10 * 24 * 60 * 60
)
