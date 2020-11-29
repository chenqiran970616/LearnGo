package main

import "fmt"

func main() {
	/*
		变量的声明
		var name type
		var 是声明变量的关键字，name 是变量名，type 是变量的类型。
		当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。所有的内存在 Go 中都是经过初始化的。
	*/
	var a int
	fmt.Println(a)
	var b float32
	fmt.Println(b)
	var c bool
	fmt.Println(c)
	var d string
	fmt.Println(d)
	var e *int
	fmt.Println(e)

	/*
		批量格式
		使用关键字 var 和括号，可以将一组变量定义放在一起
	*/
	var (
		a1 int
		b1 string
		c1 []float32
		d1 func() bool
		e1 struct {
			x int
		}
	)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println(e1)

	/*
		简短格式
		名字 := 表达式
	*/
	x := 100
	y, z := 1, "abc"
	fmt.Println(x, y, z)

}
