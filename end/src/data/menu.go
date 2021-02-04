package data

type MenuHigh struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	MenuId   string `json:"menuId" db:"menu_id"`
	ParentId int    `json:"parentId" db:"parent_id"`
	Path     string `json:"path" db:"path"`
	Name     string `json:"name" db:"name"`
	Icon     string `json:"icon" db:"icon"`
}
type MenuHighWrap struct {
	*MenuHigh
	Children []*MenuHighWrap `json:"children"`
}
type MenuHighSlice []*MenuHigh

func generator(value *MenuHighWrap, tree []*MenuHighWrap) []*MenuHighWrap {
	menu := []*MenuHighWrap{}
	for _, v := range tree {
		if v.Id != value.Id && v.ParentId == value.Id {
			menu = append(menu, v)
		}
	}
	return menu
}
func parse(tree MenuHighSlice) []*MenuHighWrap {
	m := []*MenuHighWrap{}
	menu := []*MenuHighWrap{}
	for _, v := range tree {
		vv := &MenuHighWrap{
			v,
			[]*MenuHighWrap{},
		}
		m = append(m, vv)
	}
	for _, v := range m {
		v.Children = generator(v, m)
		if v.ParentId == 0 {
			menu = append(menu, v)
		}
	}
	return menu
}
func (m *MenuHighSlice) Do() []*MenuHighWrap {
	return parse(*m)
}
