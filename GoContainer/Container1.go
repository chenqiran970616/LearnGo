package main

import "fmt"

func main() {
	/*
		数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，所以在Go语言中很少直接使用数组。
		和数组对应的类型是 Slice（切片），Slice 是可以增长和收缩的动态序列，功能也更灵活，但是想要理解 slice 工作原理的话需要先理解数组
	*/

	/*
		数组的声明语法如下：
		var 数组变量名 [元素数量]Type
	*/

	/*
		语法说明如下所示：
		数组变量名：数组声明及使用时的变量名。
		元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，元素数量不能含有到运行时才能确认大小的数值。
		Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组。
	*/

	//数组的每个元素都可以通过索引下标来访问，索引下标的范围是从 0 开始到数组长度减 1 的位置，内置函数 len() 可以返回数组中元素的个数。
	var e [3]int             // 定义三个整数的数组
	fmt.Println(e[0])        // 打印第一个元素
	fmt.Println(e[len(e)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range e {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range e {
		fmt.Printf("%d\n", v)
	}

	//默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，对于数字类型来说就是 0，同时也可以使用数组字面值语法，用一组值来初始化数组：
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"
	fmt.Println(q[0]) // "1"

	//在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算，因此，上面数组 q 的定义可以简化为：
	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n", p) // "[3]int

	//数组的长度是数组类型的一个组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型，数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。
	//p = [4]int{1, 2, 3, 4} // 编译错误：无法将 [4]int 赋给 [3]int

	//比较两个数组是否相等
	//如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（==和!=）来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int

	//遍历数组——访问每一个数组元素
	//遍历数组也和遍历切片类似，代码如下所示：
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	for k, v := range team { //使用 for 循环，遍历 team 数组，遍历出的键 k 为数组的索引，值 v 为数组的每个元素值。
		fmt.Println(k, v)
	}

}
