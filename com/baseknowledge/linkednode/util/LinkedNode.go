package util

import (
	"fmt"
	"reflect"
)

type LinkedNode struct {
	Data interface{}
	Next *LinkedNode
}

// 创建链表
func (node *LinkedNode) Create(Data ...interface{}) {
	if node == nil {
		return
	}

	head := node

	for i := 0; i < len(Data); i++ {
		newNode := new(LinkedNode)
		newNode.Data = Data[i]
		newNode.Next = nil
		node.Next = newNode
		node = node.Next
	}
	node = head
}

// 打印链表
func (node *LinkedNode) Print() {
	if node == nil {
		return
	}
	if node.Data != nil {
		fmt.Println(node.Data, " ")
	}
	node.Next.Print()
}

// 链表长度
func (node *LinkedNode) Length() int {
	if node == nil {
		return -1
	}
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
	}
	return i
}

// 插入数据 -> 尾插
func (node *LinkedNode) InsertByHead(Data interface{}) {
	if node == nil {
		return
	}
	if Data == nil {
		return
	}
	// 创建新节点
	newNode := new(LinkedNode)
	newNode.Data = Data
	newNode.Next = node.Next
	// 将新节点放在当前节点后面
	node.Next = newNode
}

// 插入数据 -> 尾插
func (node *LinkedNode) InsertByTail(Data interface{}) {
	if node == nil {
		return
	}
	// 查找链表末尾节点
	for node.Next != nil {
		node = node.Next
	}

	// 创建新节点 将其插入末尾节点尾部
	newNode := new(LinkedNode)
	newNode.Data = Data
	newNode.Next = nil
	node.Next = newNode
}

// 插入数据(下标) 位置
func (node *LinkedNode) InsertByIndex(index int, Data interface{}) {
	if node == nil {
		return
	}
	if index < 0 {
		return
	}
	preNode := node

	for i := 0; i < index; i++ {
		preNode = node
		if node == nil {
			return
		}
		node = node.Next
	}

	// 创建新节点
	newNode := new(LinkedNode)
	newNode.Data = Data
	newNode.Next = node

	// 上一个节点链接当前节点
	preNode.Next = newNode
}

// 删除数据 (下标) 位置
func (node *LinkedNode) DeleteByIndex(index int) {
	if node == nil {
		return
	}
	if index < 0 {
		return
	}
	preNode := node
	for i := 0; i < index; i++ {
		preNode = node
		if node == nil {
			return
		}
		node = node.Next
	}

	// 将上一个指针域节点指向node的下一个节点
	preNode.Next = node.Next

	// 销毁当前节点
	node.Data = nil
	node.Next = nil
	node = nil
}

// 删除数据(数据)
func (node *LinkedNode) DeleteByData(Data interface{}) {
	if node == nil {
		return
	}
	if Data == nil {
		return
	}
	preNode := node
	for node.Next != nil {
		preNode = node
		node = node.Next
		// 判断interface 存储的数据类型是否相同
		// reflect.DeepEqual()
		if reflect.TypeOf(node.Data) == reflect.TypeOf(Data) &&
			node.Data == Data {
			preNode.Next = node.Next

			// 销毁
			node.Next = nil
			node.Data = nil
			node = nil

			// 如果添加return 表示删除第一个相同的数据
			// 如果不添加return 表示删除所有相同的数据
			return
		}
	}
}

// 查找数据(数据)
func (node *LinkedNode) Search(Data interface{}) int {
	if node == nil {
		return -1
	}
	if Data == nil {
		return -1
	}
	i := 0
	for node.Next != nil {
		i++
		if reflect.TypeOf(node.Data) == reflect.TypeOf(Data) &&
			node.Data == Data {
			return i - 1
		}
		node = node.Next
	}
	return -1
}

// 销毁链表
func (node *LinkedNode) Destroy() {
	if node == nil {
		return
	}
	node.Next.Destroy()

	node.Next = nil
	node.Data = nil
	node = nil
}
