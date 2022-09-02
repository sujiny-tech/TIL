# Docker Compose cmd 정리
> Docker로 배포된 DIF Universal Resolver를 접하면서 알게된 명령어를 정리한다📝

+ Docker Compose : 여러 개의 Container로 구성된 app들을 관리하기 위한 간단한 
+ docker-compose.yml : 간단하게 말하면 도커 실행 옵션을 미리 적어둔 문서로, 프로젝트 루트에 파일을 만들고 실행설정을 적어둔 문서임. 터미널에서 도커 명령어를 통해 컨테이너를 실행하거나 죽일 수 있음


## cmd
+ `docker-compose --version` : 버전 확인(출력)
+ `docker-compose -f docker-compose.yml pull` : 실행전 이미지/저장소 가져오기
+ `docker-compose -f docker-compose.yml up` : docker-compose.yml을 바탕으로 여러개의 container를 생성하여 시작
+ `docker-compose -f docker-compose.yml up -d` : 백그라운드로 실행
