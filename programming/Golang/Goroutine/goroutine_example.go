package main

import (
	"fmt"
	"sync"
	"time"
)

// number 출력하는 함수
func PrintNumbers() {
	for i := 1; i <= 20; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Print(i)
	}
}

// char 출력하는 함수
func PrintChar() {
	for i := 63; i <= 83; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	//goroutine 적용 X
	PrintChar()
	fmt.Println("")
	PrintNumbers()
	fmt.Println("")

	//-----------------------------
	//goroutine 적용 O
	//1. waitGroup 사용하지 않은 버전
	go PrintChar()
	go PrintNumbers()
	time.Sleep(2 * time.Second)
	fmt.Printf("\nDone1\n")

	//2. waitGroup 사용한 버전
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		PrintChar()
	}()

	go func() {
		defer waitGroup.Done()
		PrintNumbers()
	}()
	waitGroup.Wait()
	fmt.Printf("\nDone2\n")
}
