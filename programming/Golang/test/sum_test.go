package calc

import "testing"

/*
* 간단한 테스트 함수
func TestSum(t *testing.T) {
	r := Sum(1, 2)
	if r != 3 {
		t.Error("3이 아닙니다. r=", r)
	}
}*/

//1. test 함수는 test하려는 함수앞에 Test를 항상 붙여야함
//2. Test 다음에 함수이름이 오고, 함수이름의 첫 글자는 항상 대문자로 만들어야함
//3. 항상 *testing.T 타입의 매개변수를 받음

type TestData struct {
	arg1 int
	arg2 int
	res  int
}

var testdata = []TestData{
	{2, 6, 8},
	{1, 2, 3},
	{-8, 3, -5},
	{22, -15, 7},
	{6, -6, 0},
	{0, 0, 0},
} //간단한 테스트 데이터지만, 함수에서 발생할 수 있는 다양한 예외 상황을 고려해서 기입해야함

func TestSum(t *testing.T) {
	for _, d := range testdata {
		r := Sum(d.arg1, d.arg2)
		if r != d.res {
			t.Errorf("%d+%d의 결과는 %d가 아닙니다. r=%d", d.arg1, d.arg2, d.res, r)
		}
	}
}
