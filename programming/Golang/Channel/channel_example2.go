package main

import (
	"fmt"
	"time"
)

func run1(ch1 chan int) {
	time.Sleep(1 * time.Second) //1초 지연 후, ch1에 데이터 송신
	ch1 <- 1
}

func run2(ch2 chan int) {
	time.Sleep(2 * time.Second) //2초 지연 후, ch2에 데이터 송신
	ch2 <- 2
}

func main() {
	// channel : goroutine간의 통신 채널
	// 고루틴끼리 메시지를 전달할 수 있는 메시지 큐(순서대로 쌓이고, 맨처음온것부터 차례대로 진행)

	ch1 := make(chan int) //size=0인 채널
	ch2 := make(chan int) //size=0인 채널

	go run1(ch1)
	go run2(ch2)

EXIT:
	for {
		select {
		case <-ch1:
			fmt.Println("finish run1") //ch1 데이터 수신

		case <-ch2:
			fmt.Println("finish run2") //ch2 데이터 수신
			break EXIT
		}
	}
}
