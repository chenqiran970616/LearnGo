package main

import "fmt"

func main() {
	/*
		Go语言map元素的删除和清空
	*/

	//使用 delete() 函数从 map 中删除键值对，使用 delete() 内建函数从 map 中删除一组键值对，delete() 函数的格式如下：
	//delete(map, 键)
	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}
	fmt.Println("----------------------------------------------------------")

}
