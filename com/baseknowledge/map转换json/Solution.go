package main

import (
	"encoding/json"
	"fmt"
)

// map转换成json
func main01() {
	m := make(map[string]interface{})
	// map转成json不区分大小写
	m["subject"] = "go语言开发"
	m["student"] = []string{"memolei", "迈莫"}
	m["price"] = 240.00
	slice, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(string(slice))

}

func main() {
	slice := []byte(`{"Subject":"go语言开发","Students":["memolei","迈莫"],"Price":240}`)

	// 创建map[string]interface{}
	m := make(map[string]interface{})
	//var temp interface{}
	err := json.Unmarshal(slice, &m)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	// map的value为interface 需要进行类型断言 获取数据内容
	for k, v := range m {
		switch tv := v.(type) {
		//case interface{}:
		//	fmt.Println("interface类型数据", k, v)
		case string:
			fmt.Println("string类型数据：", k, tv)
		case float64:
			fmt.Println("float64类型数据：", k, tv)
		case []string:
			fmt.Println("[]string类型数据：", k, tv)
		case []interface{}:
			fmt.Println("[]interface类型数据：", k, tv)
			for i, val := range tv {
				fmt.Println(i, val)
			}
		}
	}
}
