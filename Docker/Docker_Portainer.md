# Portainer 설치
> web에서 Docker를 관리할 수 있는 툴 설치 및 관련 내용들 정리📝

## Portainer 설치
+ 설치를 위해서는 Docker 설치가 되어있어야 함!
+ `docker volume create portainer_data`
+ `docker run -d -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data --restart=always portainer/portainer`
   + --restart=always : docker 재시작할 경우, portainer 컨테이너 또한 자동으로 구동하는 명령어 옵션
+ http://<해당서버ip>:9000 접속하면 첫 화면은 admin 계정 설정할 수있는 페이지가 뜸(초기접속정보 설정)
+ 관리하고자 하는 Docker 환경 선택 가능


... ing
