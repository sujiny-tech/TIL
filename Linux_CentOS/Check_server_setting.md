# 서버 환경 체크
> 세팅을 확인할 일이 있어서 찾아보았고, 이전에 알고 있던 내용과 알게 된 내용을 간단하게 정리하였음 🛠 

## 메모리 확인
+ 터미널에 ```cat/proc/meminfo```를 입력하면, **메모리에 대한 상세 스펙을 확인**할 수 있음
+ 전체 메모리/ 사용중인 메모리/ 그 외 공유메모리, 버퍼, 캐시메모리/ swap 용량을 확인하려면, ```free```를 입력하면 됨   


   + free 뒤의 **옵션**들
      + ```free -h``` : (human-readable) 사람이 읽기 쉬운 단위로 출력해줌
      + ```free -b/-k/-m/-g``` : 원하는 단위(bytes/KB/MB/GB)로 출력해줌
      + ```free -t``` : ram 메모리 총 용량과 swap 메모리의 합(total)을 확인할 수 있음
+ **메모리 점유율 확인**은 ```top```입력 시 확인 가능함   

## HDD 용량 확인
+ `df -h` : HDD 용량 확인하는 명령어

## CPU 상세 확인
+ 터미널에 ```cat/proc/cpuinfo```를 입력하면, **CPU에 대한 상세한 정보를 확인**할 수 있음
+ **코어 수**는 ```nproc```을 입력하면 확인할 수 있음

## OS 확인
+ 터미널에 ```cat /etc/redhat-release```를 입력하면,**OS 정보를 확인**할 수 있음
+ **또다른 명령어**를 통해 OS 정보 확인 가능함
   +  ```cat /etc/*release*```
   +  ```cat /etc/issue```
   +  OS **bit** 확인 = ```getconf LONG_BIT```   

## port 체크
+ `netstat -nap` : 열려있는 모든 포트 확인하는 명령어
+ `netstat -nap | grep LISTEN` : Listen 중인 포트만 확인하는 명령어
+ `netstat -nap | grep <port_num>` : <port_num>에 해당하는 특정 포트의 상태만 확인하는 명령어

## port 열기
+ `sudo iptables -I INPUT 1 -p tcp --dport <port_num> -j ACCEPT` : <port_num>에 해당하는 포트를 open하는 명령어
  
## 네트워크 속도 확인
+ `yum install ethtool net-tools` : centos의 경우, ethtool 패키지 설치하는 명령어   
+ `ifconfig` : 네트워크 인터페이스 확인
+ `ethtool <net_interface_name>` : ethtool을 통해 원하는 인터페이스 정보 조회   

## 날짜/시간 확인
+ `date` : 서버의 시간 출력
+ `timedatectl` : 정보확인(local time, universal time, Timezone ...etc)
+ 서버 시간 변경
   + `ls /usr/share/zoneinfo/Asia` : Seoul의 timezone 정보에 해당하는 파일 찾기
   + `mv /etc/localtime /etc/localtime_org` : 기존 파일백업한 다음, 해당하는 심볼릭 링크 만들기
   + `ln -s /usr/share/zoneinfo/Asia/Seoul /etc/localtime` 

# 참고
+ [참고하면 좋을 다른사람 블로그✨ -Centos7 시스템 자원/네트워크 관련 정리](https://estenpark.tistory.com/372)
