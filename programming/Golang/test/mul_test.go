package calc

import "testing"

func BechmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(-2, 22)
	}
}

//sum_test.go와는 다르게 benckmark를 수행하는 함수는 b.N만큼 for문 반복
//그리고 반복문 안에서 성능 측정할 수 있는 함수 호출함
//test 함수와는 다르게 go test -bench . 를 통해 테스트 실행

//1. test 함수는 앞에 Benchmark 로 시작
//2. Benchmark 다음 함수이름이 오고, 함수이름의 첫글자는 항상 대문자
//3. 항상 *testing.B 타입의 매개변수를 받음
