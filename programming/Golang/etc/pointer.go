package main

import "fmt"

type Hello struct {
	A string
	B int
}

func Call(a int) *Hello {
	return &Hello{"ggg", a}
}

func Cs(a *Hello) *Hello {
	a = &Hello{"hello in Cs", 5}
	return a
}

func (a *Hello) Add() interface{} {
	a.A = a.A
	a.B = a.B

	return 0.24
}

func main() {

	bh := Hello{"rer", 11}
	ss := bh.Add()
	fmt.Println(bh, ss)

	/*


		var b *int
		i := 10
		b = &i
		fmt.Println(b, *b)

		var t *Hello
		t.A = "aaa"
		t.B = 33
		fmt.Println(t)

		z := Hello{}
		z.A = "hello"
		z.B = 22

		fmt.Println(z)
	*/
	res := Call(3)
	fmt.Println(res) //*res)

	var t *Hello
	//t := &Hello{}
	//t.A = "ee"
	//t.B = 444
	t = &Hello{"hhhh", 55}
	fmt.Println(t)
}
