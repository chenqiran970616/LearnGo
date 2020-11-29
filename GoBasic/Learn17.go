package main

import (
	"fmt"
	"math"
	"time"
)

//常量可以是构成类型的一部分，例如用于指定数组类型的长度
const IPv4Len = 4

// parseIPv4 解析一个 IPv4 地址 (d.d.d.d).
func parseIPv4(s string) [4]byte {
	var p [IPv4Len]byte
	fmt.Println(p)
	// ...
	return p
}

func main() {
	/*
		Go语言常量和const关键字
		Go语言中的常量使用关键字 const 定义，用于存储不会改变的数据，常量是在编译时被创建的，即使定义在函数内部也是如此，并且只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。
		由于编译时的限制，定义常量的表达式必须为能被编译器求值的常量表达式。
		常量的定义格式和变量的声明语法类似：const name [type] = value，例如：
	*/
	const pi = 3.14159 // 相当于 math.Pi 的近似值
	/*
		在Go语言中，你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
		显式类型定义： const b string = "abc"
		隐式类型定义： const b = "abc"

		常量的值必须是能够在编译时就能够确定的，可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。
		正确的做法：const c1 = 2/3
		错误的做法：const c2 = getNumber() // 引发构建错误: getNumber() 用做值
	*/

	//和变量声明一样，可以批量声明多个常量：
	const (
		e   = 2.7182818
		pi2 = 3.1415926
	)

	/*
		所有常量的运算都可以在编译期完成，这样不仅可以减少运行时的工作，也方便其他代码的编译优化，当操作数是常量时，一些运行时的错误也可以在编译时被发现，
		例如整数除零、字符串索引越界、任何导致无效浮点数的操作等。
		常量间的所有算术运算、逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数调用都是返回常量结果：len、cap、real、imag、complex 和 unsafe.Sizeof。
		因为它们的值是在编译期就确定的，因此常量可以是构成类型的一部分，例如用于指定数组类型的长度
	*/

	parseIPv4("100")

	/*
		一个常量的声明也可以包含一个类型和一个值，但是如果没有显式指明类型，那么将从右边的表达式推断类型。在下面的代码中，time.Duration 是一个命名类型，
		底层类型是 int64，time.Minute 是对应类型的常量。下面声明的两个常量都是 time.Duration 类型，可以通过 %T 参数打印类型信息：
	*/
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
	/*
		如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前面常量的初始化表达式，对应的常量类型也是一样的。例如：
	*/
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"

	/*
		iota 常量生成器
		常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个 const 声明语句中，在第一个声明的常量所在的行，
		iota 将会被置为 0，然后在每一个有常量声明的行加一。
	*/

	//【示例 1】首先定义一个 Weekday 命名类型，然后为一周的每天定义了一个常量，从周日 0 开始。在其它编程语言中，这种类型一般被称为枚举类型。
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	/*
		无类型常量
		Go语言的常量有个不同寻常之处。虽然一个常量可以有任意一个确定的基础类型，例如 int 或 float64，或者是类似 time.Duration 这样的基础类型，但是许多常量并没有一个明确的基础类型。
		编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，可以认为至少有 256bit 的运算精度。这里有六种未明确类型的常量类型，
		分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。
		通过延迟明确常量的具体类型，不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。
	*/

	//【示例 2】math.Pi 无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方：
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	/*
		如果 math.Pi 被确定为特定类型，比如 float64，那么结果精度可能会不一样，同时对于需要 float32 或 complex128 类型值的地方则需要一个明确的强制类型转换：
		对于常量面值，不同的写法可能会对应不同的类型。例如 0、0.0、0i 和 \u0000 虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常量类型。
		同样，true 和 false 也是无类型的布尔类型，字符串面值常量是无类型的字符串类型。
	*/
	const Pi64 float64 = math.Pi
	var u float32 = float32(Pi64)
	var v float64 = Pi64
	var w complex128 = complex128(Pi64)
	fmt.Println(u, v, w)

}
