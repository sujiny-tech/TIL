# build tool


### 정의
- 빌드 관리 도구 : 프로젝트에서 필요한 xml, properties, jar 등을 자동 인식하여 빌드해주는 도구를 뜻함
- 소스코드를 컴파일/테스트/정적분석 등을 하여, 실행가능하도록 빌드해줌
- 프로젝트 정보관리, 테스트 빌드, 배포 등의 작업 뿐만 아니라 외부 라이브러리를 참조하여 자동 다운로드 및 업데이트 해줌


## Ant
- 특징
  - XML 기반의 빌드 스크립트
  - 자유로운 빌드 단위 지정
  - 간단하고 사용하기 쉬움
  - 대규모 프로젝트에서 복잡해지는 경향있음
  - LifeCycle 없음

## Maven

- 특징
  - XML 기반의 빌드 스크립트
  - LifeCyle 도입(Ant에는 빌드 기능만 있음)
  - pom.xml을 통해 Dependency 관리
    - Maven은 자동으로 라이브러리 관리해주는 기능이 있음
    - 다운로드받아서 사용하던 라이브러리에 변동사항이 있으면 자동으로 업데이트하여 적용됨

- 간단 사용법
    - pom.xml 파일을 활용해 빌드 및 관리
        - pom.xml 파일 역할
            - 프로젝트 정보 관리
            - 프로젝트에서 사용하는 외부 라이브러리 관리
            - 프로젝트의 빌드 관련 설정
    - 메이븐(Maven) 대표 태그 
      - modelVersion : maven 버전
      - groupId : 프로젝트 그룹 ID (일반적으로 대표하는 사이트 도메인을 역순으로 적어서 사용한다고 함)
      - artifactId : groupId 외, 다른 프로젝트와는 구분될 수 있는 프로젝트의 id를 작성
      - version : 프로젝트 버전
      - name : 프로젝트 이름
      - description : 프로젝트의 간략한 설명
      - properties : pom.xml 파일 내 빈번하게 사용되는 중복 상수를 정의하는 영역
      - dependencies : 프로젝트에서 의존성을 가지고 사용하는 라이브러리를 정의하는 영역
      - build : 프로젝트 빌드와 관련된 정보를 설정하는 영역
    - maven repository : https://mvnrepository.com/


## Gradle

- Groovy 스크립트를 활용한 빌드 관리도구
    - Groovy : JVM 상에서 실행되는 스크립트 언어, Java와 유사한 문법 구조를 가지며, 호환성이 아주 좋다.
    - Maven에 비해 더 빠른 처리속도를 가지며, 더 간결한 구성이 가능함
    - pom.xml을 사용하는 Maven과는 달리, build.gradle을 사용함
- 멀티 프로젝트(Multi-Project)의 빌드에 최적화하여 설계됨
- 안드로이드 프로젝트의 표준 빌드 시스템으로 채택
- 간단 사용법
    - Gradle 대표 용어
        - repositories : 라이브러리가 저장된 위치 등 설정
        - mavenCentral : 기본 Maven Repository
        - dependencies : 라이브러리 사용을 위한 의존성 설정


### 참고
- [youtube - java 빌드 관리 툴](https://www.youtube.com/watch?v=3Jp9kGDb01g)
- [Gradle 8.10.1 공식 홈페이지](https://docs.gradle.org/current/samples/sample_building_java_applications.html)
- [참고하기 좋은 다른 사람의 블로그💫](https://goddaehee.tistory.com/199)
