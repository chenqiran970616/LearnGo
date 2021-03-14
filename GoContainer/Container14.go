package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*
		列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等。
		在Go语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作。
	*/

	/*
		list 的初始化有两种方法：分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的。
			1) 通过 container/list 包的 New() 函数初始化 list
			变量名 := list.New()

			2) 通过 var 关键字声明初始化 list
			var 变量名 list.List

			列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题，例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机。

	*/
	//初始化列表
	listA := list.New()
	var listB list.List
	fmt.Println(listB)

	//在列表中插入元素
	listA.PushBack("first")   //将 fist 字符串插入到列表的尾部
	listA.PushFront("second") //将 second 放入列表，此时，列表中已经存在 fist 元素，second 这个元素将被放在 fist 的前面。

	//从列表中删除元素
	// 尾部添加后保存元素句柄
	element := listA.PushBack("fist")
	// 在fist之后添加high
	listA.InsertAfter("high", element)
	// 在fist之前添加noon
	listA.InsertBefore("noon", element)
	// 使用
	listA.Remove(element)
}
