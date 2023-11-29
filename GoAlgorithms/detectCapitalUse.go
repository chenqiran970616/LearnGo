package main

import "fmt"

func main() {
	str := "XYZ"
	arr := []byte(str)
	fmt.Println(arr)
	k := 0
	for i, _ := range arr {
		if arr[i] >= 65 && arr[i] <= 90 {
			k++
		}
	}
	if k == len(arr) || k == 0 || (arr[0] >= 65 && arr[0] <= 90) {
		fmt.Println("true")
	}
}
