package main

func main() {
	/*
		关键字
		关键字即是被Go语言赋予了特殊含义的单词，也可以称为保留字。

		Go语言中的关键字一共有 25 个：

		break	default 	func	interface	select
		case	defer	go	map	struct
		chan	else	goto	package	switch
		const	fallthrough	if	range	type
		continue	for	import	return	var


		标识符
		标识符是指Go语言对各种变量、方法、函数等命名时使用的字符序列，标识符由若干个字母、下划线_、和数字组成，且第一个字符必须是字母。通俗的讲就是凡可以自己定义的名称都可以叫做标识符。

		下划线_是一个特殊的标识符，称为空白标识符，它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用_作为变量对其它变量进行赋值或运算。

		在使用标识符之前必须进行声明，声明一个标识符就是将这个标识符与常量、类型、变量、函数或者代码包绑定在一起。在同一个代码块内标识符的名称不能重复。

		标识符的命名需要遵守以下规则：
		由 26 个英文字母、0~9、_组成；
		不能以数字开头，例如 var 1num int 是错误的；
		Go语言中严格区分大小写；
		标识符不能包含空格；
		不能以系统保留关键字作为标识符，比如 break，if 等等。

		命名标识符时还需要注意以下几点：
		标识符的命名要尽量采取简短且有意义；
		不能和标准库中的包名重复；
		为变量、函数、常量命名时采用驼峰命名法，例如 stuName、getVal；

		当然Go语言中的变量、函数、常量名称的首字母也可以大写，如果首字母大写，则表示它可以被其它的包访问（类似于 Java 中的 public）；如果首字母小写，则表示它只能在本包中使用 (类似于 Java 中 private）。

		在Go语言中还存在着一些特殊的标识符，叫做预定义标识符，如下表所示：

		append	bool	byte	cap	close	complex	complex64	complex128	uint16
		copy	false	float32	float64	imag	int	int8	int16	uint32
		int32	int64	iota	len	make	new	nil	panic	uint64
		print	println	real	recover	string	true	uint	uint8	uintptr
	*/
}
