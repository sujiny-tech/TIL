package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

//-------------------------------------------------------------------------
// Pool : 일종의 캐시, 객체를 사용한 후 보관했다가 다시 사용가능하게 하는 기능
// 1. func (p *Pool) Get()interface{} : pool에 보관된 객체 get
// 2. func (p *Pool) Put(x interface{}) : pool에 객체 put
//-------------------------------------------------------------------------

type DATA struct {
	tag string //pool tag
	buf []int  //데이터 저장용 slice
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := sync.Pool{ //pool 할당
		New: func() interface{} { //Get 함수 사용시 호출될 함수 정의
			data := new(DATA) //새 메모리 할당 (tag, buf 설정 및 공간할당)
			data.tag = "NEWDATA"
			data.buf = make([]int, 10)
			return data //할당한 메모리(구조체 객체) 리턴
		},
	}

	for i := 0; i < 10; i++ { //goroutine 10개 생성
		go func() {
			data := pool.Get().(*DATA) //pool에서 *DATA 타입으로 데이터를 가져옴
			for index := range data.buf {
				data.buf[index] = rand.Intn(100) //buf 슬라이스에 랜덤하게 값 저장
			}
			fmt.Println(data)
			data.tag = "USED" //변경 후
			pool.Put(data)    //pool에 객체 보관
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*DATA) //pool에서 *DATA 타입으로 데이터를 가져옴
			n := 0
			for index := range data.buf {
				data.buf[index] = n //buf 슬라이스에 짝수 저장
				n += 2
			}
			fmt.Println(data)
			data.tag = "USED" //사용(변경)후
			pool.Put(data)    //pool에 객체 보관
		}()
	}

	fmt.Scanln()
	//메모리 효율적으로 관리 가능하며, 수명주기가 짧은 객체는 pool에 적합 x
}
