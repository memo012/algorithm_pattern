package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
}

// 解析字符串
func stringLiteralTemplate() {
	s := "My name is {{.Name}}. I am {{.Age}} years old.\n"
	// 创建模板 并解析s这个字符串
	t, err := template.New("test").Parse(s)
	if err != nil {
		log.Fatal("Parse string literal template error:", err)
	}
	u := User{Name: "memolei", Age: 18}
	// 去执行渲染(把数据塞进模板文件中) 动作
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute string literal template error:", err)
	}
}

// 解析文件
func fileTemplate()  {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("Parse file template error:", err)
	}
	u := User{Name: "ml",Age: 19}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute file template error:", err)
	}
}

func main() {
	stringLiteralTemplate()
	//fileTemplate()
}
