package main

import "fmt"

func main() {
	/*
		Go语言中允许使用多维数组，因为数组属于值类型，所以多维数组的所有维度都会在创建时自动初始化零值，多维数组尤其适合管理具有父子关系或者与坐标系相关联的数据。

		声明多维数组的语法如下所示：
		var array_name [size1][size2]...[sizen] array_type

		其中，array_name 为数组的名字，array_type 为数组的类型，size1、size2 等等为数组每一维度的长度。
	*/
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	fmt.Println(array)
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array)
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println(array)
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array)

	//【示例 2】为二维数组的每个元素赋值
	// 声明一个 2×2 的二维整型数组
	var array2 [2][2]int
	// 设置每个元素的整型值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	fmt.Println(array2)

	//【示例 3】同样类型的多维数组赋值
	// 声明两个二维整型数组
	var array3 [2][2]int
	var array4 [2][2]int
	// 为array2的每个元素赋值
	array4[0][0] = 10
	array4[0][1] = 20
	array4[1][0] = 30
	array4[1][1] = 40
	// 将 array2 的值复制给 array1
	array3 = array4
	fmt.Println(array3)
	fmt.Println(array4)

	//【示例 4】使用索引为多维数组赋值
	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var array5 [2]int = array4[1]
	fmt.Println(array5)
	// 将数组中指定的整型值复制到新的整型变量里
	var value int = array4[1][0]
	fmt.Println(value)

}
