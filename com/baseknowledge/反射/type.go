package main


import (
	"fmt"
	"reflect"
)

func main01() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Printf("类型名:%v  种类:%v", typeOfA.Name(), typeOfA.Kind())
}
