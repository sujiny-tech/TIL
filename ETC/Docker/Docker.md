# Docker

## Docker를 활용한 배포 순서
1. Dockfile 작성
2. image 생성
3. image를 container registry에 업로드
4. 서버에서 다운로드 및 실행 

## 실습

### [예제1: 기초 실습](https://github.com/sujiny-tech/TIL/blob/main/ETC/Docker/1.%20test_docker/Dockerfile)
+ 아래 "참고" 中 3번째 링크를 참고하여, 기초 실습을 학습했음📝
+ docker build를 통해 간단하게 "Hello golang..."을 출력하는 gohello 이미지를 빌드하였음. 

### [예제2: 간단한 웹서버 배포]
+ 아래 "참고" 中 4번째 링크를 참고하여, 실습하였음

+ **image 생성 관련 명령어**
   + `docker build --tag <image_name>` : 원하는 이미지 이름을 지정하여 docker 이미지 생성하는 명령어
   + `docker build` : default로 latest로 docker 이미지 생성됨 
   + `docker image ls` : 현재경로에서 docker 이미지 목록
   + `docker image tag <"source"iamge_name> <new tag-image>`: "source" 이미지로, 새로운 tag를 달은 이미지 생성하는 명령어
   + `docker image rm <remove image_name>`: 특정 이미지 삭제하는 명령어

+ **image > container로 실행하는 명령어**
   + `docker run <image_name>` : container 내부에서 특정 이미지 실행하는 명령어
   + `docker run --publish <host_port>:<container_port> <image_name>` : container에 대한 내부 포트를 외부포트로 노출하기 위해서는 --publish 플래그 사용해야함. 
      + 즉, host와 container 포트 연결(포트포워딩)
   + `docker run -d -p <host_port>:<container_port> <image_name>` : container를 백그라운드로 실행시키기 위해서 -d, --detach 플래그 사용해야 함.
   + `docker ps` : 실행되고있는 container 프로세스 목록 확인하기 위함

...ing

### 참고
+ [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
+ [Docker Golang 관련](https://docs.docker.com/language/golang/build-images/)
+ [Docker 기초실습 참고✨👍](https://nayoungs.tistory.com/entry/Docker-Docker%EC%97%90-Go-%EB%B0%B0%ED%8F%AC%ED%95%98%EA%B8%B0)
+ [Docker 웹서버 배포 참고](https://docs.docker.com/language/golang/build-images/)
+ [Docker 초급자 온라인북✨👍](http://www.pyrasis.com/private/2014/11/30/publish-docker-for-the-really-impatient-book)
