package main

import "fmt"

/*
如果需要经常做类似的转换，可以将转换的代码封装成一个函数
*/
// 如果b为真，btoi返回1；如果为假，btoi返回0
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
func main() {
	/*
		Go语言bool类型（布尔类型）
		一个布尔类型的值只有两种：true 或 false。if 和 for 语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。
		一元操作符!对应逻辑非操作，因此!true的值为 false，更复杂一些的写法是(!true==false) ==true，实际开发中我们应尽量采用比较简洁的布尔表达式，就像用 x 来表示x==true。
	*/
	var aVar = 10
	fmt.Println(aVar == 5)  // false
	fmt.Println(aVar == 10) // true
	fmt.Println(aVar != 5)  // true
	fmt.Println(aVar != 10) // false
	/*
		布尔值可以和 &&（AND）和 ||（OR）操作符结合，并且有短路行为，如果运算符左边的值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值，因此下面的表达式总是安全的
		因为&&的优先级比||高（&& 对应逻辑乘法，|| 对应逻辑加法，乘法比加法优先级要高）
	*/
	s := 'x'
	fmt.Println(s != 'a' && s == 'x') //两个都为true，结果为true
	fmt.Println(s == 'x' || s == 'a') //只要一个为true，结果为true
	/*
		布尔值并不会隐式转换为数字值 0 或 1，反之亦然，必须使用 if 语句显式的进行转换：
	*/
	i := 0
	if s == 'x' {
		i = 1
	}
	fmt.Printf("i = %d\n", i)
	//调用函数
	c := btoi(true)
	d := btoi(false)
	fmt.Println(c)
	fmt.Println(d)
}
