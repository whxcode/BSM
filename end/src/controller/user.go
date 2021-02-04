package controller

import (
	"React-Go/src/data"
	"React-Go/src/model"
	"React-Go/src/util"
	"net/http"
)

/*
* 用户登录
 */
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	type Login struct {
		*data.User
		UserInfo *data.UserInfo `json:"userInfo"`
	}
	user := &Login{
		UserInfo: &data.UserInfo{},
	}
	model.Db.Get(user, "select * from users where account = ? and password = ?", username, password)
	model.Db.QueryRowx(`select * from user_info ui where ui.user_id = ?`, user.Id).StructScan(user.UserInfo)
	if user.Account != username {
		model.Send(w, 1, "未能找到该用户")
		return
	}
	model.Send(w, 0, user)
}

/*
* 用户登录
 */
func RealUser(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	type Login struct {
		*data.User
		UserInfo *data.UserInfo `json:"userInfo"`
	}
	user := &Login{
		UserInfo: &data.UserInfo{},
	}
	model.Db.Get(user, "select * from users where id = ?", id)

	model.Db.QueryRowx(`select * from user_info ui where ui.user_id = ?`, user.Id).StructScan(user.UserInfo)
	model.Send(w, 0, user)
}

/**
* 获取用户信息、包括访问菜单的权限
*
*username := r.PostFormValue("username")KJOIUYTRTYUIOP[654321234O9`
 */
func GetUesInfo(w http.ResponseWriter, r *http.Request) {
	role := &data.Role{}
	roleId := r.PostFormValue("roleId")
	model.Db.Get(role, "select * from roles where id = ?", roleId)
	permission := []*data.SinglePermission{}
	model.Db.Select(&permission, `SELECT 
       m.id id,
	permission_id,
	permission_name,
	icon
 FROM role_menu  rm left JOIN menus m on rm.menu_id = m.id  where rm.role_id = ?`, roleId)
	compare := &data.Permission{
		Role:        role,
		Permissions: permission,
	}

	per := &data.UserPermission{
		Role: compare,
	}

	for _, v := range permission {
		v.ActionsObject = []*data.Action{}
		model.Db.Select(&v.ActionsObject, `select * from role_actions r where r.menu_id = ? and r.role_id = ?`, v.Id, roleId)
	}
	model.Send(w, 0, per, "成功")
}

func QueryUserList(w http.ResponseWriter, r *http.Request) {
	users := []struct {
		data.User
		RoleName string `json:"roleName" db:"role_name"`
	}{}
	sql := `select 
	u.id id,account,username,u.role_id role_id,u.create_time create_time,u.update_time update_time,r.label role_name
 	from users u left join roles r on u.role_id = r.id`
	model.CommonList(w, r, &users, sql, `select count(id) from users`)
}

// 获取单个用户信息
func GetUserSingleInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.PostFormValue("id")
	users := []struct {
		data.User
		RoleName string `json:"roleName" db:"role_name"`
	}{}
	sql := `select 
	u.id id,account,username,u.role_id role_id,u.create_time create_time,u.update_time update_time,r.label role_name
 	from users u left join roles r on u.role_id = r.id where u.id = ?`
	model.Db.Select(&users, sql, userId)
	model.Send(w, 0, users[0])
}

// 添加&修改用户
func UserAdd(w http.ResponseWriter, r *http.Request) {
	userId := r.PostFormValue("id")
	userName := r.PostFormValue("userName")
	roleId := r.PostFormValue("roleId")
	password := r.PostFormValue("password")
	account := r.PostFormValue("account")
	code := 0
	var status interface{}
	if userId != "" {
		// 更新用户操作
		sql := `update users set username = ?,role_id = ? where id = ?`
		_, status = model.Db.Query(sql, userName, roleId, userId)
	} else {
		// 添加用户
		_, status = model.Db.Query(`insert into users(username,password,account,role_id) values (?,?,?,?)`,
			userName, password, account, roleId)
	}

	if status != nil {
		code = 1
	}

	model.Send(w, code, status)
}

func ModifyBaseUser(w http.ResponseWriter, r *http.Request) {
	m := util.GetQuestValue(r, "id", "brief", "email", "surname", "avatar")
	sql := `update user_info set avatar = ?,brief = ?,email = ?,surname = ? where id = ?`
	model.Db.QueryRowx(sql, m["avatar"], m["brief"], m["email"], m["surname"], m["id"])
	model.Send(w, 0, "修改成功")

}
