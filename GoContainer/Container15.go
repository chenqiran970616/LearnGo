package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		Go语言nil：空值/零值
		在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""，而指针、切片、映射、通道、函数和接口的零值则是 nil。
		nil 标识符是不能比较的
		nil 不是关键字或保留字
		不同类型 nil 的指针是一样的
		不同类型的 nil 值占用的内存大小可能是不一样的
	*/
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8
	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8
	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8
	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16

}
