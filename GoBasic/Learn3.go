package main

import "fmt"

/*
	多重赋值在Go语言的错误处理和函数返回值中会大量地使用。例如使用Go语言进行排序时就需要使用交换，代码如下：
	代码说明如下：
第 1 行，将 IntSlice 声明为 []int 类型。
第 3 行，为 IntSlice 类型编写一个 Len 方法，提供切片的长度。
第 4 行，根据提供的 i、j 元素索引，获取元素后进行比较，返回比较结果。
第 5 行，根据提供的 i、j 元素索引，交换两个元素的值。
*/
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	/*
		变量交换，传统方法
	*/
	var a int = 100
	var b int = 200
	var t int
	fmt.Println(a, b)
	t = a
	a = b
	b = t
	fmt.Println(a, b)

	/*
		变量交换，go多重赋值
	*/
	var c int = 300
	var d int = 400
	fmt.Println(c, d)
	d, c = c, d
	fmt.Println(c, d)

	/*
		多重赋值在Go语言的错误处理和函数返回值中会大量地使用。例如使用Go语言进行排序时就需要使用交换，代码如下：
	*/
}
