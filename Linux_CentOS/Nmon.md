# Nmon (Nigel's Monitor)
> 서버의 CPU, memory, DiskI/O를 간편하게 볼 수 있는 모니터링 툴 🖥✨

+ AIX(UNIX 시스템), Linux의 성능을 분석해주는 **모니터링 툴**

+ 기본적으로 터미널에서 top 명령어를 통해 확인할 수 있으나, Nmon은 결과치들을 통계로 분석해주는 Nmon Anaylser가 존재함

+ 해당 분석 툴을 통해 나오는 그래프들을 보면서 조금 더 쉽게 직관적으로 보기 위해 사용하게 되었음 📊


## Nmon 설치
+ Centos에서 설치 : `yum install nmon`
   + 설치가 안되는 경우, EPEL(Extra Packages for Enterpriese Linux) repository를 설치한 다음에 nmon 패키지 설치 명령어를 입력하면 됨 ❗❗
      + `yum install epel-release` : epel repository 설치 명령어   


+ 사용방법

   + 실시간 모니터링 : `nmon` 명령어 입력
   + 데이터 캡쳐 모드 : `nmon -f -s<sec> -c<number>`

      > 데이터 캡쳐 모드 시 사용하는 옵션들은 아래에서 설명 ➰

## Nmon 캡쳐/저장
+ 주요 옵션들
   + -f : 파일로 저장
   + -s<sec> : 몇 s당 데이터를 캡쳐할 것인지
   + -c<number> : 캡쳐 횟수 (default : 10000000)
   + -m<dir> : 파일 저장할 경로   
   
+ `nmon -f -s10 1800` : 10초 간격으로 1800번 캡쳐한 결과들을 파일로 저장 (30분 소요)
   
# Nmon Anaylzer
> 공식 사이트에서는 Nmon Anaylser라고 명칭이 되어있는데...오타인가?🤷‍♀️   
   
+ Nmon을 통해 출력된 파일은 데이터들(CPU/메모리 사용량/네트워크 등등)을 포함하는데, 출력된 파일들을 분석 툴에 입력하면 간편하게 통계 그래프가 나옴 
   
+ 구체적으로 cpu 사용량, 메모리 사용량, disk I/O, 네트워크, 프로세스, 각 cpu마다의 사용량 등을 그래프로 볼 수 있음 ❗   
   
+ 금방 설치하고 편리하게 사용할 수 있어서 좋았고, 측정 값들을 그래프로 빠르게 뽑아낼 수 있어서 분석하기에 정말 유용했음 ⭐

   
# 참고
+ [Nmon Anaylzer 💫](https://developer.ibm.com/articles/au-nmon_analyser/)
