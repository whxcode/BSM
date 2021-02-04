package controller

import (
	"React-Go/src/data"
	"React-Go/src/model"
	"net/http"
)

// 菜单查询
func QueryMenusList(w http.ResponseWriter, r *http.Request) {
	permission := []data.SinglePermission{}
	model.CommonList(w, r, &permission, `select * from menus`, `select count(id) from menus`)
}

// 菜单删除
func MenusDelete(w http.ResponseWriter, r *http.Request) {
	menuId := r.PostFormValue("menuId")
	model.Db.QueryRow(`delete from menus where id = ?`, menuId)
	model.Send(w, 0, "成功")
}

// 菜单添加
func MenuAdd(w http.ResponseWriter, r *http.Request) {
	/*menu := &data.SinglePermission{
		Icon: "0",
		PermissionName: "测试",
		PermissionId: "test",
		Actions: "[]",
	}*/
	id := r.PostFormValue("id")
	icon := r.PostFormValue("icon")
	permissionName := r.PostFormValue("permissionName")
	permissionId := r.PostFormValue("permissionId")
	actions := r.PostFormValue("actions")

	// fmt.Println(icon,permissionId,permissionName,actions)
	if id != "" {
		model.Db.QueryRowx("update menus set permission_name = ?,permission_id = ?,icon = ?,actions = ? where id = ?",
			permissionName, permissionId, icon, actions, id)

	} else {
		model.Db.QueryRowx(`insert into menus (permission_name,permission_id,icon,actions) values(?,?,?,?)`,
			permissionName, permissionId, icon, actions)
	}

	model.Send(w, 0, "操作成功")
	// fmt.Println(menu)

}

// 获取菜单详情
func MenuGet(w http.ResponseWriter, r *http.Request) {
	menuId := r.PostFormValue("menuId")
	menu := []data.SinglePermission{}
	model.Db.Select(&menu, `select * from menus where id = ?`, menuId)
	model.Send(w, 0, menu[0])
}
