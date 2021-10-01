# Nmon (Nigel's Monitor)
> 서버의 CPU, memory, DiskI/O를 간편하게 볼 수 있는 모니터링 툴 🖥✨

+ AIX(UNIX 시스템), Linux의 성능을 분석해주는 모니터링 툴
+ 기본적으로 터미널에서 top 명령어를 통해 확인할 수 있으나, Nmon은 결과치들을 통계로 분석해주는 Nmon Anaylser가 존재함
+ 해당 분석 툴을 통해 나오는 그래프들을 보면서 조금 더 쉽게 직관적으로 보기 위해 사용하게 되었음 📊

## nmon 캡쳐/저장
+ 주요 옵션들
   + -f : 파일로 저장
   + -s<sec> : 몇 s당 데이터를 캡쳐할 것인지
   + -c<number> : 캡쳐 횟수 (default : 10000000)
   + -m<dir> : 파일 저장할 경로

# Nmon Anaylser

+ Nmon을 통해 출력된 파일은 데이터들(CPU/메모리 사용량/네트워크 등등)을 포함하는데, 출력된 파일들을 분석 툴에 입력하면 간편하게 통계 그래프가 나온다.
+ ... **ING**

   
# 참고
+ [Nmon Anaylzer 💫](https://developer.ibm.com/articles/au-nmon_analyser/)