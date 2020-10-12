package main

import (
	"fmt"
	"reflect"
)

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func main02() {
	type Cat struct {}
	// 获取结构体实例的反射类型对象
	ta := reflect.TypeOf(Cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(ta.Name(), ta.Kind())

	// 获取Zero常量的反射类型对象
	tb := reflect.TypeOf(Zero)
	fmt.Println(tb.Name(), tb.Kind())
}