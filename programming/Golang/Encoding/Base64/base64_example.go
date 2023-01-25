package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	text := []byte("sujiny-tech")
	fmt.Println("---------------0.RawData--------------")
	fmt.Println("before text(raw data)     :", text)
	fmt.Println("before text(raw data)     :", string(text))
	fmt.Println("---------------1.Encode---------------")
	enc_text := base64.RawStdEncoding.EncodeToString(text)
	fmt.Println("after  text(base64Encode) :", enc_text)

	before_text, err := base64.RawStdEncoding.DecodeString(enc_text)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("---------------2.Decode---------------")
	fmt.Println("before text(base64Decode) :", before_text)
	fmt.Println("before text(base64Decode) :", string(before_text))

}
