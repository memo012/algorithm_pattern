package main

import (
	"algorithm_pattern/com/baseknowledge/linkednode/util"
	"fmt"
)

func main() {
	node := new(util.LinkedNode)
	for i := 0; i < 5; i++ {
		node.InsertByHead(i)
	}
	node.Print()
	fmt.Println(node.Length())
}
