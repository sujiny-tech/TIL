# Solidity 기본 문법
> 개인적으로 솔리디티 문법 학습 후 기록📝

## 1. Solidity 소스 파일 구성
### pragma 
- 특정 컴파일러 특징을 사용할 수 있도록 하는 키워드
  - **Version Pragma** : 컴파일/실행할때 사용하는 solidity 버전정보 기입하기 위함
  - Experimental Pragma : 아직 활성화되지 않은 컴파일러 또는 언어의 기능을 활성화하는데에 사용하는 키워드
      - ABIEncoderV2 : 임의로 중첩된 배열과 구조체를 인/디코딩 할 수 있는 인코더
        > `pragma experimental ABIEncoderV2;`
      - SMTChecker : Solidity 언어로 작성된 스마트 컨트랙트의 정적 분석을 수행하는 도구
        > `pragma experimental SMTChecker;`
### Import
- 다른 소스 파일 임포트하는 방법
  - Syntax and Semantics
     > ```import “filename”;```  
     > ```import * as symbolName from “filename”;```  
     > ```import {symboll as alias, symbol12} from “filename”;```    
     > ```import “filename” as symbolName;```
  - Path
  - Use in Actual Compilers
### Comments
- 주석
  - 1줄 주석(Single-line comment)
    > `//`
  - 여러 줄 주석(Multi-line comment)
    > `/*...*/`

       
## 2. 컨트랙트 구조
### 특징
- Solidity의 컨트랙트는 객체 지향 언어의 클래스와 유사함
- **상태변수, 함수, 함수 수정자, 이벤트, 구조체 유형 및 열거형 유형**이 있음
- 컨트랙트는 다른 컨트랙트에서 상속될 수도 있음
- 라이브러리 및 인터페이스에 대한 컨트랙트도 있음

### 구조
(1) 상태 변수(State Variables)   
- 컨트랙트 저장소에 영구적으로 저장되는 변수
- 모든 function이 접근 가능하며, function에 의해 변경된 값은 계속해서 저장됨
     
  > ```
  > pragma solidity >=0.4.0 <0.6.0;  
  > contract SimpleStorage {  
  >  uint storedData; //State variable
  >  // ...
  > }  


(2) 함수(Functions)
- 컨트랙트 내에서 실행 가능한 코드(동작 관련 코드)
- 호출은 내부/외부에서 발생할 수 있음

  > ```
  > pragma solidity >=0.4.0 <0.6.0;
  > contract SimpleAuction {
  > function bid()publicpayable {// Function// ...}
  > }

(3) 함수 변경자(Function Modifiers)
- 함수를 수정하는 데에 권한을 가짐
- 선언적 방식(declarative way)으로 함수의미 수정
- 함수처럼 파라미터를 받을 수 있고, 바디에 특정 조건을 만족하는지 확인하는 코드가 들어감
  
  > ```
  > pragma solidity >=0.4.22 <0.6.0;
  > contract Purchase {
  >     addresspublic seller;
  > 
  > modifier onlySeller() {// Modifierrequire(
  >             msg.sender == seller,
  >             "Only seller can call this."
  >         );
  > _;
  >     }
  >
  > function abort()publicview onlySeller {// Modifier usage//
  > ...}
  > }

(4) 이벤트(Events)
- EVM 로깅 기능을 갖춘 편리한 인터페이스
- 컨트랙트가 작업 수행하는 동안 로그를 남길 수 있게 해주는 기능
  
... ing

## 참고
- [솔리디티 파고들기](https://solidity-kr.readthedocs.io/ko/latest/layout-of-source-files.html)
- [잘 정리된 다른 사람의 블로그⭐](https://velog.io/@mae-zung/Solidity-%EA%B8%B0%EB%B3%B8-%EB%AC%B8%EB%B2%95)
