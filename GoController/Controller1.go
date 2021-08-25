package main

import "fmt"

func conn() {
	a := 1
	fmt.Println(a)
	return
}
func main() {
	/*
		Go语言if else（分支结构）
	*/
	condition := true
	condition2 := false
	//单分支操作if
	if condition {
		//do
	}
	//多分支操作if```else
	if condition {
		//do
	} else {
		//do
	}

	//多分支
	if condition {
		//do
	} else if condition2 {
		//do
	} else {
		//do
	}

	//特殊写法
	if err := conn; err != nil {
		return
	}
}
