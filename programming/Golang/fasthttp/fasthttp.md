# net/http
> 기본적으로 go에서 사용하는 net/http 패키지   


# fasthttp

+ 패키지 설치 : `go get github.com/valyala/fasthttp`

+ document 잘 참고해서 원하는 함수 잘 활용 할 것.
   + ctx.Response.SetBody랑 ctx.Response WriteTo 둘다 사용해서 body랑 header가 동시에 보내졌었음
   + 함수 동작을 잘 확인하고 사용해야 됨
   + `ctx.Response.SetBody`
   + `ctx.Response.WriteTo`
   + SetBody vs SetBodyString 성능 비교 관련 [issue #722](https://github.com/valyala/fasthttp/issues/722)


+ **...ing**


+ 관련 이슈 발생/필요 시, 참고했던 글✨
   + response body 관련 : [github valyala/fasthttp issue #411](https://github.com/valyala/fasthttp/issues/411)
   + POST method 세팅 관련 : [github valyala/fasthttp issue #202](https://github.com/valyala/fasthttp/issues/202)
   + POST method 송신 관련 : [github valyala/fasthttp issue #558](https://github.com/valyala/fasthttp/issues/558)
   + [GO fasthttp document 👍](https://pkg.go.dev/github.com/valyala/fasthttp#section-readme)
