package main

import "fmt"

func main() {
	/*
		在计算机中，复数是由两个浮点数表示的，其中一个表示实部（real），一个表示虚部（imag）。
		Go语言中复数的类型有两种，分别是  complex128（64 位实数和虚数）和 complex64（32 位实数和虚数），其中 complex128 为复数的默认类型。
		复数的值由三部分组成 RE + IMi，其中 RE 是实数部分，IM 是虚数部分，RE 和 IM 均为 float 类型，而最后的 i 是虚数单位。
	*/

	//声明复数的语法格式如下所示，其中 name 为复数的变量名，complex128 为复数的类型，“=”后面的 complex 为Go语言的内置函数用于为复数赋值，x、y 分别表示构成该复数的两个 float64 类型的数值，x 为实部，y 为虚部。
	x, y := 1, 2
	var name complex128 = complex(float64(x), float64(y))
	fmt.Println(name)
	//简写
	a, b := 3, 4
	name2 := complex(float64(a), float64(b))
	fmt.Println(name2)

	//对于一个复数z := complex(x, y)，可以通过Go语言的内置函数real(z) 来获得该复数的实部，也就是 x；通过imag(z) 获得该复数的虚部，也就是 y。
	fmt.Println(name * name2)
	fmt.Println(real(name * name2))
	fmt.Println(imag(name * name2))

}
