package main

import "fmt"

func main() {
	/*
		Go语言append()为切片添加元素

	*/
	//Go语言的内建函数 append() 可以为切片动态添加元素，代码如下所示：
	var a []int
	a = append(a, 1) // 追加1个元素
	fmt.Println(a)
	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	fmt.Println(a)
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包，我们不能传递另一个切片作为参数，它必须是一个或多个 type 类型的参数。因此，我们将使用解包操作符 … 将切片分解到一系列参数中 ( 以满足*append*函数的参数要求)。
	fmt.Println(a)
	//不过需要注意的是，在使用 append() 函数为切片动态添加元素时，如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。
	fmt.Println("-----------------------------------------")

	//切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……，代码如下：
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	/*
		第 21 行，声明一个整型切片。
		第 23 行，循环向 numbers 切片中添加 10 个数。
		第 24 行，打印输出切片的长度、容量和指针变化，使用函数 len() 查看切片拥有的元素个数，使用函数 cap() 查看切片的容量情况。
		通过查看代码输出，可以发现一个有意思的规律：切片长度 len 并不等于切片的容量 cap。
	*/
	fmt.Println("-----------------------------------------")
	//在切片的开头添加元素：在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次，因此，从切片的开头添加元素的性能要比从尾部追加元素的性能差很多。
	var c = []int{1, 2, 3}
	fmt.Println("初始化c为")
	fmt.Println(c)
	c = append([]int{0}, a...) // 在开头添加1个元素
	fmt.Println("在c开头添加元素")
	fmt.Println(c)
	c = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println("在c开头添加一个切片")
	fmt.Println(c)
	fmt.Println("-----------------------------------------")

	//因为 append 函数返回新切片的特性，所以切片也支持链式操作，我们可以将多个 append 操作组合起来，实现在切片中间插入元素：
	var d = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("初始化d", d)
	d = append(d[:3], append([]int{6}, d[3:]...)...) // 在第i个位置插入x 第3个位置插入6
	fmt.Println("d插入6", d)
	d = append(d[:3], append([]int{1, 2, 3}, d[3:]...)...) // 在第i个位置插入切片 第3个位置插入1，2，3切片
	fmt.Println("d插入123", d)

}
