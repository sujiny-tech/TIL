# Docker 설치
> centos7에 docker, docker-compose 설치🔥

## CentOS7 Docker 설치
+ Docker 저장소 설치
   + `curl -fsSL https://get.docker.com/ | sh
+ Docker 서비스 시작
   + `sudo systemctl start docker`
+ Docker 서비스 상태 확인
   + `sudo systemctl enable docker`
+ root 계정이 아닌 특정 계정(username)에서 docker 사용하기 위한 설정
   + `sudo usermod -aG docker <username>`
   + 명령어 입력 후(설정 후) 계정 재접속 필요❗❗
+ Docker 버전 확인
   + `docker --version`
+ Docker 동작 체크 : hello-world 컨테이너 실행 확인
   + `docker run hello-world`

## CentOS7 Docker Compose 설치
+ docker도 설치해야 함
+ Docker-Compose 설치
   + `curl -L "https://github.com/docker/compose/releases/download/v2.6.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose`
   + docker-compose 버전 체크
      + [해당 링크 : GitHub docker-compose release 정보](https://github.com/docker/compose/releases)

+ Docker-Compose 실행 권한 부여
   + `sudo chmod -x /usr/local/bin/docker-compose

+ Docker-Compose 버전 확인
   + `docker-compose --version`

