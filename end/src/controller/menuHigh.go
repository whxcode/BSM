package controller

import (
	"React-Go/src/data"
	"React-Go/src/model"
	"net/http"
	"strconv"
)

func QueryMenuHighList(w http.ResponseWriter, r *http.Request) {
	sql := `select * from menu_high`
	var m data.MenuHighSlice = []*data.MenuHigh{}
	model.Db.Select(&m, sql)
	mm := m.Do()
	model.Send(w, 0, &mm)
}

func MenuHighAdd(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	name := r.PostFormValue("name")
	path := r.PostFormValue("path")
	icon := r.PostFormValue("icon")
	menuId := r.PostFormValue("menuId")
	parentId := r.PostFormValue("parentId")
	pid := 0
	if parentId != "" {
		pid, _ = strconv.Atoi(parentId)
	}
	sql := `insert into menu_high(title,menu_id,parent_id,path,name,icon) values(?,?,?,?,?,?)`
	model.Db.QueryRowx(sql, title, menuId, pid, path, name, icon)
	model.Send(w, 0, "成功")
}

func deep(list []*data.MenuHigh) {
	sql := `select * from menu_high where parent_id = ?`
	data := []*data.MenuHigh{}
	for _, v := range list {
		model.Db.Select(&data, sql, v.Id)
		for _, vv := range data {
			model.Db.Query("delete from menu_high where id = ?", vv.Id)
		}
		deep(data)
	}
}

func MenuHighDel(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	sql := `select * from menu_high where id = ?`
	data := []*data.MenuHigh{}
	model.Db.Select(&data, sql, id)
	model.Db.Query("delete from menu_high where id = ?", id)
	go deep(data)
	model.Send(w, 0, "成功")
}
