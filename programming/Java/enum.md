# enum
> java 모듈 작업 중, enum 개념에 대해 정리의 필요성을 느낌.. 따라서 기본 개념들 위주로 정리한다✏️

## 정의
+ 정의
  + 서로 관련있는 상수들끼리 모아 상수들을 정의하기 위해 사용됨
  + enum 클래스형을 기반으로 한 클래스형 선언
+ 문법 : ```enum 열거체 이름 {상수1, 상수2, ...}```
+ 코드 : ```enum Fruit {BANANA, APPLE, BERRY, ...}```
+ 사용법 : ```Fruit.APPLE```

## 특징
+ enum 클래스
  + 열거형으로 선언된 순서에 따라 0부터 index 값을 가짐 (순차적 증가)
  + enum 열거형으로 지정된 상수들은 모두 대문자로 선언함
  + 열거형 변수들을 선언한 후, 마지막에 세미콜론(;)을 찍지 않음
  + 상수와 특정 값을 연결시킬 경우, 마지막에 세미콜론(;)을 붙여줘야 함
  + JDK 1.5이상의 버전에서만 사용가능함
+ enum 메서드
    + 주로 사용되는 메서드는 values(), ordinal(), valueOf()
    + **values()**  : 열거된 모든 원소들을 순서대로 enum 타입의 배열로 반환 
      + ENUM$$VALUES의 카피로, 너무 자주 호출하는 것은 좋지 않다고 함
    + **ordinal()** : 해당 원소에 열거된 순서를 정수 값으로 반환
    + **valueOf()** : 입력받은 값과 열거형에서 일치하는 이름을 갖는 원소를 반환
      + 주어진 String과 일치하는 원소가 없는 경우, IllegalArgumentException 예외 발생

    ```
    public enum Fruit {
       BANANA, APPLE, BERRY
       
       public static void printType() {
          for(Fruit type : values()) {
             System.out.println(type);
          }
       }
       
       public static void printOrder() {
          for(Fruit type : values()) {
             System.out.println(type.ordinal());
          }
       }
       
       public static Fruit convertToType(String fruitName) {
          return valueOf(fruitName);
       }
    }
    ```

## 장점
1. 데이터들 간의 연관관계 표현
2. 상태와 행위를 한곳에서 관리
3. 데이터 그룹관리


## 참고
+ [참고하기 좋은 다른사람의 블로그 - enum 예제](https://mine-it-record.tistory.com/204)
+ [우아한 기술 블로그 - Java Enum 활용기](https://techblog.woowahan.com/2527/)
+ [정리잘된 다른사람 블로그](https://bcp0109.tistory.com/334)
+ [정리잘된 다른사람 블로그](https://velog.io/@red-sprout/Java-enum%EC%9D%80-%EC%99%9C-%EC%93%B0%EB%8A%94%EA%B1%B8%EA%B9%8C-feat.-%EC%9A%B0%EC%95%84%ED%95%9C%ED%98%95%EC%A0%9C%EB%93%A4-%EA%B8%B0%EC%88%A0%EB%B8%94%EB%A1%9C%EA%B7%B8)
