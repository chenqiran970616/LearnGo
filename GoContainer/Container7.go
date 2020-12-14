package main

import "fmt"

func main() {
	/*
		切片其实就是多个相同类型元素的连续集合，既然切片是一个集合，那么我们就可以迭代其中的元素，
		Go语言有个特殊的关键字 range，它可以配合关键字 for 来迭代切片里的每一个元素
	*/
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	//当迭代切片时，关键字 range 会返回两个值，第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本
	fmt.Println("--------------------------------------------------------------")

	//range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	// 创建一个整型切片，并赋值
	slice2 := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice2 {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
	//因为迭代返回的变量是一个在迭代过程中根据切片依次赋值的新变量，所以 value 的地址总是相同的，要想获取每个元素的地址，
	//需要使用切片变量和索引值（例如上面代码中的 &slice[index]）。
	fmt.Println("--------------------------------------------------------------")

	//如果不需要索引值，也可以使用空白标识符（下划线）_来忽略这个值，
	// 创建一个整型切片，并赋值
	slice3 := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示其值
	for _, value := range slice3 {
		fmt.Printf("Value: %d\n", value)
	}
	fmt.Println("--------------------------------------------------------------")

	//关键字 range 总是会从切片头部开始迭代。如果想对迭代做更多的控制，则可以使用传统的 for 循环，代码如下所示。
	// 创建一个整型切片，并赋值
	slice4 := []int{10, 20, 30, 40}
	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice4); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice4[index])
	}

}
