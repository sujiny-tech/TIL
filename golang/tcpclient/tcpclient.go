package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8000") //tcp프로토콜, 127.0.0.1:8000서버에 연결
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() //main 끝나기 직전에 tcp 연결 닫고

	//서버에서 받은 데이터 출력하는 고루틴
	go func(c net.Conn) {
		data := make([]byte, 4096) // 4096 사이즈의 byte slice 생성

		for {
			n, err := c.Read(data) //서버에서 받은 데이터 읽음
			if err != nil {        //err 처리
				fmt.Println(err)
				return
			}
			fmt.Println(string(data[:n])) //data print
			time.Sleep(1 * time.Second)

		}

	}(client)

	//서버에 데이터를 보내는 고루틴
	go func(c net.Conn) {
		i := 0
		for {
			s := "Hello " + strconv.Itoa(i)

			_, err := c.Write([]byte(s)) //서버로 데이터 보냄
			if err != nil {              //err 처리
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Second) // 1초마다 hello 문자열에 숫자 1씩 증가시켜서 보냄
		}
	}(client)
	fmt.Scanln()
}
