# net/http 
> 기본적으로 go에서 사용하는 net/http 패키지   


## [TODO] function

+ ...


## Error
+ **net/http 관련 에러** 🔥
   + `http: panic serving 127.0.0.1:xxxxx: runtime error: invalid memory address or nil pointer dereference goroutine ...`
   + 원인 : 함수 수행 시, 중간에 의도하지 않았지만 nil값이 만들어지는 경우가 존재하여 해당 에러가 발생하였음.
   + 해결 : 따라서 nil 값에 대해 예외처리 등으로 해결하였음 ❗

+ 서버 동작 중 **예상치 못한 에러 발생 시 처리 방법** 🔥
   + 서버 내에서 예상치 못한 에러가 발생하여 서버가 죽거나 멈출 때, 처리가 필요함.
      > 에러가 발생했을 때, 서버를 초기 동작모드로 처리(복구)해줘야 할 필요가 있음 ❗

   + `recover` 함수를 이용해 처리 가능   

      > <img src="https://user-images.githubusercontent.com/72974863/150079005-dcc88574-da5c-4dc4-8c6a-087c01ed92e9.png">   

    + [stackoverflow 관련 글 참고](https://stackoverflow.com/questions/28745648/global-recover-handler-for-golang-http-panic/28746725)

## 연관
+ [fasthttp](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/fasthttp/fasthttp.md)
