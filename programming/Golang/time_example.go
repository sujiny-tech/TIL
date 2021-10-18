package main

import (
	"fmt"
	"time"
)

func main() {

	startTime := time.Now() //current time #1
	secs := startTime.Unix()
	nanos := startTime.UnixNano()
	fmt.Println("cur         : ", startTime)
	fmt.Println("cur(unix)   : ", startTime.Unix()) //unix time stamp
	fmt.Print("cur(yymmdd) : ")
	fmt.Println(time.Now().Date()) //

	millis := nanos / 1000000
	fmt.Println("cur(sec)    : ", secs)
	fmt.Println("cur(nanos)  : ", nanos)
	fmt.Println("cur(milis)  : ", millis)

	fmt.Println("startTime   : ", time.Unix(secs, 0))
	fmt.Println("startTime   : ", time.Unix(0, nanos))
	//delay (sleep)
	//time.Sleep(1 * time.Microsecond) //10^-6
	//time.Sleep(1 * time.Nanosecond)  //10^-9
	//time.Sleep(1 * time.Second)      //1sec

	elapsedTime := time.Since(startTime)

	fmt.Printf("Total elapsedTime : %s\n", elapsedTime)
}
