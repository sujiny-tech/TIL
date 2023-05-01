package main

import (
	"fmt"
	"time"
)

func sendHello(ch chan string) {
	time.Sleep(time.Second)
	ch <- "hello"
}
func sendWorld(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "world"
}

func receive(ch1 chan string, ch2 chan string, done chan string) {
	var msg string
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Msg : ", msg1)
			msg += msg1

		case msg2 := <-ch2:
			fmt.Println("Msg : ", msg2)
			msg += msg2
			done <- msg
			return
		default:
			fmt.Println("Default...") //1초마다 계속 출력됨(ch1,ch2에 데이터가 없다면)
			time.Sleep(time.Second)
		}
	}
}
func main() {
	ch1 := make(chan string)  //string 타입 메시지를 송수신하는 채널 1
	ch2 := make(chan string)  //string 타입 메시지를 송수신하는 채널 2
	ch3 := make(chan string)  //string 타입 메시지를 송수신하는 채널 3
	go sendHello(ch1)         //"hello" 메시지 받기(1초지연) --> goroutine으로 실행
	go sendWorld(ch2)         //"world" 메시지 받기(3초지연) --> goroutine으로 실행
	go receive(ch1, ch2, ch3) //순서대로 데이터 받으면 종료 --> goroutine으로 실행

	result := <-ch3
	fmt.Println("result : ", result)

}
