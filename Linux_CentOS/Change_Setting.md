# 서버 환경변수 변경
> load test 진행 중, 서버 환경변수를 변경해줘야 할 일이 발생해서 정리해본다✏  

## ulimit
+ 프로세스의 자원의 limit를 설정하는 명령으로, soft/hard 한도로 나뉨   

+ soft의 경우, 새로운 프로그램 생성 시 default로 적용되는 한도를 뜻하고, hard는 soft의 최대 한도를 뜻함    

+ 서버의 ulimit 조회
   + `ulimit -a` : soft 한도 출력 (-a를 -aS으로 입력시 동일)
   + `ulimit -aH` : hard 한도 출력


## [ERR] too many open files 
> load test 진행하다가 접했던 에러🚨

+ 발생 원인
   + OS에 요청할 수 있는 최대 Open file 갯수의 제한을 초과해서 발생한 에러   
   + 즉, 해당 서버에서 설정된 max open files limit 초과해서 발생  


+ 해결 방법
   + 서버 내 max-files를 체크한 뒤, 이를 초과하지 않은 범위내에서 max open files의 한도를 변경하면 됨
   + **... ing**




## 그 외
