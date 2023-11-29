package main

import "fmt"

func main() {
	var a []int
	a = append(a, 0)
	a = append(a, 1)
	fmt.Println(a)
	k := 0
	for i := 2; i < 5; i++ {
		k = a[i-1] + a[i-2]
		fmt.Println(k)
	}
}
