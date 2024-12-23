package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func write(ch chan string) {
	for {
		select {
		//写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	/* select可以同时监听一个或多个channel，直到其中一个channel ready
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)

	// 跑2个协程，写数据
	go test1(output1)
	go test2(output2)

	// select 监控
	select {
	case sl := <-output1:
		fmt.Println("sl=", sl)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
	*/

	/* 如果多个channel同时ready，则随机选择一个执行
	// 创建2个管道
	int_chan := make(chan int, 1)
	strint_chan := make(chan string, 1)
	go func() {
		int_chan <- 1
	}()
	go func() {
		strint_chan <- "hello"
	}()
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-strint_chan:
		fmt.Println("string:", value)
	}

	fmt.Println("main结束")
	*/

	/* 可以用于判断管道是否存满 */
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}
