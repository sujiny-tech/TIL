# Locust 부하 테스트 Tool 🛠
> API load test tool로 쉽게 테스트 해볼 수 있는 유용한 툴에 대해 정리 ❗


## 필요한 세팅들
+ python 설치 필요
+ locust 설치 필요 : ```pip install locust```

## 가능한 것들
+ 손쉽게 작성해서 원하는 User로 설정해서 원하는 시나리오대로 테스트 가능하다.
   + 원하는 User 수, 1초당 접속할 User 수(spawn rate) 지정 가능   
   + 스크립트를 통해 원하는 테스트 시나리오대로 작성할 수 있는데, 간단한 테스트 예제는 널려있으므로 패스...
   + 실행은 ```locust -f .\해당파일.py``` 명령어를 통해 실행함


+ 테스트 결과를 그래프로 확인할 수 있고, 다운로드 가능하다.
   + http://localhost:8089 접속을 통해 확인 가능   
   + 웹을 통해 Request per second(RPS), Response times 확인 가능


+ 원하는 테스트 루틴대로 작성해서 순차적으로 함수를 call할 수 있다.
   + 현재 locust 버전에 맞게 document를 참고하여 작성해야 함
   + 이전 버전과는 달리 없어진 함수나 이름이 변경된 함수가 있어서 확인 필수 ❗❗❗
   + 원하는 테스트 루틴대로 작성하기 위해 많은 stackoverflow를 참고하였음


### 참고 ✨
+ [Locust Document](https://docs.locust.io/en/stable/)
+ [참고했던 stackoverflow 중 유용했던 글 👍](https://stackoverflow.com/questions/59832109/stop-locust-when-specified-number-of-user-tasks-complete)   
