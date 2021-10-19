# JNI(Java Native Interface)

+ **STEP1** : native code로 java 소스 코드 작성    
+ **STEP2** : 소스 컴파일해서 * .class 파일 생성   
+ **STEP3** : javah를 이용해서 dll prototype이 포함되어 있는  * .h 파일 자동 생성
   + javah로 헤더 생성하는 과정에서 발생한 에러 `Error: Could not find class file for '패키지명.클래스명'.`   

      + stackoverflow의 글을 보고 해결 완료 : [참고 링크 ✨](https://stackoverflow.com/questions/19137201/javah-tool-error-could-not-find-class-file-for-hellojni)   


+ TODO : 생성된 * .h를 이용해 * .cpp 파일 작성

# JNA(Java Native Access)
> 첨부되어 있는 jna.jar import →  jna.jar 패키지를 이용해 다양한 dll 로드하여 사용 가능

+ JNI와는 달리 만들어져있는 dll을 수정하지않고, 최사