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

	var data int = 0
	var rmMutex = new(sync.RWMutex) //Read/write mutex 생성

	go func() {
		for i := 0; i < 3; i++ {
			rmMutex.Lock()
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond) //10ms 지연
			rmMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rmMutex.RLock()
			fmt.Println("read 1 : ", data)
			time.Sleep(1 * time.Second) //1s 지연
			rmMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rmMutex.RLock()
			fmt.Println("read 2 : ", data)
			time.Sleep(2 * time.Second) //2s 지연
			rmMutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
*/
