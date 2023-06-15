package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 CPU 사용하도록 설정

	var mutex = new(sync.Mutex)    //뮤텍스 객체 생성
	var cond = sync.NewCond(mutex) //뮤텍스를 이용해서 조건변수 생성

	c := make(chan bool, 3) //size가 3인 bool형 비동기채널 생성

	//-------------------------------
	//1. goroutine 3개 생성, 대기 처리
	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true
			fmt.Println("Wait BEGIN : ", n)
			cond.Wait() // Wait()함수 사용 : 해당 goroutine 대기
			fmt.Println("Wait END   : ", n)
			mutex.Unlock()
		}(i)
	}

	//-------------------------------
	//2. 비동기 채널 값 꺼내기 (goroutine 3개 모두 실행될 때 까지 기다림)
	for i := 0; i < 3; i++ {
		<-c //비동기 채널에서 값 꺼내기
	}

	//-------------------------------
	//3. 대기 중인 goroutine 깨우기
	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("SIGNAL     : ", i)
		cond.Signal() //Singal() 함수 사용 : 대기 중인 goroutine 하나씩 깨우기
		mutex.Unlock()
	}

	fmt.Scanln()
}
