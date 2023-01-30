# fmt
> 가장 처음 접하는 패키지로, c언어 표준 입출력 라이브러리 stdio 헤더와 동일한 역할을 수행함📝   


## fmt Index
> 특히 나는 append, print/printf/println/Sprint를 자주 쓰는 편임❗   

> <img src="https://user-images.githubusercontent.com/72974863/214731610-58238b6a-2167-4e08-aee9-73165d5dbe41.png">

### Append
+ Append(b []byte, a ... any) []byte
   + byte slice에 변수 append하는 함수
+ Appendf(b []byte, format string, a ...any) []byte
   + 지정한 형식에 따라, byte slice에 값 append하는 함수
+ Appendln(b []byte, a ...any) []byte
   + Append 함수처럼 변수를 추가하고, 그다음에 개행문자(\n)가 들어감

### 출력
+ Errorf(format string, a ... any) error
   + error 포맷팅하는 함수
+ Print(a ...any) (n int, err error)
   + 출력하는 함수로, 내부적으로는 Fprint을 사용함
+ Printf(format string, a ...any) (n int, err error)
   + 지정한 형식에 따라, 값을 출력하는 함수
+ Println(a ...any) (n int, err error)
   + 마지막 개행문자(\n)이 들어간 출력 함수
+ Fprint(w io.Writer, a  ... any) (n int, err error)
   + F가 붙은 함수들은 파일(File) 출력을 의미함
   + 즉, 값을 그대로 문자열로 만들어서 파일에 저장하는 함수
   + io.Writer 타입을 받아 출력하는 함수로, Print는 표준출력인 os.Stdout을 넘겨 터미널에 기록하는 원리임
+ Fprintf(w io.Writer, format string, a ... any) (n int, err error)
   + Fprint처럼 동일하지만, 지정한 형식에 따라 파일에 저장하는 함수
+ Fprintln(w io.Writer, a ... any) (n int, err error)
   + Fprint처럼 동일하지만, 마지막에 개행문자(\n)를 붙여서 파일에 저장하는 함수


### 입력
+ Scan(a ...any) (n int, err error)
   + 콘솔로 입력 받아서, 특정 변수에 저장하는 함수
+ Scanf(format string, a ...any) (n int, err error)
   + 지정한 형식에 따라, 값을 입력하는 함수
+ Scanln(a ...any) (n int, err error)
   + 마지막에 개행문자(\n)가 들어가는 입력함수
+ Sprint(a ...any) string
   + 값을 그대로 문자열로 만들어주는 함수
+ Sprintf(format string, a ...any) string
   지정한 형식에 따라(형식지정하여), 문자열을 만듦
+ Sprintln(a ...any) string
   + 문자열 마지막에 개행문자(\n)가 붙음.
+ Fscan(r io.Reader, a ...any) (n int, err error)
   + 파일을 읽고, 공백이나 개행문자로 구분된 문자열에서 입력을 받는 함수
+ Fscanf(r io.Reader, format string, a ...any) (n int, err error)
   + 형식을 지정하여 파일로부터 입력을 받는 함수
+ Fscanln(r io.Reader, a ...any) (n int, err error)
   + 파일을 읽고, 공백으로 구분된 문자열에서 입력을 받는 함수
+ Sscan(str string, a ...any) (n int, err error)
   + 공백, 개행문자로 구분된 문자열에서 입력을 받는 함수
+ Sscanf(str string, format string, a ...any) (n int, err error)
   + 문자열에서 형식을 지정하여 입력을 받는 함수
+ Sscanln(str string, a ...any) (n int, err error)
   + 공백으로 구분된 문자열에서 입력을 받는 함수




# 참고
+ [golang fmt package](https://pkg.go.dev/fmt)
