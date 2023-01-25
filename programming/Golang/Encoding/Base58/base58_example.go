package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
)

func main() {

	text := []byte("sujiny-tech")
	fmt.Println("---------------0.RawData--------------")
	fmt.Println("before text(raw data)     :", text)
	fmt.Println("before text(raw data)     :", string(text))
	fmt.Println("---------------1.Encode---------------")
	enc_text := base58.Encode(text)
	fmt.Println("after  text(base58Encode) :", enc_text)

	before_text := base58.Decode(enc_text)
	fmt.Println("---------------2.Decode---------------")
	fmt.Println("before text(base58Decode) :", before_text)
	fmt.Println("before text(base58Decode) :", string(before_text))

}
