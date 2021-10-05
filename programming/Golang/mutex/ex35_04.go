package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print_hello() {
	fmt.Println("Hello !")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 CPU 사용하도록 설정

	once := new(sync.Once) // Once 생성

	for i := 0; i < 3; i++ { //goroutine 3개 생성
		go func(n int) {
			fmt.Println("goroutine : ", n)
			once.Do(print_hello) //but, goroutine 한번만 실행
		}(i)
	}

	fmt.Scanln()
	//한번만 실행되므로, 복잡하게 반복되는 부분에서 초기화를 위해 사용한다고 한다.
}
