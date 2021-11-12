package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type INFO struct {
	Name  string
	Age   int64
	Birth string
}

func main() {

	InPeople := &INFO{
		Name:  "Sujiny-tech",
		Age:   0,
		Birth: "YYMMDD",
	}
	var OutPeople INFO

	fmt.Println("------------1.Mashal/Unmarshal-------------") // bytes, slice, string에는 Marshal/Unmarshal 함수 적합

	fmt.Println(InPeople)
	jbytes, _ := json.Marshal(InPeople)
	fmt.Println(jbytes)

	json.Unmarshal(jbytes, &OutPeople)
	fmt.Println(OutPeople)

	fmt.Println("-----------2.Encoder/Decoder---------------")
	json.NewEncoder(os.Stdout).Encode(&InPeople) // 표준 입출력, 파일 (Reader/Writer) 인터페이스를 사용해서 stream 기반으로 동작시 Encoder/Decoder 사용
	json.NewEncoder(os.Stdin).Encode(&OutPeople)

	fmt.Printf("%+v\n", OutPeople)

}
