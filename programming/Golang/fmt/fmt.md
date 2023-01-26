# fmt
> 가장 처음 접하는 패키지로, c언어 표준 입출력 라이브러리 stdio 헤더와 동일한 역할을 수행함📝   


## fmt Index
> 특히 나는 append, print/printf/println/Sprint를 자주 쓰는 편임❗   

> <img src="https://user-images.githubusercontent.com/72974863/214731610-58238b6a-2167-4e08-aee9-73165d5dbe41.png">

### Append
+ Append(b []byte, a ... any) []byte
   + byte slice에 변수 append하는 함수
+ Appendf(b []byte, format string, a ...any) []byte
+ Appendln(b []byte, a ...any) []byte

### 출력
+ Errorf(format string, a ... any) error
   + error 포맷팅하는 함수
+ Print(a ...any) (n int, err error)
   + 출력하는 함수로, 내부적으로는 Fprint을 사용함
+ Printf(format string, a ...any) (n int, err error)
+ Println(a ...any) (n int, err error)
+ Fprint(w io.Writer, a  ... any) (n int, err error)
   + F가 붙은 함수들은 파일(File) 출력을 의미함
   + io.Writer 타입을 받아 출력하는 함수로, Print는 표준출력인 os.Stdout을 넘겨 터미널레 기록하는 원리임
+ Fprintf(w io.Writer, format string, a ... any) (n int, err error)
+ Fprintln(w io.Writer, a ... any) (n int, err error)


### 입력
+ Scan(a ...any) (n int, err error)
+ Scanf(format string, a ...any) (n int, err error)
+ Scanln(a ...any) (n int, err error)
+ Sprint(a ...any) string
+ Sprintf(format string, a ...any) string
+ Sprintln(a ...any) string
+ Fscan(r io.Reader, a ...any) (n int, err error)
+ Fscanf(r io.Reader, format string, a ...any) (n int, err error)
+ Fscanln(r io.Reader, a ...any) (n int, err error)
+ Sscan(str string, a ...any) (n int, err error)
+ Sscanf(str string, format string, a ...any) (n int, err error)
+ Sscanln(str string, a ...any) (n int, err error)

**... ing**


# 참고
+ [golang fmt package](https://pkg.go.dev/fmt)
