# net/http 
> 기본적으로 go에서 사용하는 net/http 패키지   


## function
+ func HandleFunc(pattern string, handler func(ResponseWriter, * Request))
   + pattern 인자에 해당하는 값의 경로로 요청이 들어올 시, 두번째 인자의 함수를 수행할 수 있도록 매핑시켜줌
   + 즉 url 패털과 연결할 함수를 등록시켜주는 핸들러 함수임.

+ func ListenAndServe(addr string, handler Handler) error 
   + addr에 대한 HTTP 서버 시작

+ func Get(url string) (resp * Response, err error)
   + 입력한 url에 대해 GET 요청(웹페이지 상 리소스 요청할 때 사용)
   + 단점 : 헤더를 세밀하게 컨트롤 불가능. 이 경우 NewRequest 객체를 직접 생성해서 http.Client 객체를 통해 호출하면 됨.

+ func Post(url, contentType string, body io.Reader) (resp * Response, err error)
   + 입력한 url에 대해 POST 요청
   + 단점 : GET의 경우와 같은 단점을 가짐.


##  type : Request
   + func NewRequest(method, url string, body io.Reader) (* Request, error)
      + method : "GET", "POST", "PUT", "DELETE"
      + url : addr
      + body : 전달할 요청데이터   


        ```
        //example-Get Request
        req, err := http.NewRequest("GET", "http://example.com/test/search", nil)
        if err!=nil {
           panic(err)
        }
        defer resp.Body.Close()
        ```


## Case : Multi port
+ `http.NewServeMux()`함수와 `http.ListenAndServe`함수를 활용하면 multi port 설정 가능❗

  ```
  serverMuxA := http.NewServeMux()
  serverMuxA.HandlerFunc("/hello", helloFunc)

  serverMuxB := http.NewServeMux()
  serverMuxB.HandlerFunc("/hi", hiFunc)
  
  go func() {
     http.ListenAndServe(":8081", serverMuxA)
  }()

  http.ListenAndServe(":8082", serverMuxB)
  ```
  

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
+ [stackoverflow : Can I setup multi port from one web app with Go?](https://stackoverflow.com/questions/23693520/go-how-to-combine-two-or-more-http-servemux)
