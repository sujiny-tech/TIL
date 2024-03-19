# os 
> 파일/디렉토리 관련 다룰 일이 있어서, 사용했던 함수 정리📝  
> go 1.16 이후부터는 io/ioutil 지원 중단 ([go.dev/doc/go1.16 참고](https://go.dev/doc/go1.16#ioutil))


## 파일
### 1. 파일 쓰기
+ `func os.Create(name string) (*os.File, error)`
   + 입력 : 파일명
+ `func CreateTemp(dir, pattern string) (*os.File, error)`
   + 임시파일을 생성할 때 사용   
     > 해당하는 경로에 입력한 접두어를 갖는 임시파일을 생성하며, 프로그램 종료시 임시파일은 자동 삭제됨
   + 입력 : 생성될 디렉토리 경로(기본값:"/tmp"), 접두어(기본값:"tmp")
   

### 2. 파일 읽기
+ `func Open(name string) (*os.File, error)`
   + 입력 : 파일명

### 3. 파일 삭제
+ `func os.Remove(name string) error`
   + 하나의 파일을 삭제할 때 사용
   + 입력 : 파일명
+ `func os.RemoveAll(path string) error`
   + 디렉토리를 포함하여 모두 삭제할 때 사용
   + 입력 : 삭제하려고하는 경로(ex: path+"/..")

## 디렉토리
### 1. 디렉토리 생성
+ `func os.Mkdir(path string, perm fs.FileMode) error`
   + 하나의 디렉토리 생성할 때 사용
   + 입력 : 폴더명(ex: "test"), 퍼미션모드상수(32비트 부호없는 정수형 상수값)
+ `func os.MkdirAll(path string, perm fs.FileMode) error`
   + 디렉토리 계층(하위 디렉토리) 구분하여 생성할 때 사용
   + 입력 : 폴더명(ex: "test/1/2/3/"), 퍼미션모드상수(32비트 부호없는 정수형 상수값)

## 이외
+ **os에 따른 분기 처리**
   + `const GOOS string = goos.GOOS`
      + runtime 패키지에서 정의된 GOOS 상수 이용
      + os에 따라 darwin, freebsd, linux, windows 등을 반환함

## 예제
### [1️⃣예제: PEM 파일 생성](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/os/file_example/1.%20pemfile_example.go)

   


## 참고
+ [golang os package🌟](https://pkg.go.dev/os#Create)
+ [참고한 다른사람의 블로그🌟](https://zerostarting.tistory.com/39)
