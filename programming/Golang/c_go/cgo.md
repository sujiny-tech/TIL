# C언어 연동
> C언어로 만든 소스나 필요한 함수를 go에서 사용하는 법 메모📝   

## 연동을 위한 세팅
+ Ubuntu
   + `sudo apt-get install gcc` 를 통해 gcc 설치

+ Centos
   + `sudo yum install gcc` 를 통해 gcc 설치

+ Windows 32bit/64bit
   + https://www.mingw-w64.org/ 에서 MinGW-W64 설치
   + vscode에서 ctrl+shift+P 눌러서 C/C++ 구성 설정 필요함 ❗❗
   + 만약, mingw-w64에서 다운로드받은 installer에서 'the file has been donwloaded incorrectly!'에러가 발생한다면❓
      + zip파일 통째로 다운로드 받아, 압축 해제해서 C:\Program Files\mingw-w64 경로 안에 파일 넣어주면 해결됨❗❗


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
      
   + 정적 라이브러리 만들기 !!!

      + `gcc -c <src_name>.c` : 해당 소스 컴파일    

      + `ar rc <lib_name>.a <컴파일결과파일>.o` : 컴파일된 파일을 공유 라이브러리로 변환
      + `ranlib <lib_name>.a` : 라이브러리에 색인저장   

   + #cgo CFLAGS로 [gcc 컴파일 옵션](https://github.com/sujiny-tech/TIL/blob/main/Linux_CentOS/gcc.md) 지정하고, #cgo LDFLAGS로 링크 옵션 지정해줘야 함
      + -D : 매크로 설정 옵션
      + -I : 헤더파일 경로 옵션
      + -L : 라이브러리 경로 옵션
      + -l : 라이브러리 링크 옵션
+ c로 만든 대용량 소스는 go에서 사용 가능 !!! 
   > 위의 방법대로 진행 가능 !!!
   
   
## 관련 Error 🔥

+ `func literal () used as value`
   + c로 구현해놓은 함수는 반환값이 없는데, cgo로 작성했을때 반환값을 받으려고 했을때 발생한 에러.
   + c에서 구현한 함수의 형태 그대로 사용해야 함.


+ MSVC 의존성있는 헤더파일 등을 사용했을 때, gcc 컴파일러 시 제대로 컴파일되지 않은 경우 발생하였음.
   + visual studio 2019로 빌드했을 때에는 문제가 없었으나, gcc 컴파일 실행 시 발생
   + 최대한 MSVC에 의존성 없도록 소스를 대체해주거나, 의존성없도록 코드를 구현하는 것이 좋음!!!! 
     > 의존성있는 해당 소스를 가져와서 할 순 있으나, 대체적으로 의존성있는 코드는 문제가 자주 발생됨!

+ `fatal error: unexpected signal during runtime execution`   
   + `[signal SIGSEGV: segmentation violation ....]`
      + **SIGSEGV**는 포인터가 가리키는 메모리가 사용하면 안되는 곳인 메모리공간을 사용하려고할 때 발생한다고 함.
         > 대표적으로 Null pointer derefernce(null pointer가 가리키는 데이터를 읽거나 쓰려고)할 때임.   
         > 또는 할당받은 메모리 공간 이외를 건드렸을 때/초기화되지 않은 pointer 사용할 때 발생함.
   + `SIGABRT: abort`
      + **SIGABRT**는 비정상적인 종료를 뜻함. 메모리 할당이 안된 메모리에 접근하거나 예외처리를 위해 정상적으로 사용하는 경우 발생함.
         > 주로 malloc(메모리 할당)처리가 되지 않아서 발생함.

   + **사용하는 포인터에 대한 Null 체크**를 하거나, **할당해제된 포인터가 사용되는 구간이 있는지 확인 필요**함.   
   


+ 이와는 별개로 한번에 관련된 패키지 설치하는 명령어 존재함 ❗
   + `go mod tidy` : 해당 소스 내에서 사용하는 패키지를 모두 다운받음 

## 참고
- [다른언어(C,Python,Java)에서 Go 언어 함수 호출 방법](https://okky.kr/articles/405775)
