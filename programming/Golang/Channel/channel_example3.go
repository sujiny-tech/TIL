package main

import (
	"fmt"
	"time"
)

func send2Channel(ch chan<- string) {
	ch <- "hello"
}

func receiveFromChannel(ch <-chan string) {
	result := <-ch
	fmt.Println(result)
}

func main() {
	ch := make(chan string)
	go send2Channel(ch)       //channel에 "hello" 메시지 담기
	go receiveFromChannel(ch) // channel에 담긴 메시지 출력
	<-time.After(time.Second) //출력후 종료하기 위해 지연시키는 것임.
}
