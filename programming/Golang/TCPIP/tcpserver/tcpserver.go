package main

import (
	"fmt"
	"net"
)

//net 패키지 - tcp function 제공해줌
func requestHandler(c net.Conn) {
	data := make([]byte, 4096) // 4096사이즈의 바이트 슬라이스 생성

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽고
		if err != nil {        //err 처리
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:n])) //데이터 출력

		_, err = c.Write(data[:n]) //클라이언트로 데이터 보냄
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000") //TCP 프로토콜에 8000 포트로 연결받고
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close() //main 끝나기전에 연결 대기 닫고

	for {
		conn, err := ln.Accept() // 클라이언트 연결되면 tcp 연결 리턴
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close() //main 끝나기전에 tcp 연결 닫고

		go requestHandler(conn) //패킷을 처리할 함수를 고루틴으로 실행
		//고루틴 = 가벼운 스레드, 현재 수행 흐름과 별개의 흐름을 만듦, 비동기적으로 함수루틴 실행
		//여러 코드 동시에 실행하는데 사용
	}
}

//http://pyrasis.com/book/GoForTheReallyImpatient/Unit03
