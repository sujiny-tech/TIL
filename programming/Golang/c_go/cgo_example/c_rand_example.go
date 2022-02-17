package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func main() {
	r := C.rand()
	fmt.Println(r)
}
