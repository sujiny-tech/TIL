package main

import f "fmt"

func main() {
	/*
	   // #1
	   row := 5
	   col := 9
	   mid := col / 2
	   f.Println("check", mid)
	   for i := 0; i < row; i++ {
	      for j := 0; j < col; j++ {
	         if j >= (mid-i) && j <= (mid+i) {
	            f.Printf("*")
	         } else {
	            f.Printf(" ")
	         }
	         //f.Printf("*")
	         // i=0 -> j= mid
	         // i=1 -> j= mid-1, mid, mid+1
	         // i=2 -> j= mid-2, mid-1, mid, mid+1, mid+2
	         // i=3 -> j= mid-3, mid-2, mid-1, mid, mid+1, mid+2, mid+3
	      }
	      f.Println()
	   }
	*/

	/*
		row := 6

		for i := 0; i <= row; i++ {
			for k := 0; k < row-i; k++ {
				f.Print(" ")
			} //6, 5, 4, ... 틀 맞춰주기

			var n int = 1
			for j := 1; j <= i; j++ {
				n *= j // n!
			}

			for j := 0; j <= i; j++ {
				var r int = 1
				for x := 1; x <= j; x++ {
					r *= x // r!

				}
				var c int = 1
				for y := 1; y <= i-j; y++ {
					c *= y // (n-r)!

				}
				f.Print(n/(r*c), " ")

				//f.Println(j, "! = ", r)
				//f.Print(i-j, "! = ", c)
				//f.Println()
			}
			f.Println()
		}
	*/

	/*
	   // 2 default로 space +
	   // 1row 1 ()
	   for i := 0; i < row; i++ {
	       for j := 1; j < col; j++ {
	           //default space
	           if j == 1 {
	               for k := 0; k < 2; k++ {
	                   f.Print(" ")
	               }
	           } else if j >= (mid-i) && j <= (mid+i) { //j>=2
	               f.Print(i+1, " ")
	           } else {
	               f.Print(" ")
	           }
	       }
	       f.Println()
	   }
	*/
	/*
	   for i := 0; i < row; i++ {
	       //홀
	       if (i+1)%2 == 0 {
	           for j := 0; j < col; j += 2 {
	               //mid-i
	               if j >= (mid-i) && j <= (mid+i) {
	                   f.Print(" ", i+1, " ")
	               } else {
	                   f.Print(" ")
	               }
	           }
	       } else { // 짝
	           for j := 1; j < col; j += 2 {
	               //mid-i
	               if j >= (mid-i) && j <= (mid+i) {
	                   f.Print(" ", i+1, " ")
	               } else {
	                   f.Print(" ")
	               }
	           }
	       }
	       f.Println()
	   }
	*/
	//num + " "
	/*
	   // #3
	   row := 5
	   col := 9
	   mid := col / 2
	   for i := 0; i < row; i++ {
	      for j := 0; j < col; j++ {
	         if j > (mid-(row-i)) && j < (mid+(row-i)) {
	            f.Printf("*")
	         } else {
	            f.Printf(" ")
	         }
	      }
	      f.Println()
	   }
	*/

	// #4
	row := 9
	col := 9
	mid := col / 2
	for i := 0; i < row; i++ {
		// 위
		c := i
		if i <= mid {
			for j := 0; j < col; j++ {
				//위쪽 삼각형
				if j >= (mid-i) && j <= (mid+i) {
					if j < mid {
						f.Print(c + 1)
						c--
					} else if j == mid {
						f.Print(1)
						c = 1
						c++
					} else {
						f.Print(c)
						c++
					}
				} else {
					f.Print(" ")
				}
			}
			// 아래
		} else {
			c := row - i
			//아래쪽 삼각형
			for j := 0; j < col; j++ {
				if j > (mid-(row-i)) && j < (mid+(row-i)) {
					if j < mid {
						f.Print(c)
						c--
					} else if j == mid {
						f.Print(1)
						c = 1
						c++
					} else {
						f.Print(c)
						c++
					}
				} else {
					f.Print(" ")
				}
			}
		}
		f.Println()
	}

}
