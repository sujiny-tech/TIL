# c++에서 c 함수 사용 

+ extern "C" Keyword 사용   

+ **간단한 예제 코드** ✨

   + 먼저, c로 구현된 함수 (calculator.c)가 다음과 같다고 할 때, 
      > <img src="https://user-images.githubusercontent.com/72974863/124121467-e15c8100-daaf-11eb-94ba-e3aed3d82faf.png">   
      > TIL/prograaming/extern_c_calculator.c
      
   + 함수에 해당하는 헤더파일에서 아래와 같이 추가하여야 함.   
      > <img src="https://user-images.githubusercontent.com/72974863/124122407-03a2ce80-dab1-11eb-940e-41e94005956b.png">    
      > TIL/prograaming/extern_c_calculator.h    
      

   + c++ 파일에서는 다음과 같이 c 함수를 사용할 수 있음.
   
      > <img src="https://user-images.githubusercontent.com/72974863/124122661-52506880-dab1-11eb-9d66-8523c5dd10eb.png">   
      > TIL/prograaming/extern_c_example.hpp   
      
      > <img src="https://user-images.githubusercontent.com/72974863/124122750-6bf1b000-dab1-11eb-98cf-5cdfb2576f94.png">   
      > TIL/prograaming/extern_c_example.cpp
      
