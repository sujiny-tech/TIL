# javadoc


## 개념
+ JDK와 함께 패키지로 제공되는 도구로, Java 소스 코드의 문서를 생성하는 데 도움을 주는 도구
+ 컴파일 시 모든 주석은 지워지므로, 프로그램 성능에는 영향이 없다고 함
+ 주석을 공식적인 도구로 작성하기 좋은 방법 중 하나라고 생각함


## 사용법
### 1. javadoc tag를 통해 작성
+ "/** <내용> *"과 같은 틀을 지님
+ IntelliJ의 경우 "/**" 타이핑 후 Enter 키를 누르면, 자동으로 틀이 작성됨
+ 태그
   + `@version` : 구현체(클래스, 메서드, 변수 등)에 대한 버전을 나타내는 태그
   + `@author` : 작성자를 정의하기 위한 태그
   + `@deprecated` : 해당 구현체가 곧 삭제, 업데이트 중단을 의미하는 태그
   + `@param` : 메서드의 매개변수(인자 값)을 설명하기 위한 태그
   + `@throws` : 코드에서 throw (예외상황) 정의하기 위한 태그
   + `@return` : 메서드 반환 값 정의하기 위한 태그

### 2. javadoc html 생성 방법
+ 크게 2가지로 나뉨
   + cmd에서 javadoc command를 통해 생성하는 방법
   + IntelliJ에서 지원하는 Javadoc UI를 통해 생성하는 방법 등으로 나뉨


# 참고
+ [참고한 다른 사람의 블로그](https://velog.io/@ming/JavaDoc-%EC%A3%BC%EC%84%9D-%EC%95%8C%EA%B3%A0%EC%93%B0%EC%9E%90)
+ [참고한 다른사람의 블로그](https://agileryuhaeul.tistory.com/entry/Javadoc%EC%9D%B4%EB%9E%80-Javadoc-%EC%82%AC%EC%9A%A9%EB%B0%A9%EB%B2%95)
+ [참고한 다른사람의 블로그💫](https://parkadd.tistory.com/138)
