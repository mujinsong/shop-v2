package model

// RoleCreateUpdateBase 创建/修改管理员基类
type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

// RoleCreateInput 创建管理员
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput 创建管理员返回结果
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}
