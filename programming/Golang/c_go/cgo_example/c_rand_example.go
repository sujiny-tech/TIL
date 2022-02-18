package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func main() {
	r := C.rand()
	//c의 rand함수 사용하려면 주석안에 해당 헤더파일 포함 필요
	//위의 주석과 C를 import하는 해당부분이 이렇게 붙어있어야 함!!
	fmt.Println(r)
}
