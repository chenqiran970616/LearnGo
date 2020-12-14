package main

import "fmt"

func main() {
	/*
		切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型），
		这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内。

		Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合，如果将数据集合比作切糕的话，切片就是你要的“那一块”，
		切的过程包含从哪里开始（切片的起始位置）及切多大（切片的大小），容量可以理解为装切片的口袋大小。
	*/

	/*
		从数组或切片生成新的切片
		切片默认指向一段连续内存区域，可以是数组，也可以是切片本身。

		从连续内存区域生成切片是常见的操作，格式如下：
		slice [开始位置 : 结束位置]

		语法说明如下：
		slice：表示目标切片对象；
		开始位置：对应目标切片对象的索引；
		结束位置：对应目标切片的结束索引。
	*/

	//从数组生成切片，代码如下：
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	fmt.Println("-----------------------------------------")

	//其中 a 是一个拥有 3 个整型元素的数组，被初始化为数值 1 到 3，使用 a[1:2] 可以生成一个新的切片 [2]
	/*
		从数组或切片生成新的切片拥有如下特性：
		取出的元素数量为：结束位置 - 开始位置；
		取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
		当缺省开始位置时，表示从连续区域开头到结束位置；
		当缺省结束位置时，表示从开始位置到整个连续区域末尾；
		两者同时缺省时，与切片本身等效；
		两者同时为 0 时，等效于空切片，一般用于切片复位。
		根据索引位置取切片 slice 元素值时，取值范围是（0～len(slice)-1），超界会报运行时错误，生成切片时，结束位置可以填写 len(slice) 但不会报错。
	*/

	//1) 从指定范围中生成切片
	//切片和数组密不可分，如果将数组理解为一栋办公楼，那么切片就是把不同的连续楼层出租给使用者，出租的过程需要选择开始楼层和结束楼层，
	//这个过程就会生成切片，示例代码如下：
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])
	fmt.Println("-----------------------------------------")

	//2) 表示原有的切片
	//生成切片的格式中，当开始和结束位置都被忽略时，生成的切片将表示和原切片一致的切片，并且生成的切片与原切片在数据内容上也是一致的，代码如下：
	b := []int{1, 2, 3}
	fmt.Println(b[:])
	//b 是一个拥有 3 个元素的切片，将 a 切片使用 a[:] 进行操作后，得到的切片与 a 切片一致

	//3) 重置切片，清空拥有的元素
	//把切片的开始和结束位置都设为 0 时，生成的切片将变空，代码如下：
	c := []int{1, 2, 3}
	fmt.Println(c[0:0])
	fmt.Println("-----------------------------------------")

	/*
		直接声明新的切片
		除了可以从原有的数组或者切片中生成切片外，也可以声明一个新的切片，每一种类型都可以拥有其切片类型，表示多个相同类型元素的连续集合，因此切片类型也可以被声明，切片类型声明格式如下：
		var name []Type

		其中 name 表示切片的变量名，Type 表示切片对应的元素类型。
	*/
	// 声明字符串切片，切片中拥有多个字符串
	var strList []string
	// 声明整型切片，切片中拥有多个整型数值
	var numList []int
	// 声明一个空切片，将 numListEmpty 声明为一个整型切片，本来会在{}中填充切片的初始化元素，这里没有填充，所以切片是空的，但是此时的 numListEmpty 已经被分配了内存，只是还没有元素。
	var numListEmpty = []int{}
	// 输出3个切片，切片均没有任何元素，3 个切片输出元素内容均为空。
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)      //声明但未使用的切片的默认值是 nil，strList 和 numList 也是 nil，所以和 nil 比较的结果是 true。
	fmt.Println(numList == nil)      //声明但未使用的切片的默认值是 nil，strList 和 numList 也是 nil，所以和 nil 比较的结果是 true。
	fmt.Println(numListEmpty == nil) //numListEmpty 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false。
	//切片是动态结构，只能与 nil 判定相等，不能互相判定相等。声明新的切片后，可以使用 append() 函数向切片中添加元素。
	fmt.Println("-----------------------------------------")

	//使用 make() 函数构造切片
	/*
		如果需要动态地创建一个切片，可以使用 make() 内建函数，格式如下：
		make( []Type, size, cap )

		其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
	*/
	e := make([]int, 2)
	d := make([]int, 2, 10)
	fmt.Println(e, d)
	fmt.Println(len(e), len(d))
	//其中 e 和 d 均是预分配 2 个元素的切片，只是 e 的内部存储空间已经分配了 10 个，但实际使用了 2 个元素。
	//容量不会影响当前的元素个数，因此 e 和 d 取 len 都是 2。
}
