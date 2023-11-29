package main

import (
	"fmt"
)

func main() {
	var queries []string
	queries = append(queries, "cbd")
	var words []string
	words = append(words, "zaaaz")

	var min byte
	sum := 0
	for i, _ := range queries {
		arr := []byte(queries[i])
		for j, _ := range arr {
			if arr[j+1] > arr[j] {
				min = arr[j]
			}
		}
	}
	for i, _ := range queries {
		arr := []byte(queries[i])
		for j, _ := range arr {
			if arr[j] == min {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
