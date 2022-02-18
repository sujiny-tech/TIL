package main

/*
#include <stdio.h>

int sum_fun(int a, int b) {
	return a+b;
}
void print_hello() {
	printf("Hello, world! \n");
}*/
import "C"
import "fmt"

func main() {
	var a, b int = 1, 2
	C.print_hello()
	r := C.sum_fun(C.int(a), C.int(b))
	fmt.Println(r)

}
