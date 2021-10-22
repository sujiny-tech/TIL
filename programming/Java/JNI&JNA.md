# JNI(Java Native Interface)

+ Virtual Machine에서 실행되는 java 프로그램이 해당 플랫폼의 **native code(C/C++)에 접근할 수 있도록 제공하는 interface**를 뜻함.
+ C/C++로 구현된 **동적 라이브러리와 연동**하기 위해 사용함 

+ JNI를 통해 C/C++ dll 사용하기
   > <U>자바에서 공식적으로 지원하는 인터페이스</U>로, 정해진 형태로 c++ 라이브러리를 제작하거나 **JNI 규격을 따르면서 기존 라이브러리를 호출하는 C++ 래핑 라이브러리를 제작**해야 함   


   + **STEP1** : native code로 java 소스 코드 작성    
   + **STEP2** : 소스 컴파일해서 * .class 파일 생성    
   + **STEP3** : javah를 이용해서 dll prototype이 포함되어 있는  * .h 파일 자동 생성   


      + javah로 헤더 생성하는 과정에서 발생한 에러 
         + `Error: Could not find class file for '패키지명.클래스명'.`   


            > stackoverflow의 글을 보고 해결 완료 : [참고 링크 ✨](https://stackoverflow.com/questions/19137201/javah-tool-error-could-not-find-class-file-for-hellojni)   

- - - - - - - - - - - - - - - - 

# JNA(Java Native Access)
> 첨부되어 있는 jna.jar import →  jna.jar 패키지를 이용해 다양한 dll 로드하여 사용 가능

+ 커뮤니티에서 제작한 라이브러리로, 깃허브에서 다운로드 받고 프로젝트에 포함시켜야 함   

  > JNI와는 달리 래핑 라이브러리를 만들지 않아도 간편하게 기존 라이브러리를 자바에서 사용 가능
  + **STEP1** : 깃허브에서 해당 라이브러리 다운로드 → [github 링크 ✨](https://github.com/java-native-access/jna)   




+ **TODO** : JNI 실행 후, JNA 예제 작성 🖥

   
   
# 참고 ⭐
+ [JNI unsigned char 변환 관련 stackoverflow](https://stackoverflow.com/questions/25259095/jni-android-jbytearray-to-unsigned-char-and-viceversa)
+ [JNI Tutorial ✨](https://sungcheol-kim.gitbook.io/jni-tutorial/chapter13)
