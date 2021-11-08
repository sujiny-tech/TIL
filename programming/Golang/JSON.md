# JSON Enconding/Decoding
> 몇개월 전, 데이터 전달 시 사용해봤던 JSON 정리

## JSON function
### json.Encoder/json.Decoder
+ 표준 입출력, 파일 (Reader/Writer) 인터페이스를 사용해서 stream 기반으로 동작시 Encoder/Decoder 사용   

### json.Marshal/json.Unmarshal
+  bytes, slice, string에는 Marshal/Unmarshal 함수 적합   
 

## ✨ 주의할 점 ✨

   + JSON 형식이 한글자라도 다르면, 에러가 나기 때문에 형식에 맞게 잘 구현해야 함.   


   + Python과는 다르게, golang에서는 기본적으로 value에 대해 base64 인코딩이 적용 되어 있었음.   


      > 따라서, 언어마다 추가되는 부분이 있는지 체크 필요 ❗   


### [👉간단한 샘플 코드는 여기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/JSON_example.go)    


# 참고
+ [go package document 💫](https://pkg.go.dev/encoding/json)
