# Error list 정리
> Linux 사용하면서 발생했던 에러들을 정리하고, 간단한 해결 방법을 정리하자 ➰

+ -bash : syntax error near unexpected token `('
   + 중간에 괄호부분이 에러 원인인데, 괄호부분에 대해 쌍따옴표로 감싸주면 에러 해결 ❗  


# 명령어 정리
> 사용했던 것들 생각날때마다 기록하자 📝

+ linux 사용자 변경 (CentOS)
   + cli 환경에서 `su - [user_name]` 명령어 입력

+ 비밀번호 변경 
   1. cli 환경에서 `passwd` 입력
      + 현재 비밀번호 입력 후, 변경할 비밀번호 입력하면 끝 ❗

+ 현재 디렉토리 정보 출력(print working directory)
   + cli 환경에서 `pwd` 입력

+ 파일 생성, 날짜 변경
   + `touch ,filename]` : 파일 생성
   + `touch -c [filename]` : 파일에 관한 시간정보를 현재시간으로 변경
   + `touch -t YYYYMMDDhhmm [filename]` : 시간정보를 YYYYMMDDhhmm(년/월/일/시/분)으로 변경

+ 사용자 추가
   + `sudo useradd -m [user_name]` : 사용자 추가    
   + `sudo usermod -a -G sudo [user_name]` : 해당 사용자가 sudo 명령어 사용할 수 있도록 변경   

+ 서버 내 테스트시 발생했던 에러 `too many open files`    

   + 서버 환경 설정해야 됨 (openfiles 값 설정필요) 
   + ulimit 명령어의 -n 옵션을 통해 설정 가능
      + `ulimit -n 40960`
   + /ext/security/limits.conf 파일에 추가/수정을 통해 설정 가능
      + `<계정명> soft nofile 40960` 
      + `<계정명> hard nofile 40960`  
   + [다른 사람의 블로그 참고👍](https://knoow.tistory.com/220)   


ing...

