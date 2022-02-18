package main

/*
#include <stdlib.h>

typedef struct _PERSON {
	char *name;
	int age;
}PERSON;

PERSON* create(char *name, int age){//해당 구조체를 이용해 메모리 할당
	PERSON *p =(PERSON *)malloc(sizeof(PERSON));

	p->name = name;
	p->age= age;

	return p; //할당한 메모리 반환
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var p *C.PERSON //c로 작성한 구조체 선언

	name := C.CString("Sujiny-tech") //c언어의 CString형으로 변환
	age := C.int(100)                //c언어의 int형으로 변환

	p = C.create(name, age) //해당 구조체에 메모리 할당

	fmt.Println("PERSON name :", C.GoString(p.name)) //리턴받은 포인터는 c언어의 문자열 포인터(char*)이므로 go의 형식에 맞게 바꿔줌!
	fmt.Println("PERSON age  :", p.age)

	C.free(unsafe.Pointer(name)) //c에서는 malloc으로 할당된 메모리 free시킴
	C.free(unsafe.Pointer(p))    //C.CString으로 만든 문자열 포인터, c언어의 malloc으로 할당한 포인터 c언어의 free로 메모리 해제 필요
}
