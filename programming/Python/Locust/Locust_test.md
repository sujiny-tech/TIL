# Locust 부하 테스트 Tool 🛠
> API load test tool로 쉽게 테스트 해볼 수 있는 유용한 툴에 대해 정리 ❗


## 필요한 세팅들
+ python 설치 필요
+ locust 설치 필요 : ```pip install locust```

## 가능한 것들
+ 웹사이트의 로드 테스트/시스템이 처리할 수 있는 동접자 수를 파악 가능함
+ 서버에서 지원하는 초당 요청 건수 (Request per second)를 대략적으로 알아낼 수 있음
+ 손쉽게 작성해서 원하는 시나리오대로 테스트 가능함
   + 원하는 User 수, 1초당 접속할 User 수(spawn rate) 지정 가능 
     > locust는 한국어로 '메뚜기'라는 의미를 지니는데.. 테스트 중 메뚜기 떼가 웹사이트를 공격한다는 것으로부터 착안된 이름이라고 한다.   
      
   + 스크립트를 통해 원하는 테스트 시나리오대로 작성할 수 있는데, 간단한 테스트 예제는 널려있으므로 패스...
   + 실행은 ```locust -f .\해당파일.py``` 명령어를 통해 실행함

+ 테스트를 진행하는 도중에 테스트환경을 변경하여, 또다른 세팅의 테스트를 이어서 할 수 있음
   + 중간에 User 수, 1초당 접속할 User 수를 변경할 수 있음 ❗❗
      > <img src="https://user-images.githubusercontent.com/72974863/172284547-9a6fffb1-2033-465c-997e-3dcb6b6bc28a.png" width="40%" hegiht="40%">   

+ 테스트 결과를 그래프로 확인할 수 있고, 다운로드 가능함
   + http://localhost:8089 접속을 통해 확인 가능   
   + 웹을 통해 Request per second(RPS), Response times 확인 가능


+ 원하는 테스트 루틴대로 작성해서 순차적으로 함수를 call할 수 있음
   + 현재 locust 버전에 맞게 document를 참고하여 작성해야 함
   + 이전 버전과는 달리 없어진 함수나 이름이 변경된 함수가 있어서 확인 필수 ❗❗❗
   + 원하는 테스트 루틴대로 작성하기 위해 많은 stackoverflow를 참고하였음


## 발생했던 에러들
+ locust response의 status_code값이 0이 발생되는 경우   
  
   + 과도한 요청을 서버에서 받쳐주지 못해 생기는 에러, 시간초과나 연결오류가 원인
   + 따라서, 요청 수를 줄이면 해당 에러가 발생하지 않음
   + [참고한 stackoverflow✨](https://stackoverflow.com/questions/17317162/locust-got-0-response-status-code-and-none-content)

   + **추가적으로** IDC 서버내에서 로드 테스트를 진행했을 때, 중간중간 해당 에러가 발생했었음. **그러나** 내부 PC를 통해 진행했을 때, 해당 에러가 발생하지 않았음 ❗❗❗❗


### 참고 ✨
+ [Locust Document](https://docs.locust.io/en/stable/)
+ [참고했던 stackoverflow 중 유용했던 글 👍](https://stackoverflow.com/questions/59832109/stop-locust-when-specified-number-of-user-tasks-complete)   
