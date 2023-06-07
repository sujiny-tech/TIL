package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct {
	tag    string
	buffer []int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	pool := sync.Pool{ // Pool 할당
		New: func() interface{} { // Get 함수 사용될 때 호출될 함수 정의
			data := new(Data)             // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 반환
		},
	}

	//-------------------------------------------------------
	// 10개 goroutine으로 *Data 타입 데이터 랜덤 값 저장 및 보관
	for i := 0; i < 10; i++ { // goroutine 10개 생성
		go func() {
			data := pool.Get().(*Data) // Pool에서 *Data 타입으로 데이터를 가져옴
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 슬라이스에 랜덤 값 저장
			}
			fmt.Println(data) // data 내용 출력
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // Pool에 객체 보관
		}()
	}

	//-------------------------------------------------------
	// 10개 goroutine으로 *Data 타입 데이터 짝수 값 저장 및 보관
	for i := 0; i < 10; i++ { // goroutine 10개 생성
		go func() {
			data := pool.Get().(*Data) // Pool에서 *Data 타입으로 데이터를 가져옴
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n //슬라이스에 짝수 저장
				n += 2
			}
			fmt.Println(data) // data 내용 출력
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // Pool에 객체 보관
		}()
	}

	fmt.Scanln()
}
