package main

import f "fmt"

func main() {

	var result []int
	n := 3
	t := 10 //

	for j := 0; j < t; j++ {
		f.Println()

		var arr []int = make([]int, n)

		f.Println("@@@[", j+1, "] Enter a num @@@")

		for i := 0; i < len(arr)-2; i++ {
			f.Scanf("%d\n", &arr[i])
		}
		//버퍼 비우는 함수?
		/*
			for i:=0 i<n; i++{
				f.Scanf("%d", &arr[i])
			}
			bp:=copy(arr)
		*/

		result = append(result, arr...)
	}
	f.Println("finish!!:", result, len(result))

}
