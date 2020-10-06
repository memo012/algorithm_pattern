package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("法师", 97)
	scene.Store("老郑", 100)
	scene.Store("兵哥", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("法师"))
	// 根据键删除对应的键值对
	scene.Delete("法师")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(key, value interface{}) bool {
		fmt.Println(key, "--", value)
		return true
	})
}
