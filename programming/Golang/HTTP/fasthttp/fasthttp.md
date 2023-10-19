# fasthttp

+ 패키지 설치 : `go get github.com/valyala/fasthttp`

+ document 잘 참고해서 원하는 함수 잘 활용 할 것.
   + ctx.Response.SetBody랑 ctx.Response WriteTo 둘다 사용해서 body랑 header가 동시에 보내졌었음
   + 함수 동작을 잘 확인하고 사용해야 됨
   + `ctx.Response.SetBody`
   + `ctx.Response.WriteTo`
   + SetBody vs SetBodyString 성능 비교 관련 [issue #722](https://github.com/valyala/fasthttp/issues/722)


## Do function
+ func Do(req * Request, resp * Response) error
   + 입력받은 request를 수행하고, 해당 request에 대한 response를 받음.
   + 입력할 reqeust에는 전체 url이 포함되어야함.
+ func DoDeadline(req * Request, resp * Response, deadline time.Time) error
   + Do 함수의 동작 + 입력받은 deadline까지 response를 기다리는 함수.
+ func DoRedirects(req * Request, resp * Response, maxRedirectsCount int) error
   + Do 함수의 동작 + maxRedirectsCount 리디렉션에 따라 지정된 response를 반환됨.
   + 리디렉션 수가 지정한 maxRedirectsCount를 초과할 시, ErrTooManyRedirects(error) 반환됨.
+ func DoTimeout(req * Request, resp * Response, timeout time.Duration) error
   + Do 함수의 동작 + 입력받은 timeout까지 response를 기다리는 함수.
   + 해당 timeout까지 response가 반환되지 않으면 ErrTimeout(error)이 반환됨.


## fasthttp.RquestCtx
+ ctx.QueryArgs().Peek(<key>)
   + (예) HTTP POST로 `http://localhost:8080/user?name=sujin` 과 같이 쿼리를 날릴 때, key(name)에 해당하는 value를 추출하려면?
   + 이때 사용되는 함수로, 쿼리에 대해 특정 키 값에 대해 입력한 Value 값을 가져오는 함수
+ ctx.UserValue(<key>)
   + (예) HTTP GET으로 `http://localhost:8080/user/{name}`과 같이 라우팅 설정된 서버에 대해
     `http://localhost:8080/user/sujin` 으로 값을 입력할 때, 입력된 값을 추출하려면?
   + 이때 사용되는 함수로, 사용자가 HTTP GET Method를 통해 입력한 특정 키 값에 대해 value를 가져오는 함수

## ErrorList
+ `error when reading request headers cannot find http request method`
  + defer releaseResponse(resp * Response) 사용한 위치 때문에 발생하는 에러로 보임
     + golang fasthttp package에 따르면, releaseResponse 함수 선언하자마자 response에 접근불가한 것으로 적혀있음
     + 즉, releaseResponse 함수 뒤에는 Response 값에 접근 불가(Response 확인 다 완료된 다음에 해당함수 사용할 것)
  + golang ReleaseRequest(req * Request) 도 마찬가지임
  + [참고한 golang fasthttp package Reference 링크](https://pkg.go.dev/github.com/valyala/fasthttp#ReleaseResponse) 
  + [동일한 에러에 대한 다른 의견(stackoverflow 링크)](https://stackoverflow.com/questions/59723515/fasthttp-error-when-reading-request-headers-invalid-header-key-http-1-1-r-nus)
  + [github golang/fasthttp issue #159](https://github.com/valyala/fasthttp/issues/159)



+ **...ing**


# 참고
+ 관련 이슈 발생/필요 시, 참고했던 글✨
   + response body 관련 : [github valyala/fasthttp issue #411](https://github.com/valyala/fasthttp/issues/411)
   + POST method 세팅 관련 : [github valyala/fasthttp issue #202](https://github.com/valyala/fasthttp/issues/202)
   + POST method 송신 관련 : [github valyala/fasthttp issue #558](https://github.com/valyala/fasthttp/issues/558)
   + [GO fasthttp document 👍](https://pkg.go.dev/github.com/valyala/fasthttp#section-readme)
