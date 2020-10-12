package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var a int = 1024

	// 获取变量A的反射值对象
	ta := reflect.ValueOf(a)

	// 获取interface{} 类型的值 通过类型断言转换
	var ga = ta.Interface().(int)
	fmt.Println(ga)

	var columns []string
	aa, b, c := "1", "2", "3"
	columns = append(columns, fmt.Sprintf("%s %s %s", aa, b, c))
	fmt.Println(columns)
	desc := strings.Join(columns, ",")
	fmt.Println(desc)

}
