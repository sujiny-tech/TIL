# Docker 명령어 정리
> Docker로 배포된 DIF Universal Resolver를 접하면서 알게된 명령어를 정리한다📝

## Docker 명령어
### 1. image
+ `docker build -f <directory_dockerfile> -t <image_name>` : image_name으로 directory_dockerfile에 있는 dockerfile로 빌드(images 생성)
+ `docker images` : docker images 조회
+ `docker tag <already_created_image_name:tag> <new_image_Name>:tag` : 기존 이미지에 대해 새로운 태그 추가하여 생성
+ `docker image rm <image_name>` : image_name에 해당하는 image 삭제
+ `docker rmi $(docker images -q)` : 모든 docker image 삭제
  
### 2. container
+ `docker container ls -a` : 모든 container 조회
+ `docker container rm <container_name>` : container_name에 해당하는 특정 container 삭제
+ `docker container stop <container_name>` : container_name에 해당하는 특정 container 중지
+ `docker container start <container_name>` : (container_name에 해당하는)중지된 container를 시작
+ `docker container restart <container_name>` : (container_name에 해당하는) container를 재시작
+ `docker ps` : docker 프로세스 출력
+ `docker logs <container_name>` : 특정 container에 대한 로그 출력
+ `docker logs --tail <num> <container_name>` : 특정 container에 대한 로그 중 마지막 n번째 줄 선택하여 출력
+ `docker logs --since <unix_timestampe> <container_name>` : 특정 container에 대한 로그를 유닉스타임으로 특정 시간이후 로그 출력 가능
+ `docker logs -f -t <container_name>` : -f 옵션으로 실시간 로그 확인가능하고, -t는 timestamp와 함께 로그 출력가능하게 함
   + `/var/lib/docker/containers/${CONTAINER_ID}/${CONTAINER_ID}-json.log` : 해당 경로로 json 형태로 docker 내부에 저장됨
+ `docker container prune` : stop된 docker container 삭제
+ `docker stop $(docker ps -a -q) ` : 모든 container stop

### 3. network
+ `docker network ls` : docker network 조회
+ `docker network create <network_name>` : network_name으로 네트워크 생성
   + 특징
      + default 값은 "bridge" 드라이버로 세팅됨
      + 컨테이너가 동일한 네트워크에 있는 경우, 이름으로 서로를 찾을 수 있음
   + 옵션
      + `docker network create --driver <driver> <network_name>` : 특정 드라이버로 설정하여 네트워크 생성
         + bridge : 하나의 호스트 컴퓨터 내에서 여러 컨테이너들이 서로 통신할 수 있도록 해줌.
         + host : 호스트 네트워크 환경을 그대로 사용할 수 있음. 내부 IP 할당필요없이 곧바로 사용. 즉 호스트 컴퓨터와 동일한 네트워크에서 컨테이너를 실행 시키기 위해 사용됨.
         + overlay : 여러 docker 데몬이 서로 연결될 수 있음. 여러 호스트에 분산되어 돌아가는 컨테이너들 간에 네트워킹을 위해 사용됨.
         + none : 아무런 네트워크를 사용하지 않는 것으로, 외부와 통신하지 않을 때 사용함.
+ `docker network insepct <network_name>` : network_name에 해당하는 네트워크의 상세 정보를 출력
+ `docker network disconnect <driver> <container_name>` : 컨테이너를 특정 네트워크로부터 연결 해제
+ `docker network rm <network_name>` : network_name에 해당하는 네트워크 삭제
+ `docker network prune` : 아무 컨테이너와 연결되지 않은 (사용하지 않는) 네트워크를 한번에 모두 삭제

## Docker compose 명령어

+ `docker-compose --version` : 버전 확인(출력)
+ `docker-compose -f docker-compose.yml pull` : 실행전 이미지/저장소 가져오기
+ `docker-compose -f docker-compose.yml up` : docker-compose.yml을 바탕으로 여러개의 container를 생성하여 시작
+ `docker-compose -f docker-compose.yml up -d` : 백그라운드로 실행
+ `docker-compose stop` : 현재 디렉터리에 있는 docker-compose를 통해 구동한 서비스 중지
+ `docker-compose down` : 현재 디렉터리에 있는 docker-compose를 통해 구동한 서비스 중지 & 구동시 사용된 컨테이너 제거(초기화시킴)

## 참고
+ [docker network 관련 참고한 블로그💫](https://imjeongwoo.tistory.com/113)
+ [docker network 관련 참고한 블로그🌟](https://www.daleseo.com/docker-networks/)
