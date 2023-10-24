# Docker
> Docker 배포 순서, DockerFile 작성 등 전반적인 내용을 기록📝

## Docker를 활용한 배포 순서
1. Dockfile 작성
2. image 생성
3. image를 container registry에 업로드
4. 서버에서 다운로드 및 실행 

## Dockerfile 기본 명령어
> 그 외 명령어는 "참고" 中 1번째 링크를 참고
+ **FROM** : 생성할 이미지의 베이스가 되는 이미지를 지정하는 명령어로, 반드시 지정해야하는 값이며 반드시 한 번이상 입력해야 함
   + ex) ```FROM golang:1.19-alpine```  
+ **MAINTAINER** : 이미지를 생성한 개발자 등을 나타내는 명령어
+ **COPY** : 호스트 컴퓨터에 있는 디렉터리나 파일을 Docker 이미지 내부로 복사하기 위해 사용되는 명령어
   + ex) ```COPY ~/model ./model```     
+ **RUN** : 이미지를 만들기 위해, 컨테이너 내부에서 특정 파일을 실행해야할 때 사용하는 명령어임. Dockfile로부터 도커이미지를 빌드하는 순간 실행됨.
   + ex) ```RUN apt-get update```     
+ **CMD** : 컨테이너가 시작될 때마다 실행되는 명령어로 한번만 사용할 수 있음  
+ **WORKDIR** : shell의 cd 명령문과 같이 컨테이너상에서 작업 디렉토리로 전환을 위해 사용되는 명령어  
+ **EXPOSE** : Dockerfile 빌드로 생성된 이미지에서 노출할 포트를 설정할 때 사용하는 명령어
   + ex) ```EXPORT 8080``` 
+ **ENV** : 환경 변수 할당할 때 사용하는 명령어  


## 실습

### [예제1: 기초 실습](https://github.com/sujiny-tech/TIL/tree/main/Docker/1.%20test_docker/Dockerfile)
+ 아래 "참고" 中 3번째 링크를 참고하여, 기초 실습을 학습했음📝
+ docker build를 통해 간단하게 "Hello golang..."을 출력하는 gohello 이미지를 빌드하였음. 

### 예제2: 간단한 웹서버 배포
+ 아래 "참고" 中 4번째 링크를 참고하여, 실습하였음

+ **image 생성 관련 명령어**
   + `docker build --tag <image_name> .` : 원하는 이미지 이름을 지정하여 docker 이미지 생성하는 명령어
   + `docker build` : default로 latest로 docker 이미지 생성됨 
   + `docker image ls` : 현재경로에서 docker 이미지 목록
   + `docker image tag <"source"iamge_name> <new tag-image>`: "source" 이미지로, 새로운 tag를 달은 이미지 생성하는 명령어
   + `docker image rm <remove image_name>`: 특정 이미지 삭제하는 명령어

+ **image > container로 실행하는 명령어**
   + `docker run <image_name>` : container 내부에서 특정 이미지 실행하는 명령어
   + `docker run --publish <host_port>:<container_port> <image_name>` : container에 대한 내부 포트를 외부포트로 노출하기 위해서는 --publish 플래그 사용해야함. 
      + 즉, host와 container 포트 연결(포트포워딩)
   + `docker run -d -p <host_port>:<container_port> <image_name>` : container를 백그라운드로 실행시키기 위해서 -d, --detach 플래그 사용해야 함.
   + `docker run -d -p <host_port>:<container_port> --name <new_name> <image_name>` : new_name으로 container 이름 지정하는 명령어



+ **프로세스 확인/중지하는 명령어**
   + `docker ps` : 실행되고있는 container 프로세스 목록 확인하기 위함, CONTAINER ID/IMAGE/COMMAND/CREATED/STATUS/PORTS/NAMES가 순차적으로 출력됨.
   + `docker stop <NAMES>`  : 입력한 NAMES에 해당하는 Container 중지하는 명령어 (container id를 입력해도 됨)
   + `docker ps -all` : 실행중이거나 중지된 모든 container를 출력하는 명령어, 또는 -a를 입력해도 됨


# Docker-compose

+ Docker Compose : 여러 개의 Container로 구성된 서비스를 구축하고 실행 관리할 수 있는 기능임
   + docker-compose.yml : 간단하게 말하면 도커 실행 옵션을 미리 적어둔 문서로, 프로젝트 루트에 파일을 만들고 실행설정을 적어둔 문서임. 터미널에서 도커 명령어를 통해 컨테이너를 실행하거나 죽일 수 있음




## Docker Compose를 활용한 배포 순서
1. 각각의 container의 Dockerfile 작성(기존에 공개된 이미지를 사용하는 경우는 해당없음)
2. docker-compose.yml 작성 후, 각각 독립된 container의 실행 정의 실시
3. `docer-compose up` cmd를 실행하여 docker-compose.yml로 정의한 container publish함



### 참고
+ [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
+ [Docker Golang 관련](https://docs.docker.com/language/golang/build-images/)
+ [Docker 기초실습 참고✨👍](https://nayoungs.tistory.com/entry/Docker-Docker%EC%97%90-Go-%EB%B0%B0%ED%8F%AC%ED%95%98%EA%B8%B0)
+ [Docker 웹서버 배포 참고](https://docs.docker.com/language/golang/build-images/)
+ [Docker 초급자 온라인북✨👍](http://www.pyrasis.com/private/2014/11/30/publish-docker-for-the-really-impatient-book)
