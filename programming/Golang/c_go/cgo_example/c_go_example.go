package main

/*
#include <stdio.h>

int sum_fun(int a, int b) {
	return a+b;
}
void print_hello() {
	printf("Hello world~~~~ \n");
}*/
import "C"
import "fmt"

func main() {
	C.print_hello()
	r := C.sum_fun(C.int(2), C.int(22))
	fmt.Println(r)

}
