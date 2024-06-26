# test 패키지
> go test 패키지를 통해 테스트 함수 작성 및 실행 가능, 또한 각 함수의 성능을 측정하는 벤치마킹까지 가능 ✨

## 테스트 작성법
+ **파일 형식**
   + "_ test.go" 형식의 이름으로 파일(테스트 함수) 작성해야 됨
+ **코드 내 함수 형식**
   + 코드 내 테스트 함수는 "Test"으로 시작해야 함
   + 테스트 함수는 테스트하려고 하는 함수 앞에 Test를 꼭 붙여줘야 됨 (ex. TestSum, TestMul ...)    
   + Test 다음에 함수이름이 오는데, 함수이름의 첫 글자는 항상 대문자로 시작    
   + Test 함수의 파라미터에 항상 * testing.T 타입의 매개변수를 적어줌
+ **벤치마킹**
   + benchmark를 통해 원하는 함수의 성능을 측정 가능함
   + 코드 내 함수 형식 : 성능 측정하려는 함수 앞에 Benchmark를 꼭 붙여줘야 함 (ex. BenchmarkSum, BenchmarkMul ...)
      + Benchmark 다음에 함수이름이 오는데, Test 함수와 동일하게 함수이름의 첫 글자는 항상 대문자로 시작
      + Test 함수와 마찬가지로, Benchmark 함수의 파라미터에 항상 * testing.T 타입의 매개변수를 적어줌
   + benchmark를 수행하면, b.N만큼 수행하게 되는데 이는 benchmark한 함수의 성능이 균일하게(안정적으로) 나올때까지 조정되는 것임✨
     > [참고: Go testing package](https://pkg.go.dev/testing#hdr-Benchmarks)


## ✨**테스트 실행 명령어** 
   + `go test -run` : 해당 디렉토리에 존재하는 _ test.go 파일 테스트 실행     
   + `go test -benchmem -run=^$ -bench ^"(BenchmarkSum)"$` : 함수에 관해 할당된 메모리와 loop count, 1회 연산당 실행시간(ns/op)을 확인 가능    
      + loop count는 벤치마킹한 해당 함수를 반복한 전체 횟수를 의미   
      + 할당된 메모리는 Bytes/op, Allocs/op로 1회 함수 연산 당 할당된 바이트 수, 1회 함수 연산 당 할당된 메모리를 의미    
      + ns/op와 loop count를 확인할 수 있으므로, TPS(Transaction Per Seconds)도 도출 가능 ❗
   + [이외의 옵션 존재 📄](https://pkg.go.dev/cmd/go/internal/test)

## **그 외** 
   + 테스트 셋업을 위해 반드시 한번 실행되어야 하는 함수가 존재할 때, 처리 방법 
   + b.Run 함수를 활용하여, 하위 벤치마크 테스트를 작성함(즉, 요 부분에 기존 테스트 진행하려고 하는 부분을 작성)
   + b.Run 함수 외의 부분은 반드시 한번만 수행하게 됨 ❗❗
   + [참고 - stackoverflow : Golang benchmark setup database](https://stackoverflow.com/questions/73782483/golang-benchmark-setup-database)

     
## ✨예제
### 1️⃣간단한 샘플 코드 : [sum_test.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/test/sum_test.go)     
### 2️⃣간단한 샘플 코드 : [mul_test.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/test/mul_test.go)

