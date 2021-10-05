# MySQL 설치 및 Workbench 연동

## Windows 10 설치
+ www.mysql.com/downloads 사이트 내에서 "MySQL Community Serve" 다운로드
   + 무료 = Community (GPL;General Public License)   

+ 로그인 필수 아니므로 중간에 'No thanks' 클릭 후 다운하고, 기본 설정대로 'Execute', 'Next' 클릭

+ 기본 포트로 3306으로 설정되어있는데, 기존에 사용하는 포트와 충돌나지 않도록 잘 설정해야함 
   + 되도록 8080, 8081, 1521(Oracle) 등은 피해서 설치하는게 좋을 것 ❗❗

+ 비밀번호 설정도 빠짐없이 잘해주고, 엄격한 보안을 위해 잘 설정하고 잊지 않도록..! ➰

## CentOS 7

+ yum을 통해 MySQL 설치 : `sudo yum install mysql-community-server`
+ 설치 완료되면 MySQL 서비스 시작하고, 아래를 통해 부팅시 자동으로 시작할 수 있도록 설정
   + `sudo systemctl enable mysqld`
   + `sudo systemctl start mysqld`   

+ MySQL 서버가 처음으로 시작되면 root 사용자에 대해 임시 password가 생성됨. 다음 명령어를 통해 해당 임시 password를 찾을 수 있음 ❗ 
   + `sudo grep 'temporary password' /var/log/mysqld.log`

+ MySQL_secure_installation 명령을 통해 Mysql 설치의 보안을 향상시킬 수 있음   
   + `sudo mysql_secure_installation`
   + 해당 명령어 입력시, 익명 사용자 제거/로컬 컴퓨터에 대한 루트 사용자 엑세스 제한/테스트 데이터 베이스 제거에 관한 메시지를 띄우는데, 나는 모든 질문에 y를 입력했음.

## Workbench 연동

+ MySQL 설치 시 기본 설정으로 진행했다면, Workbench 설치가 되었을 것임. 혹시 설치가 되지 않았더라면 MySQL installer 에서 Workbench를 add시켜주면 됨 ❗
+ Workbench 설치 후, 아래 그림과 같이 순차적으로 진행해주면 됨 
   > <img src="https://user-images.githubusercontent.com/72974863/136013332-60285143-6306-403f-becc-5c9ea11a437f.png" width="70%" hegiht="70%">   


+ 나는 위에서 설치한 Windows(local)와 CentOS 7서버의 MySQL을 연동해서 아래와 같이 연결된 것을 확인할 수 있음
   > <img src="https://user-images.githubusercontent.com/72974863/136013855-cb63fa12-76bd-4609-97dd-0e09bf8601f6.png" width="70%" hegiht="70%">
