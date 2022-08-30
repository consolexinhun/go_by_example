package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// 在我们的例子中，我们将从两个通道中选择。
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 各个通道将在若干时间后接收一个值，这个用来模拟例如
	// 并行的 Go 协程中阻塞的 RPC 操作
	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	// 我们使用 `select` 关键字来同时等待这两个值，并打
	// 印各自接收到的值。

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
