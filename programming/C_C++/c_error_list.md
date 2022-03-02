# Error list 정리


+ **스택오버플로우** 에러

   + **C6262 Excessiv Stack Usage*
    
   > <img src="https://user-images.githubusercontent.com/72974863/155907203-cdadd9e4-1372-4b55-9a95-967083400d01.png">    
      
   + 설정된 스택 사이즈를 초과해서 사용하기 때문에, 갑자기 메모리가 리셋되고 다른 변수들의 주소로 사용되는 것을 확인하였음 🔥
      
   + 해결방안1) 코드 내 스택 사이즈를 더 크게 설정해주는 방법
      
   + 해결방안2) 스택에 할당된 변수들을 전역변수 또는 힙으로 할당해주는 방법
   + 나의 경우엔, 여러 함수에서 공통으로 사용되는 변수들을 전역으로 처리해주고, 내부에서 사용하고 없애도되는 변수는 힙으로 할당해주었음 ✨
    
+ warning C4996  
   + 'strlwr' 관련하여 the POSIC name for this item is deprecated. Insteade, use the ISO C and C++ conformant name ...
   + 해결방안1)'#define _CRT_NONSTDC_NO_DEPRECATE' 를 추가하여 해당 경고 표시를 막는 방법   

   + 해결방안2) 프로젝트 속성 페이지에서 C/C++ --> Advanced --> Disable Specific Warnings 에 해당 에러코드 기입하여 막는 방법
   + [참고 - stackoverflow](https://stackoverflow.com/questions/46916437/itoa-the-posix-name-for-this-item-is-deprecated)

...
