# 서버에 golang 설치

+ golang 설치 파일 다운로드   

  + `cd /usr/local/src` : 해당 폴더로 이동    

  + `sudo curl -O https://dl.google.com/go/go1.15.5.linux-amd64.tar.gz` : go 1.15.5 다운로드
  + [go 다운로드 사이트](https://golang.org/dl/)에서 최신버전과 이전버전 다운로드 가능

+ 다운로드 파일 압축풀기
   + `sudo tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz`   

   + `cd /usr/local/go/bin` 폴더로 이동 후, `./go version` 입력해서 go version go1.15.5 linux/amd64가 출력되면 성공

+ 환경변수 설정
   + `vi ~/.bash_profile` 입력 후, 원하는 path를 지정 (vi 명령어를 이용해서 입력하고 저장)
      > go 실행을 어느곳에서든 자유롭게 하기 위해서 설정한다. 

      + export GOPATH=$HOME/go   

      + export PATH=$PATH:$GOPATH/bin   

      + export PATH=$PATH:/usr/local/go/bin   

   + `source ~/.bash_profile` 입력해줘야 반영됨 ✨


## 설치 시 이슈
+ go build 되지않은 경우
   + `go : go.mod file not found in current directory  on any parent directory; see 'go help modules'`
   + 해결 : `go env -w GO111MODULE=auto` 명령어 입력
+ go 파일 실행을 위해 관련 package 다운로드 시, git 명령어 인식안되는 경우
   + git 설치하면 package get 가능
   + centos7 기준, `yum install git` 명령어를 통해 git 설치
   
