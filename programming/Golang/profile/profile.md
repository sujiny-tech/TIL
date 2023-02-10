# Go profile
> go언어로 짠 소스를 **프로파일링**하는 방법 정리📝       
> → "**프로파일링**" : 메모리 사용량, 함수별 CPU 점유시간, tracing 등 어플리케이션을 동적으로 분석하는 것   
> 이를 잘 활용한다면 CPU 사용부분,메모리 누수 발견하고,**장기간 안정적으로 유지되도록 개발**하는 데에 매우 유용함❗   


## 방법
1. go tool pprof 프로그램
2. runtime/pprof 패키지
3. net/http/pprof 패키지
4. test 패키지

## 1. go tool pprof
+ 프로파일을 시각화해주는 툴로써, golang 설치시 같이 설치됨 
+ 먼저 사용하기 전에, graphviz를 설치해야 함
   + 설치가 안되어있으면, 프로파일 시 에러발생함 (graphviz가 설치되어있지 않다, 설치해야된다는 관련 에러메시지가 뜰 것임)
+ 2~4를 통해 얻은 프로파일 파일들을 시각화해주는 툴로 생각하면 됨
+ 명령어
   + `go tool pprof -http :<port> <profile File>`
      + "localhost:<port>"에 접속해서 <profile File>에 대한 시각화된 프로파일을 확인할 수 있음.   
   + `go tool pprof -http :<port> <profile HTTP Endpoint>`
      + "localhost:<port>"에 접속해서 <profile HTTP Endpoint>에 대한 시각화된 프로파일을 확인할 수 있음.    
+ 시각화된 프로파일을 이해하기위해 알아야될 개념
   + **Flat** : 함수가 직접적으로 수행하는 부분에 대한 부하
   + **Cum** : 함수가 실행되기 위해 수행되는 모든 부분에 대한 부하
      ```
      func FuncA() { // FuncA()에 대한 Cum은 FuncA() 전체에 대한 부하를 뜻하고
         FuncB()      
         FuncC()
         a:= "sujiny-tech" // FuncA()에 대한 Flat은 직접적으로 변수를 선언하는 것처럼, 
         b:= 2023          // 직접적으로 수행하는 부분에 대한 부하를 뜻함!
         ...
      }
      ```   
   
## 2. runtime/pprof
+ runtime/pprof 패키지는 CLI와 같이 한번 실행되고, 종료되는 App 프로파일을 위해 사용되는 패키지
+ 해당하는 패키지에서는 CPU, Memory Heap에 대한 프로파일만 얻을 수 있음
   
+ cpu 프로파일 설정
> <img src="https://user-images.githubusercontent.com/72974863/216878976-7c60036a-ccd9-418e-9395-4d9d52a64755.png" width=40% height=40%>   

+ memory 프로파일 설정
> <img src="https://user-images.githubusercontent.com/72974863/216879113-747be115-2fcb-4767-917b-de63a4efeba6.png" width=40% height=40%>

+ `go tool pprof -http :8080 cpu.out` : CPU 사용량 등 프로파일링(cpu.out에 결과값 바이너리 파일로 생성되면서, 해당 포트에 웹사이트를 띄움)
+ `go tool pprof -hhtp :9090 mem.out` : Memory 사용량 등 프로파일링(mem.out에 결과값 바이너리 파일로 생성되면서, 해당 포트에 웹사이트를 띄움)


## 3. net/http/pprof 
+ net/http/pprof 패키지는 서버와 같이 계속 동작중인 App의 프로파일을 위해 사용되는 패키지
+ http Handler 하위에 다양한 프로파일을 얻을 수 있는 Endpoint를 등록하면 됨
   + `/debug/pprof/profile` : cpu 프로파일
   + `/debug/pprof/heap` : 메모리 힙 프로파일
   + `/debug/pprof/block` : block 프로파일
   + `/debug/pprof/threadcreate` : thread 프로파일
   + `/debug/pprof/goroutine` : goroutine 프로파일
   + `/debug/pprof/mutex` : mutex 프로파일   
   

### 4. test
+ go test 에서 단위테스트 실행 시, 프로파일링 수행도 가능함
+ `go test ./... -cpuprofile cpu.out -memprofile mem.out -blockprofile block.out -mutexprofile mutex.out`를 통해 프로파일을 얻을 수 있음
   + `go tool pprof http :8080 cpu.out`
  
  

# 참고
+ [github - google/pprof](https://github.com/google/pprof)
+ [go package - runtime/pprof](https://pkg.go.dev/runtime/pprof)
+ [go package - net/http/pprof](https://pkg.go.dev/net/http/pprof)
+ [잘 정리한 다른사람의 블로그🤓👍](https://ssup2.github.io/programming/Golang_Profiling/)
