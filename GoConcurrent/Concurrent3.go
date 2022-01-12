package main

import (
	"fmt"
	"time"
)

//channel通道，默认为nil

//channel写数据
//func main() {
//	c := make(chan int)
//	go writeChan(c, 666)
//	time.Sleep(10*time.Second)
//}

func writeChan(c chan int, x int) {
	fmt.Println("start", x)
	c <- x
	//执行完之后函数就进入阻塞，因为c中数据没有其他goroutine获取。
	close(c)
	fmt.Println("end:", x)
}

//channel接收数据
func main() {
	c := make(chan int)
	go writeChan(c, 666)
	time.Sleep(10 * time.Second)
	fmt.Println("read:", <-c)
	if _, ok := <-c; ok {
		fmt.Println("open")
	} else {
		fmt.Println("close")
	}
}

//以channel 作为函数参数 双向通道 c chan int  单向通道 c chan <- int
func single(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
}
