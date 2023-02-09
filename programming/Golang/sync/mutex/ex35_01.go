package main

import (
	_ "fmt"
	_ "runtime"
	_ "sync"
	_ "time"
)

/*

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 CPU 사용하도록 설정

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched() //다른 고루틴이 cpu 사용할수있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()
			runtime.Gosched() //다른 고루틴이 cpu 사용할수있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) //2초 지연

	fmt.Println(len(data))
}

*/
