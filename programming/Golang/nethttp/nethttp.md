# net/http 
> 기본적으로 go에서 사용하는 net/http 패키지   


## [TODO] function

+ ...


## Error
+ **net/http 관련 에러** 🔥
   + `http: panic serving 127.0.0.1:xxxxx: runtime error: invalid memory address or nil pointer dereference goroutine ...`
   + 원인 : 함수 수행 시, 중간에 의도하지 않았지만 nil값이 만들어지는 경우가 존재하여 해당 에러가 발생하였음.
   + 해결 : 따라서 nil 값에 대해 예외처리 등으로 해결하였음 ❗


## 연관
+ [fasthttp](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/fasthttp/fasthttp.md)
