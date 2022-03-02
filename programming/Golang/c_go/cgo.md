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

### [예제 : C의 rand함수 go에서 사용하기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/c_go/cgo_example/c_rand_example.go)   
### [예제 : C로 만든 함수 go에서 사용하기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/c_go/cgo_example/c_go_example.go)
+ Go 언어의 자료형을 C언어 함수에 입력값으로 넣으려면, 변환 필요 ❗
+ 해당 예제에서는 Go 언어의 int형을 C.int로 변환하여 입력값으로 넣음 (C.schar, C.uint, C.float, C.short, C.ushort ...)

### [예제 : C의 구조체 및 문자열 포인터 사용하기](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/c_go/cgo_example/c_memory_example.go)
+ c 함수에서 리턴받은 c언어 형식의 변수를 바로 go에서 사용할 수 없으므로, 그때 변환 필요
+ C.CString으로 만든 문자열 포인터, c언어의 malloc으로 할당한 포인터는 반드시 C.free로 메모리 해제시켜줘야함 ❗
   + go에서 c에 관한 메모리 해제할 때, unsafe.Pointer로 변환하여 C.free함수에 넣어줌


## 궁금증 🤔
+ c에서 생성한 라이브러리 사용 가능여부 체크
   > 체크 중 [stackoverflow : How to use C library in Golang(v1.3.2)](https://stackoverflow.com/questions/31868482/how-to-use-c-library-in-golangv1-3-2)

   + 동적 라이브러리 만들기
      + `gcc -fPIC -c <src_name>.c` : 해당 소스를 컴파일   

      + `gcc -shared -o <lib_name>.so <컴파일결과파일>.o` : 컴파일된 파일을 공유 라이브러리로 변환
      
   + 정적 라이브러리 만들기

      + `gcc -c <src_name>.c` : 해당 소스 컴파일    

      + `ar rc <lib_name>.a <컴파일결과파일>.o` : 컴파일된 파일을 공유 라이브러리로 변환
      + `ranlib <lib_name>.a` : 라이브러리에 색인저장   

   + #cgo CFLAGS로 gcc 컴파일 옵션 지정하고, #cgo LDFLAGS로 링크 옵션 지정해줘야 함
      + -D : 매크로 설정 옵션
      + -I : 헤더파일 경로 옵션
      + -L : 라이브러리 경로 옵션
      + -l : 라이브러리 링크 옵션
+ c로 만든 대용량 소스는 go에서 사용 가능한지, 가능하다면 어떻게 사용가능한지 체크


