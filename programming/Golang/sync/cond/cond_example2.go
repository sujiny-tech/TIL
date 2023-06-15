package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 CPU 사용하도록 설정

	var mutex = new(sync.Mutex)    //뮤텍스 객체 생성
	var cond = sync.NewCond(mutex) //뮤텍스에 대한 조건변수 생성

	c := make(chan bool, 3) //size가 3인 bool형 비동기채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { //goroutine 3개 생성
			mutex.Lock()
			c <- true
			fmt.Println("Wait BEGIN : ", n)
			cond.Wait() //Wait() 함수 사용 : 해당하는 goroutine 대기 처리
			fmt.Println("Wait END   : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c //비동기 채널에서 하나하나씩 값 꺼내기
	}

	mutex.Lock()
	fmt.Println("!!! BROADCAST !!!")
	cond.Broadcast() //Broadcast() 함수 사용 : 대기 중인 모든 goroutine 깨우기
	mutex.Unlock()

	fmt.Scanln()
}
