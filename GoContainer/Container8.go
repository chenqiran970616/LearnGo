package main

import "fmt"

func main() {
	/*
		多维切片

		Go语言中同样允许使用多维切片，声明一个多维数组的语法格式如下：
		var sliceName [][]...[]sliceType

		其中，sliceName 为切片的名字，sliceType为切片的类型，每个[ ]代表着一个维度，切片有几个维度就需要几个[ ]。
	*/

	//下面以二维切片为例，声明一个二维切片并赋值，代码如下所示。
	//声明一个二维切片
	var slice [][]int
	//为二维切片赋值
	slice = [][]int{{10}, {100, 200}}
	fmt.Println(slice)
	fmt.Println("--------------------------------------------------------------")

	// 上面的例子简写，声明一个二维整型切片并赋值，外层的切片包括两个元素，每个元素都是一个切片，
	// 第一个元素中的切片使用单个整数 10 来初始化，第二个元素中的切片包括两个整数，即 100 和 200。
	slice2 := [][]int{{10}, {100, 200}}
	fmt.Println(slice2)
	// 使用append，为第一个切片追加值为 20 的元素
	// Go语言里使用 append() 函数处理追加的方式很简明，先增长切片，再将新的整型切片赋值给外层切片的第一个元素，当上面代码中的操作完成后，再将切片复制到外层切片的索引为 0 的元素
	slice2[0] = append(slice[0], 20)
	fmt.Println(slice2)

}
