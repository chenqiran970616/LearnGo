package main

import "fmt"
import "sync"

//sync使用
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//每当有一个goroutine运行+1
		go func(x int) {
			defer wg.Done()
			//通过done来-1，在启动的goroutine中
			fmt.Print("", x)
		}(i)
	}
	fmt.Printf("\n%#v\n", wg)
	wg.Wait()
	//等待运行完成，计数器=0
	fmt.Println("----------------")
}
