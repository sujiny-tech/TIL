# time 패키지
> golang에서 타임스탬프를 여러 형태로 나타내려고 조사한 것 정리📝

## 👉간단한 샘플 코드 : [time_example.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/time/time_example.go)
> <img src="https://user-images.githubusercontent.com/72974863/199625925-7b5c9144-d056-4ae6-81ca-285541a501e7.png" width="70%" height="70%">
+ line10 : 현재 시간
+ line11 : 현재 시간을 초로 변환
+ line12 : 현재 시간을 나노초로 변환
+ line18 : 현재 시간을 밀리초로 변환
+ line23 : UTC 시간(초까지)으로 출력
+ line24 : UTC 시간(나노초까지)으로 출력
+ line30 : line10~line30까지의 실행시간

## 👉형식 변환(Format)
+ go에서는 패턴 기반의 레이아웃을 통해 포맷팅/파싱을 지원함

+ RFC 3339 형식으로 현재 시간
   + `time.Now().Format("2006-01-02 15:04:05")`
   + `time.Parse(time.RFRC3339, <변환하려는 string>)`

+ ISO 8601 형식으로 변환
   + `time.Now().Format("2006-01-02T15:04:05.000Z")`


### 참고
+ [go time package✨](https://pkg.go.dev/time)
