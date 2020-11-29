package main

import (
	"fmt"
	"net"
)

func main() {
	/*
		变量初始化的标准格式
		var 变量名 类型 = 表达式
	*/
	var a int = 100
	fmt.Println(a)
	/*
		在标准格式的基础上，将 int 省略后，编译器会尝试根据等号右边的表达式推导 hp 变量的类型。
	*/
	var b = 100
	fmt.Println(b)
	/*
		下面是编译器根据右值推导变量类型完成初始化的例子。
		第 1 和 2 行，右值为整型，attack 和 defence 变量的类型为 int。
		第 3 行，表达式的右值中使用了 0.17。由于Go语言和C语言一样，编译器会尽量提高精确度，以避免计算中的精度损失。所以这里如果不指定 damageRate 变量的类型，Go语言编译器会将 damageRate 类型推导为 float64，我们这里不需要 float64 的精度，所以需要强制指定类型为 float32。
		第 4 行，将 attack 和 defence 相减后的数值结果依然为整型，使用 float32() 将结果转换为 float32 类型，再与 float32 类型的 damageRate 相乘后，damage 类型也是 float32 类型。
	*/
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)
	/*
		短变量声明并初始化
		这是Go语言的推导声明写法，编译器会自动根据右值类型推断出左值的对应类型。
		注意：由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。
	*/
	c := 100
	fmt.Println(c)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		//错误处理
		fmt.Println(conn)
	}

}
