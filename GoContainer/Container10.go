package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		Go语言遍历map（访问map中的每一个键值对）
	*/

	//for range
	scene := make(map[string]int)
	scene["apple"] = 10
	scene["banana"] = 20
	scene["cat"] = 30
	for k, v := range scene {
		fmt.Println(k, v)
	}
	fmt.Println("----------------------------------------------------------")

	//只遍历map的值
	for _, v := range scene {
		fmt.Println(v)
	}
	fmt.Println("----------------------------------------------------------")

	//只遍历map的键
	for k, _ := range scene {
		fmt.Println(k)
	}
	fmt.Println("----------------------------------------------------------")
	//注意：遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果。
	//获取特定顺序的遍历结果，需要先排序。
	scene2 := make(map[string]int)
	// 准备map数据
	scene2["route"] = 66
	scene2["brazil"] = 4
	scene2["china"] = 960
	// 声明一个切片保存map数据
	var sceneList []string
	// 将map数据遍历复制到切片中
	for k := range scene2 {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	// 输出
	fmt.Println(sceneList)
}
