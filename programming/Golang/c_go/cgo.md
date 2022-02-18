# C언어 연동
> C언어로 만든 소스나 필요한 함수를 go에서 사용하는 법 메모📝   


## 연동을 위한 세팅
+ Ubuntu
   + `sudo apt-get install gcc` 를 통해 gcc 설치

+ Centos
   + `sudo yum install gcc` 를 통해 gcc 설치

+ Windows
   + https://www.mingw-w64.org/ 에서 MinGW-W64 설치

## GO언어에서 C언어 함수 사용

+ [예제 : C언어의 rand함수 go에서 사용하기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/c_go/cgo_example/c_rand_example.go)   
+ [예제 : C언어로 직접 만든 함수 go에서 사용하기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/c_go/cgo_example/c_go_example.go)
   + Go 언어의 자료형을 C언어 함수에 입력값으로 넣으려면, 변환 필요 ❗
   + 해당 예제에서는 Go 언어의 int형을 C.int로 변환하여 입력값으로 넣음 (C.schar, C.uint, C.float, C.short, C.ushort ...)

## C언어의 구조체 및 문자열 포인터 사용

## TODO...
