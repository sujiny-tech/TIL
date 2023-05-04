package main

import "fmt"

func main() {
	//----------------------------------------------
	c1 := make(chan int) //int 형 channel, size=1

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- i //channel에 i값 보내기
		}
		close(c1) //channel 닫기
	}()

	for i := range c1 { //channel에 있는 값 꺼내기
		// 채널에 값이 몇개나 들어올지 모르기때문에 값을 계속 꺼내기 위해 range 사용
		fmt.Println(i) //꺼낸 값 출력
	}

	//----------------------------------------------
	fmt.Println("<<<<<<<<<<IsClosedChannel?<<<<<<<<<<")
	c2 := make(chan int, 1) //int 형 channel, size=1
	go func() {
		c2 <- 1
	}()

	val, ok := <-c2                     //channel 닫혔는지 확인하기 위함
	fmt.Println("IsClose ? :", val, ok) //i, true : channel이 열려있고, 담겨있는 값 출력
	close(c2)                           //channel 닫기

	val, ok = <-c2                      //channel 닫혔는지 확인하기 위함
	fmt.Println("IsClose ? :", val, ok) //0, false : channel이 열려있고, 담겨있는 값 출력
}
