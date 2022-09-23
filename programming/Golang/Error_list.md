# go 관련 에러리스트 정리
> go를 사용하면서 겪었던 에러 및 해결방안 기록📝

+ go build 시 아래와 같은 사항이 발생한 경우
  > <img src="https://user-images.githubusercontent.com/72974863/191879872-a06fc911-e2ca-4b9a-b903-59e9395002a6.png" width=70% height=70%>
  + go를 버전업하기 위해 재설치했을 때, 이전 버전의 파일이 제대로 삭제되지 않은 경우 발생하는 것으로 보임
  + 해결방안 : go를 재설치하는 것이 깔끔
     + go env를 통해 GOROOT를 확인하고 해당 폴더 삭제 후, 재설치 진행
     + 보통 Default가 이므로 /usr/local/go 폴더일 것임
     + 따라서, `sudo rm -rf /usr/local/go`를 통해 삭제
     + `sudo wget https://dl.google.com/go/go1.19.1.linuex-amd64.tar.gz` 명령어를 통해 지금 최신버전인 1.19.1를 설치하였음

... ing
