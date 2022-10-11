# Docker 명령어 정리
> Docker로 배포된 DIF Universal Resolver를 접하면서 알게된 명령어를 정리한다📝

## Docker 명령어
+ `docker build -f <directory_dockerfile> -t <image_name>` : image_name으로 directory_dockerfile에 있는 dockerfile로 빌드(images 생성)
+ `docker images` : docker images 조회
+ `docker tag <target_image> <new_tag_image>` : target_image를 new_tag_image로 태그 변경
+ `docker image rm <image_name>` : image_name에 해당하는 image 삭제
+ `docker container ls -a` : 모든 container 조회
+ `docker container rm <container_name>` : 특정 container 삭제
+ `docker ps` : docker 프로세스 출력
+ `docker logs <container_name>` : 특정 container에 대한 로그 출력
+ `docker logs --tail <num> <container_name>` : 특정 container에 대한 로그 중 마지막 n번째 줄 선택하여 출력
+ `docker logs --since <unix_timestampe> <container_name>` : 특정 container에 대한 로그를 유닉스타임으로 특정 시간이후 로그 출력 가능
+ `docker logs -f -t <container_name>` : -f 옵션으로 실시간 로그 확인가능하고, -t는 timestamp와 함께 로그 출력가능하게 함
   + `/var/lib/docker/containers/${CONTAINER_ID}/${CONTAINER_ID}-json.log` : 해당 경로로 json 형태로 docker 내부에 저장됨
+ `docker container prune` : stop된 docker container 삭제
+ `docker stop $(docker ps -a -q) ` : 모든 container stop
+ `docker rmi $(docker images -q)` : 모든 docker image 삭제


## Docker compose 명령어

+ `docker-compose --version` : 버전 확인(출력)
+ `docker-compose -f docker-compose.yml pull` : 실행전 이미지/저장소 가져오기
+ `docker-compose -f docker-compose.yml up` : docker-compose.yml을 바탕으로 여러개의 container를 생성하여 시작
+ `docker-compose -f docker-compose.yml up -d` : 백그라운드로 실행

