package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용하도록 설정

	var mutex = new(sync.Mutex)    //mutex생성
	var cond = sync.NewCond(mutex) //mutex를 이용해서 조건변수 생성

	//----------------------------------------------------------------------
	// 조건변수 : 대기하고있는 객체 하나만 깨우거나 여러개 동시에 깨울때 사용
	// 1. func (c *Cond) Wait() : 고루틴 실행을 멈추고 대기
	// 2. func (c *Cond) Signal() : 대기중인 고루틴 하나만 깨움
	// 3. func (c *Cond) Broadcast() : 대기중인 모든 고루틴 깨움
	//----------------------------------------------------------------------

	c := make(chan bool, 3) //asy channel 생성

	for i := 0; i < 3; i++ { //고루틴 3개 생성
		go func(n int) {
			mutex.Lock() //mutex 잠금, cond.Wait() 보호 시작
			c <- true    // channel c에 true 송신
			fmt.Println("wait begin : ", n)
			cond.Wait() //조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() //mutex 잠금 해제, cond.Wait() 보호 끝
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // channel에서 값을 꺼냄, 고루틴 모두 실행될때까지 기다림
	}

	//3. Broadcast()
	mutex.Lock() //metex 잠금, cond.Broadcast() 보호 시작
	fmt.Println("<< broadcast ... >>")
	cond.Broadcast() //대기중인 모든 고루틴 깨움
	mutex.Unlock()   //mutex 잠금해제, cond.Broadcast() 보호 끝

	/*
		2. Signal()
			for i := 0; i < 3; i++ {
				mutex.Lock() //mutex 잠금, cond.Signal() 보호시작
				fmt.Println("signal : ", i)
				cond.Signal()  //대기중인 고루틴 하나씩 깨움
				mutex.Unlock() //mutex 잠금 해제, cond.Signal() 보호 끝
			}
	*/
	fmt.Scanln()
}
