package main

import (
	"React-Go/src/controller"
	"React-Go/src/data"
	"React-Go/src/model"
	"React-Go/src/util"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	handler := data.CreateServe()
	model.MyServe = handler
	serve := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	handler.AddPost("/login", controller.Login)                  // 用户登录
	handler.AddPost("/getUserInfo", controller.GetUesInfo)       // 获取用户权限
	handler.AddPost("/queryUsersList", controller.QueryUserList) // 获取用户列表
	handler.AddPost("/getUserSingleInfo", controller.GetUserSingleInfo)
	handler.AddPost("/userAdd", controller.UserAdd)
	handler.AddPost("/modifyBaseUser", controller.ModifyBaseUser)
	handler.AddPost("/realUser", controller.RealUser)

	handler.AddPost("/queryMenusList", controller.QueryMenusList) // 查询菜单
	handler.AddPost("/menusDel", controller.MenusDelete)          // 菜单删除
	handler.AddPost("/menuAdd", controller.MenuAdd)               // 菜单删除
	handler.AddPost("/menuGet", controller.MenuGet)               // 获取菜单详情

	handler.AddPost("/queryRolesList", controller.QueryRolesList)                 // 获取角色列表
	handler.AddPost("/roleModifyOperatorMenu", controller.RoleModifyOperatorMenu) // 修改角色下对某一个菜单的操作权限
	handler.AddPost("/roleGetQueryMenus", controller.RoleGetQueryMenus)           // 获取某个角色下面有那些菜单项
	handler.AddPost("/roleAdd", controller.RoleAdd)

	// 新的管理菜单
	handler.AddPost("/queryMenuHighList", controller.QueryMenuHighList)
	handler.AddPost("/menuHighAdd", controller.MenuHighAdd)
	handler.AddPost("/menuHighDel", controller.MenuHighDel)
	// 文件系统
	handler.AddPost("/uploadFile", controller.UploadFile)
	// 运行静态资源文件
	controller.RunAssetsServe(util.Public, ":8001")

	fmt.Println(serve.ListenAndServe())
}
