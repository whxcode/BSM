package base

type array struct {
	data []interface{}
	len  int
}

func (a *array) Push(e interface{}) {
	a.data = append(a.data, e)
	a.len++
}
func (a *array) Pop() interface{} {
	l := a.len
	a.len--
	v := a.data[a.len]
	a.data = a.data[0 : l-1]
	return v
}
func (a *array) Len() int {
	return a.len
}
func (a *array) Del(pos int) interface{} {
	if pos > a.len || pos < 0 {
		return nil
	}
	v := a.data[pos]
	a.data = append(a.data[0:pos], a.data[pos+1:a.len]...)
	a.len--
	return v
}
func (a *array) Each(fn func(v interface{}, i int)) { // On
	for i, v := range a.data {
		fn(v, i)
	}
}
func (a *array) Find(fn func(v interface{}) bool) interface{} {
	for _, v := range a.data {
		if fn(v) {
			return v
		}
	}
	return nil
}

func (a *array) Get(pos int) interface{} {
	if pos > a.len || pos < 0 {
		return nil
	}
	return a.data[pos]
}

// other
func (a *array) Map(fn func(v interface{}, i int) interface{}) *array {
	if a.len == 0 {
		return nil
	}
	n := CreateArray()
	for i, v := range a.data {
		n.Push(fn(v, i))
	}
	return n
}

/**
* 创建一个动态数组
 */
func CreateArray() *array { // O1
	l := &array{}
	l.len = 0
	l.data = make([]interface{}, 0)
	return l
}
