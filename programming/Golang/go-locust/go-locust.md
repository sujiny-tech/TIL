# go-locust
> python load test tool인 locust에 대해 golang에서도 사용 가능한 방법❗

## 예제
+ 아래 참고 중 첫번째 github에서처럼, 커스텀해서 직접 golang에서 locust 테스트를 동작시킬 수 있음
+ go-locust github에서 client.go와 client_test.go 파일 두개가 있음
   + client.go는 서버 샘플(/swarm을 통해 실제 locust 동작시킴)
   + client_test.go 해당하는 서버주소/접속자수/1초당증가할수(hatchRate) 등을 조절한 클라이언트 샘플
+ 해당 github에서처럼, golang에서도 locust를 커스텀하여 로드테스트 할 수 있음

# 참고
+ [go-locust github](https://github.com/amila-ku/go-locust)
+ [go-locust-client document](https://pkg.go.dev/github.com/amila-ku/go-locust-client)
