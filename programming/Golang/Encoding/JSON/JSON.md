# JSON Enconding/Decoding
> 몇개월 전부터 계속 개발할 때 사용 중인 JSON 정리📝

## JSON function
### json.Encoder/json.Decoder
+ 표준 입출력, 파일 (Reader/Writer) 인터페이스를 사용해서 stream 기반으로 동작시 Encoder/Decoder 사용   

### json.Marshal/json.Unmarshal
+  bytes, slice, string에는 Marshal/Unmarshal 함수 적합   
 
### json.MarshalIndent
+ 들여쓰기를 추가해서 가독성을 높이고 싶을때, json.MarshlIndent 함수를 사용
   + `func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)`   
   + ex) ``` b, _ := json.MarshalIndent(data, "", "  ")  ```


## ✨ 주의할 점 ✨

   + JSON 형식이 한글자라도 다르면, 에러가 나기 때문에 형식에 맞게 잘 구현해야 함.   

   + Python과는 다르게, go에서 []byte 타입의 value에 대해 base64 인코딩이 적용 되어 있었음. 그래서 Python으로 송수신 시 base64 인코딩 된 값을 확인했었음.   
 
      > 따라서, 언어마다 JSON 패키징 시 다른 부분이 있는지 체크 필요 ❗    

      > 아래 이미지 참고 (golang json document)

      > <img src="https://user-images.githubusercontent.com/72974863/141250417-a3151ce5-a169-4a2c-b580-8a7548ea34d5.png">   

## ✨JSON 태그✨
+ golang 구조체의 필드에서 태그(tag)를 이용해서 메타정보를 추가하여 확장이 가능함 ❗❗❗
   + 태그한 필드의 데이터가 존재하지 않는 경우, 해당 필드는 null값으로 출력됨
   + 데이터가 존재하지 않는 경우, 해당 필드를 제외하고 marshal을 수행할수있도록 ",omitempty" 추가❗❗
+ golang 구조체를 여러 형태로 Marshal하는 경우가 존재할때, 멀티 태그(Multi tag)를 달아서 처리 가능함 ❗❗❗
   + 말 그대로, json나 xml형태로 리턴해줘야하는 경우 멀티태그로 둘다 태그를 달아주면 처리 가능함
   + ex) ``` Name string `json:"name" xml:"name"` ```
+ json 태그에 "-" 값을 설정해서, marshal/unmarshal 하지 않도록 설정 가능함❗❗❗   
   + ex) ``` Name string `json:"-"` ```

## 그외
+ 하나를 위한 통신은 문제 없겠지만, 다양한 곳으로부터 요청/응답이 왔다갔다 할 수 있음. 따라서 필수 Json 값 외에 인터페이스를 이용해서 유연하게 처리할 수 있음.   
   + ex) API마다 response데이터가 다를 수 있으므로, 다음과 처리 할 수 있음.   
      ``` 
       type Response struct {   
             Code         int         `json:"code"`   
             Msg          string      `json:"msg"`   
             ResponseData interface{} `json:"responseData"`      
      ```
      
   + 응답메시지 구조가 다를 때 사용할 수 있음.


### [👉간단한 샘플 코드는 여기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/JSON/JSON_example.go)    


## 관련 에러
+ <, >, & 문자 JSON 포함 시 깨짐 현상    


   > 참고 : [stackoverflow : How to stop json.Marshal from escaping < and >?](https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and)
   + Response에 >, < 문자 포함 시 깨져서 오는 현상 발생하였음
   + <, > 는  "\u003c" 및 "\u003e"로 escape되어 일부 브라우저에서 html로 잘못해석하지 않도록 함
   + &도 마찬가지로 "\u0026"으로 escape됨
   + 따라서 **json.marshal에서 escape를 비활성화해야 됨**
      ```
      enc:=json.NewEncoder(<data..>)
      enc.SetEscapeHTML(false)
      ```





# 참고
+ [go package document 💫](https://pkg.go.dev/encoding/json)
+ [JSON 형태 검증/보기좋게 정리해주는 사이트👍](https://jsonlint.com/)
+ [JSON >> string 으로 변환해주는 사이트✨](https://jsontostring.com/)
+ [JSON 가이드](https://www.sohamkamani.com/golang/json/)
