package controller

import (
	"React-Go/src/data"
	"React-Go/src/model"
	"encoding/json"
	"net/http"
	"strconv"
)

// 查询角色列表
func QueryRolesList(w http.ResponseWriter, r *http.Request) {
	sql := `select r.id id,r.label label,r.create_time create_time,r.update_time update_time from roles r`
	type Action struct {
		Id       int    `json:"id" db:"id"`
		MenuId   int    `json:"menuId" db:"menu_id"`
		RoleId   int    `json:"roleId" db:"role_id"`
		Operator string `json:"operator" db:"operator"`
	}
	type Operator struct {
		Id                   int       `json:"id" db:"id"`
		MenuId               int       `json:"menuId" db:"menu_id"`
		RoleId               int       `json:"roleId" db:"role_id"`
		OperatorId           int       `json:"operatorId" db:"operator_id"`
		MenuName             string    `json:"menuName" db:"menu_name"`
		Actions              string    `json:"actions" db:"actions"`
		AlreadyActionsObject []*Action `json:"alreadyActionsObject"`
		AlreadyActions       string    `json:"alreadyActions"`
	}
	type Role struct {
		data.Role
		Actions   string      `json:"actions" db:"actions"`
		Operators []*Operator `json:"operators"`
	}
	roles := []*Role{}

	model.CommonList(w, r, &roles, sql, `select r.id id,r.label label,r.create_time create_time,r.update_time update_time from roles r`, func(data interface{}) interface{} {
		roles := data.(*[]*Role)

		for _, v := range *roles {
			v.Operators = []*Operator{}
			innerSql := `select
						m.id id,m.id menu_id,rm.role_id role_id,m.permission_name menu_name,m.actions actions
						from menus m left join role_menu rm on m.id = rm.menu_id  where rm.role_id = ?`
			model.Db.Select(&v.Operators, innerSql, v.Id)
		}

		for _, v := range *roles {
			for _, vv := range v.Operators {
				sql := `select * from role_actions ra WHERE ra.role_id = ? and ra.menu_id = ?`
				vv.AlreadyActionsObject = []*Action{}
				model.Db.Select(&vv.AlreadyActionsObject, sql, v.Id, vv.MenuId)
			}
		}

		return roles
	})
}

// 删除角色
func RoleDel(w http.ResponseWriter, r *http.Request) {
	roleId, _ := strconv.Atoi(r.PostFormValue("roleId"))
	if roleId == 1 {
		model.Send(w, 2, "不能删除该角色")
		return
	}
	sql := `delete from roles where id = ?`
	model.Db.Query(sql, roleId)
	model.Send(w, 0, "成功")
}

// 获取角色详情
func RoleGet(w http.ResponseWriter, r *http.Request) {
	roleId, _ := strconv.Atoi(r.PostFormValue("roleId"))
	sql := `select * from roles where id = ?`
	role := &struct {
		data.Role
		Actions string `json:"actions" db:"actions"`
	}{}
	model.Db.Select(&role, sql, roleId)
	model.Send(w, 0, &role)
}

// 添加角色
func RoleAdd(w http.ResponseWriter, r *http.Request) {
	roleId := r.PostFormValue("id")
	label := r.PostFormValue("label")
	actions := r.PostFormValue("actions")
	if roleId != "" {

		// 更新操作
		model.Db.Query(`update roles set label = ? where id = ?`, label, roleId)
		model.Db.Query(`delete from role_menu where role_id = ?`, roleId)

	} else {
		// 添加新角色
		model.Db.Query(`insert into roles(label) values(?)`, label)
	}

	sliceActions := []int{}
	json.Unmarshal([]byte(actions), &sliceActions)
	// 更新菜单列表
	for _, v := range sliceActions {
		model.Db.Query(`insert into role_menu(role_id,menu_id) values(?,?)`, roleId, v)
	}

	model.Send(w, 0, "成功")
}

// 修改角色下对某一个菜单的操作权限
func RoleModifyOperatorMenu(w http.ResponseWriter, r *http.Request) {
	operatorId := r.PostFormValue("operatorId")
	roleId := r.PostFormValue("roleId")
	menuId := r.PostFormValue("menuId")
	sql := `select id from role_actions ra where ra.menu_id = ? and ra.role_id = ? and ra.operator = ?`
	id := -1
	row, _ := model.Db.Query(sql, menuId, roleId, operatorId)
	defer row.Close()
	for row.Next() {
		row.Scan(&id)
	}
	if id != -1 {
		sql := `delete from role_actions where id = ?`
		model.Db.Query(sql, id)
	} else {
		sql := `insert into role_actions(role_id,menu_id,operator) values(?,?,?)`
		model.Db.Query(sql, roleId, menuId, operatorId)
	}
	// model.Db.Query(`update role_action set actions = ? where id = ?`,actions,operatorId)
	model.Send(w, 0, "修改成功")
}

// 查询一个角色可以访问那些(顶级)菜单
func RoleGetQueryMenus(w http.ResponseWriter, r *http.Request) {
	type Menus struct {
		Id           int    `json:"id" db:"id"`
		MenuName     string `json:"menuName" db:"menu_name"`
		PermissionId string `json:"permissionId" db:"permission_id"`
		RoleId       int    `json:"roleId" db:"role_id"`
	}
	type SpecialRole struct {
		data.Role
		Menus []Menus `json:"menus"`
	}
	roleId := r.PostFormValue("roleId")
	sql := `select * from roles where id = ?`
	sqlInner := `select 
	m.id id,m.permission_name menu_name,m.permission_id permission_id,rm.role_id role_id
 	from menus m left JOIN role_menu rm on m.id = rm.menu_id WHERE rm.role_id = ?`
	data := []SpecialRole{}
	model.Db.Select(&data, sql, roleId)
	data[0].Menus = []Menus{}
	model.Db.Select(&data[0].Menus, sqlInner, roleId)

	model.Send(w, 0, &data[0])
}
