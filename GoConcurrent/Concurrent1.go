package main

import (
	"fmt"
	"time"
)

//启动单个goroutine
//func main()  {
//	go func() {
//		for i := 10; i < 20; i++ {
//			fmt.Print(" ", i)
//		}
//	}()
//	fmt.Println()
//	for i := 0; i < 10; i++ {
//		fmt.Print(" ", i)
//	}
//	time.Sleep(2*time.Second)
//}

//启动20个goroutine
func main() {
	nullTime := "2000-01-01 00:00"
	nullime, _ := time.Parse("2006-01-02 15:04", nullTime)
	fmt.Println(nullime)
	timeString := "2018-08-01 20:08"
	nowTime, _ := time.Parse("2006-01-02 15:04", timeString)
	fmt.Println(nowTime)
}
