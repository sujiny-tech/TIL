# MySQL 사용자 
> 자주 사용하는 명령어는 기록하면서 점차 익숙해져야겠다 ❗

## 사용자 계정 추가
+ 먼저 MySQL 접속
   + `mysql -u root -p`   


+ 계정 전 현재 생성된 **사용자 계정 체크**
   + `use mysql;`   


   + `select host, user from  user;`
+ **사용자 추가**
   + `create user <user_id>;`
   + `create user <user_id>@localhost identified by <password>;` : password까지 설정   


      > ❗ 외부 접근 권한을 부여하려면, host부분을 '%'로 해서 똑같은 계정 추가해야 함



## 사용자 권한 부여

+ 모든 원격지에서 접속 권한 추가   

   + `grant all pribileges on <DB_name>.* to <user_id>@'%' identified by <password>;`   


+ 해당 사용자에게 모든 데이터베이스 모든 테이블에 권한 부여   

   + `grant all privileges on *.* to <user_id>@'%' identified by <password> with grant option;`

+ **권한 확인**
   + `show grants for <user_id>@localhost;`   


   + `show grants for <user_id>@'%';`
   
   
## 변경사항 반영(reload)    

+ **변경된 내용을 메모리에 반영**하기 위해 해당 명령어 사용   

   + `flush privileges;`   

➰ 해당 명령어는 성능에 영향을 준다고 함. 습관적으로 해당 명령어를 통해 reload 하는 경우가 있다고 하는데 이는 엄청난 부하가 된다고 함    

## 참고
+ 버전별 조금 차이 있음 👉 [다른사람의 블로그 참고](https://fruitdev.tistory.com/206)

