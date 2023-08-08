# Fuzzing
> golang version 1.18에서 추가된 기능 중 하나의 기능이며,   
> Fuzzing(Fuzz testing)은 소프트웨어 테스트 기법으로 무작위 테스트 방식


## 개념
+ 하드웨어나 소프트웨어 테스트를 위한 **무작위 테스트 방식**으로, 유효하면서 예상치않은 **무작위 데이터를 입력하는 방법**
  
   + 개발자가 생각치 못한 버그가 있는 **엣지 케이스를 찾는 데에 사용하는 테스트 기법** 
+ 주로 소프트웨어나 컴퓨터 시스템들의 보안 문제를 테스트하기 위해 사용됨

## 사용법
+ **기존 단위테스트와의 구문 차이**
  
  + 기존 go test 단위 테스트
       + 함수명에 "Test"를 붙여줬음
       + 입력 파라미터로 "t *testing.T" 사용
       + "go test -run=함수명"으로 특정 함수에 대해 테스트 진행
         
  + Fuzzing 테스트
       + "Fuzz"를 함수명에 붙여줘야 함
       + 입력 파라미터로 "f *testing.F" 사용
       + "go test -fuzz=Fuzz"로 퍼징 테스트 실행   
         -fuzztime 플래그를 통해 퍼징 테스트에 걸리는 시간을 제한할 수 있음
  
## 예제
+ 예제1) Reverse() 함수
  ```
  // Reverse() : 문자열 역순으로 변환하는 함수   
  func Reverse(s string) (string, error) {
        if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    b := []byte(s)
    for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b), nil
  }
  ```
  
  ```
  // TestReverse() : Reverse 함수에 대한 단위 테스트
  // tc 3가지의 경우에 대해 결과값 체크   
  func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev := Reverse(tc.in)
        if rev != tc.want {
                t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
  }
  ```

  ```
  // FuzzReverse() : Reverse 함수에 대한 퍼징 테스트
  // 랜덤 string값을 무작위 대입   
  func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
  }
  ```
+ 예제1)의 3번째 퍼징 테스트 결과 해석
  + Failing input written to testdata/fuzz/FuzzReverse/af692587... 디렉터리의 파일 확인해보면, 실패의 경우 입력값 기입되어있음
       + string("泃") 에 대해 오류 발생
       + 해당하는 泃 유니코드 문자는 utf-8에서 한글자지만, 2byte이상의 공간을 차지하므로.. []byte 기준으로 앞뒤를 뒤집으면 예상치 못한 값이 되어버림
       + 따라서 특정 엣지 케이스를 찾아내는 데에 도움이 되는 테스트 기법임


## 참고
+ [golang fuzz tutorial](https://go.dev/doc/tutorial/fuzz)
+ [golang fuzz](https://go.dev/security/fuzz/)
