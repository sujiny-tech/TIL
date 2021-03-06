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

+ Error: expected '=', ',', ';', 'asm' or 'atrribute' before ~~
   + ~~ 부분이 선언이 되어있지 않아서 발생
   + ~~ 부분이 '{' 이라면... ';'이나 괄호 입력 하지않아서 발생한 것   


+ Error: conflicting types for ....
   + gcc 컴파일 시 발생한 에러     
   + 함수 형식이 맞지 않은 경우, 또는 main 함수 뒤에서 함수를 정의한 경우 발생(즉, 함수 사용하기 전에 함수 선언이 되었는지 체크할 것)   
   + 나의 경우엔, 함수 형식 중 매개변수 하나의 타입이 맞지 않은 경우에 발생하였음. 해당 부분 수정하여 처리하였음!
   + 추가적으로, 선언했는데로 에러가 난다면... 선언이 여러군데에 되어있는지 체크할 것(선언이 여러군데 되어있어서 발생하는 경우도 있다고 함)   


+ **32bit --> 64bit 소스호환 관련**

   + [참고하면 좋을 링크](https://wiki.kldp.org/wiki.php/32bitCodeTo64bit)   


...
