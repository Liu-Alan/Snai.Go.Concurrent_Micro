package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func write() {
	rwlock.Lock() // 写锁
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	rwlock.Unlock() // 解锁
	wg.Done()
}

func read() {
	rwlock.RLock() //读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 解锁
	wg.Done()
}

func main() {
	/* 互斥锁
	   同一时间有且只有一个goroutine获得锁
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
	*/

	/* 读写互斥锁
	   当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
	   当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
	   读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
	*/
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
