# Error list 정리
> Linux 사용하면서 발생했던 에러들을 정리하고, 간단한 해결 방법을 정리하자 ➰

+ -bash : syntax error near unexpected token `('
   + 중간에 괄호부분이 에러 원인인데, 괄호부분에 대해 쌍따옴표로 감싸주면 에러 해결 ❗  

+ 서버 내 테스트시 발생했던 에러 : **too many open files!!!**

   + 서버 환경 설정해야 됨 (openfiles 값 설정필요) 
   + ulimit 명령어의 -n 옵션을 통해 설정 가능
      + `ulimit -n 40960`
   + /ext/security/limits.conf 파일에 추가/수정을 통해 설정 가능
      + `<계정명> soft nofile 40960` 
      + `<계정명> hard nofile 40960`  
   + [다른 사람의 블로그 참고👍](https://knoow.tistory.com/220)   

+ 서버 내 테스트시 발생했던 에러 : **no space left on device**   

   + 서버 용량 부족으로 인해 발생
   + 명령어를 통해 용량 체크
      + `dh -h` 
      + `dh -i` : Inodes 값 체크 --> Inodes 값이 100%임을 확인할 수 있음
         > 이때, inode는 파일 시스템에서 사용하는 자료구조이며 정규 파일 디렉터리 등 파일 시스템에 관련된 정보를 가짐.    
         > 각 파일은 하나의 Inode를 가지고, 소유자 그룹, 권한, 파일 형태, Inode 숫자 등 파일에 관한 정보를 가짐.
         > inode에 할당된 공간은 파일 시스템 전체의 1% 정도임.


      + `for i in /*; do echo $i; find $i |wc -l; done` : / 파티션 중 어느 부분에서 inodes를 많이 사용했는지 체크
      + `for i in <해당 path>; do echo $i; find $i |wc -l; done` : 해당 path에서 어느 부분에서 inodes를 많이 사용했는지 체크 가능
   + 해당 폴더 내에서 필요없는 부분을 삭제하거나, 또는 용량이 큰 폴더로 이동시키는 방법이 있음 ❗
   

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

+ 서버 재부팅
   + `reboot` : 로그를 남기고 재부팅

+ 시스템 재부팅
   + `shutdown -r now` : 지금 시스템 재부팅

+ 삭제
   + `rm <삭제할 파일/폴더 등>` : 해당 파일/폴더 삭제
   + `rm -r <삭제할 파일/폴더 등>` : cannot remove '폴더/파일명' 에러 발생시 옵션 넣어서 명령어 입력

+ usb 마운트
   + `fdisk -l` : 디스크 정보 출력하는 명령어를 통해, usb 인식되었는지 확인 (usb와 같은 용량을 가진 정보 찾기)
   + `mount -t vfat <디스크명> <마운트할 폴더>` : usb에 해당하는 디스크를 해당 폴더에 마운트 시킴
   + `umount <마운트한 폴더>` : 마운트 한 폴더 umount



ing...

