# Gradle - Multi Project


## 구성방법
+ IntelliJ > New Project 'JAVA' > Build System 'Gradle' 로 root 프로젝트 생성
+ 좌측 상단 메뉴에서 File 메뉴 > Project Structure 메뉴 클릭

+ build.gradle
  + 프로젝트 plugins, version, dependencies, test 설정
    + plugins을 통해 사용할 언어 설정
    + group, version을 통해 프로젝트 버전, 그룹 설정
    + dependencies을 통해 구현 시 사용하는 외부 프로젝트/모듈 설정
      + 멀티 프로젝트 또는 라이브러리 빌드 설정
      + 테스트 시 사용하는 junit 설정
+ settings.gradle
  + rootProject 이름 설정
  + 프로젝트 구성에 따른, path 설정 / include 할 프로젝트 path 설정
+ gradle.properties
  + 프로젝트 디렉터리, 외부 디렉터리, 변수 설정 가능
    + ex. importDir = C:\\Users
    + ex. ProjectStructure = Library
    + 정의한 변수는 settings.gradle에서 사용할 수 있음
  

## 참고
+ [정리가 잘된 다른사람의 블로그 - 멀티 모듈 프로젝트 구성](https://dkswnkk.tistory.com/691)
