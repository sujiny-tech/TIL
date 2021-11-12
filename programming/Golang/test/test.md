# test 패키지
> go test 패키지를 통해 테스트 함수 작성 및 실행 가능, 또한 각 함수의 성능을 측정하는 벤치마킹까지 가능 ✨


+  "_ test.go" 형식의 이름으로 테스트 함수 작성해야 됨   

##  👉간단한 샘플 코드 : [sum_test.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/test/sum_test.go) **분석**    

   + 테스트 함수는 테스트하려고 하는 함수 앞에 Test를 꼭 붙여줘야 됨 (ex. TestSum, TestMul ...)    

   + Test 다음에 함수이름이 오는데, 함수이름의 첫 글자는 항상 대문자로 시작    

   + Test 함수의 파라미터에 항상 * testing.T 타입의 매개변수를 적어줌


## 👉간단한 샘플 코드 : [mul_test.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/test/mul_test.go) **분석**

   + benchmark를 통해 원하는 함수의 성능을 측정할 수 있음 ❗❗   

   + 성능 측정하려는 함수 앞에 Benchmark를 꼭 붙여줘야 함 (ex. BenchmarkSum, BenchmarkMul ...)

   + Benchmark 다음에 함수이름이 오는데, Test 함수와 동일하게 함수이름의 첫 글자는 항상 대문자로 시작

   + Test 함수와 마찬가지로, Benchmark 함수의 파라미터에 항상 * testing.T 타입의 매개변수를 적어줌

## **_ test.go 파일 실행 명령어** ✨

   + `go test -run` : 해당 디렉토리에 존재하는 _ test.go 파일 테스트 실행     

   + `go test -benchmem -run=^$ -bench ^"(BenchmarkSum)"$` : 함수에 관해 할당된 메모리와 loop count, 1회 연산당 실행시간(ns/op)을 확인 가능   
   > loop count는 벤치마킹한 해당 함수를 반복한 전체 횟수, 할당된 메모리에서는 Bytes/op, Allocs/op를 확인 가능 ❗

   + [이외의 옵션 존재 📄](https://pkg.go.dev/cmd/go/internal/test)
