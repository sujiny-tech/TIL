# Go profile
> go언어로 짠 소스를 **프로파일링**하는 방법 정리📝       
> → "**프로파일링**" : 메모리 사용량, 함수별 CPU 점유시간, tracing 등 어플리케이션을 동적으로 분석하는 것   
> 이를 잘 활용한다면 CPU 사용부분,메모리 누수 발견하고,**장기간 안정적으로 유지되도록 개발**하는 데에 매우 유용함❗   


## 방법
1. go tool pprof 프로그램
2. runtime/pprof 패키지
3. net/http/pprof 패키지

## 1. go tool pprof

**...ing**  

## 2. runtime/pprof
+ pprof 데이터를 분석하고 가공해서 웹으로 볼 수 있음
+ 먼저 사용하기 전에, graphviz를 설치해야 함
   + 설치가 안되어있으면, 프로파일 시 에러발생함 (graphviz가 설치되어있지 않다, 설치해야된다는 관련 에러메시지가 뜰 것임)

+ cpu 프로파일 설정
> <img src="https://user-images.githubusercontent.com/72974863/216878976-7c60036a-ccd9-418e-9395-4d9d52a64755.png" width=40% height=40%>   

+ memory 프로파일 설정
> <img src="https://user-images.githubusercontent.com/72974863/216879113-747be115-2fcb-4767-917b-de63a4efeba6.png" width=40% height=40%>

+ `go tool pprof -http :8080 cpu.out` : CPU 사용량 등 프로파일링(cpu.out에 결과값 바이너리 파일로 생성되면서, 해당 포트에 웹사이트를 띄움)
+ `go tool pprof -hhtp :9090 mem.out` : Memory 사용량 등 프로파일링(mem.out에 결과값 바이너리 파일로 생성되면서, 해당 포트에 웹사이트를 띄움)


## 3. net/http/pprof 프로그램
**...ing**


# 참고
+ [go package - runtime/pprof](https://pkg.go.dev/runtime/pprof)
+ [go package - net/http/pprof](https://pkg.go.dev/net/http/pprof)
+ [잘 정리한 다른사람의 블로그🤓👍](https://ssup2.github.io/programming/Golang_Profiling/)
