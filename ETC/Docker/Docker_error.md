# Docker 에러/해결방법

+ `Got permission denied while trying to connect to the Docker daemon socket at unix : ~~~~ dial unix /var/run/docker.sock:connect:permission denied`
   + root계정이 아닌 다른 계정에서 Docker cmd 입력 시 발생하는 에러
   + 원인 : 해당 계정에 대해 권한이 없음
   + 방법1: Docker.sock 권한 변경
      + `sudo chmod 666 /var/run/docker.sock`
   + 방법2: Docker 그룹 생성 후, 사용자 추가
      + `sudo groupadd docker`
      + `sudo usermod -aG docker $USER`
      + `newgrp`

...ing


### 참고
+ [권한이슈 관련해서 참고한 다른사람 깃허브✨]([https://jjunii486.tistory.com/234](https://github.com/occidere/TIL/issues/116))

