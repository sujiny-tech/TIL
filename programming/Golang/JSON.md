# JSON Enconding/Decoding
> 몇개월 전부터 계속 개발할 때 사용 중인 JSON 정리📝

## JSON function
### json.Encoder/json.Decoder
+ 표준 입출력, 파일 (Reader/Writer) 인터페이스를 사용해서 stream 기반으로 동작시 Encoder/Decoder 사용   

### json.Marshal/json.Unmarshal
+  bytes, slice, string에는 Marshal/Unmarshal 함수 적합   
 

## ✨ 주의할 점 ✨

   + JSON 형식이 한글자라도 다르면, 에러가 나기 때문에 형식에 맞게 잘 구현해야 함.   

   + Python과는 다르게, go에서 []byte 타입의 value에 대해 base64 인코딩이 적용 되어 있었음. 그래서 Python으로 송수신 시 base64 인코딩 된 값을 확인했었음.   
  

      > 따라서, 언어마다 JSON 패키징 시 다른 부분이 있는지 체크 필요 ❗    

      > 아래 이미지 참고 (golang json document)

      > <img src="https://user-images.githubusercontent.com/72974863/141250417-a3151ce5-a169-4a2c-b580-8a7548ea34d5.png">   



### [👉간단한 샘플 코드는 여기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/JSON_example.go)    


# 참고
+ [go package document 💫](https://pkg.go.dev/encoding/json)
