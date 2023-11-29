package main

import (
	"fmt"
)

//s1 = "waterbottle", s2 = "erbottlewat"
//请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）
func main() {
	s1 := "abcd"
	s2 := "acdb"
	a := []byte(s1)
	b := []byte(s2)

	for i, _ := range a {
		fmt.Println(a[i])
	}
	//for i, _ := range b {
	//	fmt.Println(a[i])
	//}
	l1 := len(s1)
	l2 := len(s2)
	fmt.Printf("len1:%v", l1)
	fmt.Printf("len2:%v", l1)
	var correct int
	correct = 0
	if l1 == l2 {
		for i, _ := range a {
			for j, _ := range b {
				if b[j] == a[i] {
					correct++
					break
				}
				fmt.Println(b[j])
				fmt.Println(i)
			}
		}
		fmt.Printf("correct:%v", correct)
		if correct == l1 {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else {
		fmt.Println("false")
	}
}
