package main

import (
	"fmt"
	"reflect"
)

//在结构体成员嵌入时使用别名

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {
}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main() {
	/*
		Go语言type关键字（类型别名）


		区分类型别名与类型定义
		定义类型别名的写法为：
		type TypeAlias = Type

		类型别名规定：TypeAlias 只是 Type 的别名，本质上 TypeAlias 与 Type 是同一个类型，就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。

	*/
	// 将NewInt定义为int类型
	type NewInt int
	// 将int取一个别名叫IntAlias
	type IntAlias = int
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

	// 声明变量ab为车辆类型
	var ab Vehicle

	// 指定调用FakeBrand的Show
	ab.FakeBrand.Show()
	// 取ab的类型反射对象
	ta := reflect.TypeOf(ab)
	// 遍历ab的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.
			Name())
	}
}
