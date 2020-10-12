package main

import (
	"fmt"
	"go/ast"
	"reflect"
)

type Member struct {
	Id int `json:"id" orm:"member_id"`
	Name string `json:"name"`
	status bool
}

func (m *Member) SetStatus(s bool) {
	m.status = s
}

func (m *Member) GetStatus() bool {
	return m.status
}

func (m Member) String() string {
	return fmt.Sprintf("id: %d, name: %s", m.Id, m.Name)
}

// 如果是值传递 那么反射获取的对象将不能修改此值 否则panic
func Parse(v interface{})  {
	// 获取v的变量值 如果是地址传递获取的是指针 值传递则为变量值
	rValue := reflect.ValueOf(v)

	// 判断是否为一个指针 如果是指针通过Elem() 获取指针指向的变量值
	rValue = reflect.Indirect(rValue)

	// 获取V的变量类型
	rType := rValue.Type()

	switch rType.Kind() {
	case reflect.Struct:
		// 遍历结构体字段
		for i := 0; i < rType.NumField(); i++ {
			field := rType.Field(i)
			// 忽略匿名字段和私有字段
			if !field.Anonymous && ast.IsExported(field.Name) {
				// 获取结构体字段的interface{} 变量
				fmt.Println(rValue.Field(i).Interface())

				// 对于值传递的结构体使用Addr() 获取结构体字段 *interface{} 变量会报错
				fmt.Println(reflect.TypeOf(rValue.Field(i).Addr().Interface()))

				// 获取字段tag
				fmt.Println(rType.Field(i).Tag.Get("json"))
			}
		}

		// 根据字段名获取字段 interface{}
		fmt.Println(rValue.FieldByName("Name").Interface())

		// 获取非指针方法数量，案例中的 Member 的 String 方法
		fmt.Println(rValue.NumMethod())
		// 获取所有方法数量
		fmt.Println(rValue.Addr().NumMethod())
		
	}
}
