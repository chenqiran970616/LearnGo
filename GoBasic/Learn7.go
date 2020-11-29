package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Go语言浮点类型（小数类型）
	    一个 float32 类型的浮点数可以提供大约 6 个十进制数的精度，而 float64 则可以提供约 15 个十进制数的精度，
		通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数并不是很大。
	*/
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	/*
		浮点数在声明的时候可以只写整数部分或者小数部分，像下面这样：
	*/
	const a = .71828 // 0.71828
	const b = 1.     // 1
	fmt.Println(a)
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)

	/*
		用 Printf 函数打印浮点数时可以使用“%f”来控制保留几位小数
	*/
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
}
