package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Subject string
	Students []string
	Price float64
}

// 结构体转换成json
func main01() {
	cl := Class{
		"go语言开发",
		[]string{"memolei", "迈莫"},
		240.00,
	}
	slice, err := json.Marshal(cl)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(string(slice))
}
// {"Subject":"go语言开发","Students":["memolei","迈莫"],"Price":240}

// json 转换成结构体
func main()  {
	slice := []byte(`{"Subject":"go语言开发","Students":["memolei","迈莫"],"Price":240}`)
	var cl Class
	err := json.Unmarshal(slice, &cl)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(cl)
}