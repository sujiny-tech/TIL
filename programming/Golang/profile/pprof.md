# pprof 패키지
> go 표준 라이브러리에서 프로파일러. cpu, mem 관련해서 프로파일 가능함.   
> CPU 사용부분, 메모리 누수 발견하고, 장기간 안정적으로 유지되도록 개발하는 데에 매우 유용함.

## runtime/pprof
+ 먼저 사용하기 전에, graphviz를 설치해야 함
   + 설치가 안되어있으면, 프로파일 시 에러발생함 (graphviz가 설치되어있지 않다, 설치해야된다는 관련 에러메시지가 뜰 것임)


+ cpu 프로파일 설정
> <img src="https://user-images.githubusercontent.com/72974863/216878976-7c60036a-ccd9-418e-9395-4d9d52a64755.png" width=40% height=40%>   



+ memory 프로파일 설정
> <img src="https://user-images.githubusercontent.com/72974863/216879113-747be115-2fcb-4767-917b-de63a4efeba6.png" width=40% height=40%>


# 참고
+ [go package - runtime/pprof](https://pkg.go.dev/runtime/pprof)
+ [go package - net/http/pprof](https://pkg.go.dev/net/http/pprof)
