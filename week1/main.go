package main

import (
	"context"
	"fmt"
	"time"
)

func test1() {
	arr := [5]string{"I", "am", "stupid”", "and", "weak"}
	for index, _ := range arr {
		switch index {
		case 2:
			arr[2] = "smart"
		case 4:
			arr[4] = "strong"
		}
	}
	fmt.Println(arr)
}

func test2() {
	base := context.Background()
	channel := make(chan int, 10)
	ticker := time.NewTicker(1 * time.Second)
	timeout1, _ := context.WithTimeout(base, 15*time.Second)
	go func(timeout1 context.Context) {
		index := 1
		for _ = range ticker.C {
			select {
			case <-timeout1.Done():
				fmt.Println("sender will be closed")
				time.Sleep(1 * time.Second)
				return
			default:

			}
			channel <- index
			fmt.Println(index)
			index++
		}
	}(timeout1)
	time.Sleep(15 * time.Second) // 15s内，sender的print只能打印到10，说明被阻塞了,不会被打印到15
	timeout2, _ := context.WithTimeout(base, 15*time.Second)
	go func(timeout2 context.Context) {
		for _ = range ticker.C {
			select {
			case <-timeout2.Done():
				fmt.Println("receiver will be closed")
				time.Sleep(1 * time.Second)
				return
			default:

			}
			fmt.Println(<-channel)
		}
	}(timeout2)
	time.Sleep(15 * time.Second) // 15s内，receiver的print只能打印到10，说明被阻塞了，不会被打印到15
	close(channel)
	time.Sleep(5 * time.Second) //休眠5s，让receiver退出
	fmt.Println("test2 exit")
}

func main() {
	test1()
	test2()
}
