package data

import (
	"net/http"
)

/**
* 用户表
 */
type User struct {
	Id         int    `json:"id" db:"id"`
	UserName   string `json:"userName" db:"username"`
	Account    string `json:"account" db:"account"`
	Password   string `json:"password" db:"password"`
	CreateTime string `json:"createTime" db:"create_time"`
	UpdateTime string `json:"updateTime" db:"update_time"`
	RoleId     int    `json:"roleId" db:"role_id"`
}

/**
* 用户详情表
 */
type UserInfo struct {
	Id      int    `json:"id" db:"id"`
	Surname string `json:"surname" db:"surname"`
	Brief   string `json:"brief" db:"brief"`
	Email   string `json:"email" db:"email"`
	Avatar  string `json:"avatar" db:"avatar" `
	UserId  int    `json:"userId" db:"user_id"`
}

/*
* 权限表
 */
type Role struct {
	Id         int    `json:"id" db:"id"`
	Label      string `json:"label" db:"label"`
	CreateTime string `json:"createTime" db:"create_time"`
	UpdateTime string `json:"updateTime" db:"update_time"`
}

/**
* 菜单表
**/
type Menu struct {
	Id          int    `json:"id"`
	Label       string `json:"label"`
	Controllers string `json:"controllers"`
	ParentId    string `json:"parentId"`
	Icon        string `json:"icon"`
}

/*
* 服务器
**/
type MyServe struct {
	W    http.ResponseWriter
	R    *http.Request
	post map[string]http.HandlerFunc
	get  map[string]http.HandlerFunc
}

type SingleActionEntitySet struct {
	Action       string `json:"action"`
	Describe     string `json:"describe"`
	DefaultCheck bool   `json:"defaultCheck"`
}

type Action struct {
	Id       int    `json:"id" db:"id"`
	MenuId   int    `json:"menuId" db:"menu_id"`
	RoleId   int    `json:"roleId" db:"role_id"`
	Operator string `json:"operator" db:"operator"`
}
type SinglePermission struct {
	Id             int       `json:"id" db:"id""`
	PermissionId   string    `json:"permissionId" db:"permission_id"`
	PermissionName string    `json:"permissionName" db:"permission_name"`
	ActionsObject  []*Action `json:"actionsObject" db:"actionsObject"`
	Actions        string    `json:"actions" db:"actions"`
	Icon           string    `json:"icon" db:"icon"`
}

/**
* 权限
 */
type Permission struct {
	*Role
	Permissions []*SinglePermission `json:"permissions"`
}

/**
* 用户权限表
**/
type UserPermission struct {
	Role *Permission `json:"role"`
}

/**
* 公用列表
 */
type CommonList struct {
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
	Data       interface{} `json:"data"`
}

/**
* 公用列表
 */
type CommonListWrap struct {
	Result *CommonList `json:"result"`
}
