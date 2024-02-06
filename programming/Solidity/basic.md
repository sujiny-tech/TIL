# Solidity 기본 문법
> 
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
(2) 함수(Functions)
(3) 

... ing

## 참고
- [솔리디티 파고들기](https://solidity-kr.readthedocs.io/ko/latest/layout-of-source-files.html)
