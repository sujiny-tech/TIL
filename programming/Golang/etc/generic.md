# generic
> go version 1.18부터 generic 추가되어 관련해서 기록📝

## Generic?
+ 타입 파라미터를 통해 하나의 함수나 타입이 **여러 타입들에 대해 동작 가능**하도록 **코드 재사용성을 늘리는 기법**

## 장점
+ 타입 파라미터를 통해 하나의 함수/타입이 **여러 타입을 지원해야되는 함수에서 동작**할 수 있도록 해줌
+ **'타입 제한자'를 이용**해 타입 파라미터로 사용되는 **타입 제한 가능**
+ **인터페이스도 타입 제한자로 사용** 가능
+ 하나의 코드로 여러 타입에 대해 재사용 가능

## 사용법
+ 다음 아래에 대해 제네릭을 적용할 수 있음
   + 함수
   + 타입 제한자 (https://pkg.go.dev/golang.org/x/exp/constraints)
   + 제네릭 타입(구조체)
 
+ 자료구조(슬라이스, 맵, 채널 등)과 같이 컨테이너 타입을 사용할 때 유용
+ 다양한 타입에 대해 비슷한 동작을 해야하는 경우 유용

## 예제
+ **예제 1) Add() 함수**
  ```
  // 정수형에 대한 덧셈을 수행하는 함수
  func Add(a, b int) int {
   return a + b
  }
  
  // 실수형에 대한 덧셈을 수행하는 함수
  func Add(a, b float32) float32 {
   return a + b
  }
  ```

  ```
  import "constraints"

  // 제네릭을 사용하여, 정수/실수에 대해 덧셈을 수행하는 함수
  func add[T constraints.Integer | constraints.Float](a, b T) T {
     return a + b
  }
  ```
  > [T constraint] : 타입 파라미터를 나타냄   
  
   
+ **예제 2) Print() 함수**
  ```
  // 모든 타입에 대해 출력하는 함수
  func Print[T1 any, T2 any](a T1, b T2) {
     fmt.Println(a, b)
  }
  ```
  > [T any] : 모든 타입이 가능하다고 표시한 타입 파라미터

   
+ **예제 3) 타입 제한이 있는 Add() 함수**
  ```
  // 타입 제한 선언
  type Integer interface {
     int8 | int16 | int32 | int64 | int
  }
  func add[T Integer](a, b T) T {
     return a + b
  }
  ```
  > add[T any]로 실행 시, 해당 타입이 + 연산자가 지원되는지 알 수 없어 에러가 발생함   
  > 따라서, 따로 타입제한을 정의하여 처리할 수 있음

   
+ **예제 4) 타입 제한이 있는 Add() 함수 - 별칭 타입까지 포함하는 경우**

  ```
  // 타입 제한 선언( ~ : 별칭 타입까지 포함)
  type Integer interface {
     ~int8 | ~int16 | ~int32 | ~int64 | ~int
  }

  type MyIntType int //별칭 타입 정의
  
  func add[T Integer](a, b T) T {
     return a + b
  }
  ```
  > ~를 통해 별칭타입까지 포함하여 타입 제한이 가능
  



## 참고
+ [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
+ [다른사람의 블로그⭐](https://goldenrabbit.co.kr/2022/01/28/%EC%83%9D%EA%B0%81%ED%95%98%EB%8A%94-go-%EC%96%B8%EC%96%B4-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D-go-%EC%A0%9C%EB%84%A4%EB%A6%AD%EC%9D%98-%EC%9D%B4%ED%95%B4/)

