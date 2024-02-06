# Errorlist
> mysql 사용하면서 발생했던 에러 리스트들..😂

+ "too many connections"     

   + 너무 많은 connection이 유지되어서 발생하는 에러 (초기설정값은 100로 설정되어있다고 한다.)   

   + 또한, wait_timeout의 초기 설정값은 288000(s)로 너무 긴 편
   + 따라서 max_connections 값을 변경해줘야되며, wait_timeout 설정 값도 낮춰줘야 함
   + wait_timeout을 설정했는데도 mysql 명령어를 쳐서 확인해도 적용이 안되어있는 경우, `show global variables like '%wait_timeout%';`을 입력하여 확인해볼 것 ❗❗❗
      > [도움이 된 다른사람의 블로그👍](https://velog.io/@army262/mysql-waittimeout)

+ "sql: no rows in result set"
   + 특정 조건에 따라 row 조회할 때, 발생한 에러
   + 해당 조건에 부합한 row가 존재하지 않을 때 발생함

+ **"Error 1054 Unknown column in where clause"**
   + query를 날릴 때, sql문에 조건문 부분을 ~ 'WHERE ='+<조건변수> 로 처리 했을 때 발생하는 에러
   + ~ WHERE = '%s'",<조건변수>로 처리해줘야 함 (즉, 변수 앞뒤로 작은 따옴표 '가 있어야 함, 작은 따옴표가 문자열의 구분 기호❗)
   + [stackoverflow 참고](https://stackoverflow.com/questions/61848379/error-1054-unknown-column-in-where-clause)
   + 추가 : only 숫자인 경우 문제 발생이 없으나, text로 시작된 경우 해당 에러가 발생함

+ "sql: expected 0 arguments, got 1"
   + golang으로 db.Exec 함수 사용해서 DELETE sql문 수행시 발생
   + db.Exec 함수에 전달하는 sql문 형식이 잘못되서 발생한 에러로 보임
      + ex) 에러발생o : db.Exec("delete from tbl_example where id = '%s' and name in ('sujin', 'sj', 'sujiny2222')", id)
      + ex) 에러발생x : db.Exec("delete from tbl_example where id = (?) and name in ('sujin', 'sj', 'sujiny2222')", id)

+ **MySQL Workbench 사용 중 동작안하는 이슈 : Error while loading image**
   + 방법1. window기준 제어판>프로그램 제거 창에서 MySQL Workbench 8.0 CE 마우스 우클릭 --> "복구" 버튼 클릭
   + 방법2. mysql installer에서 remove 후, 재 설치(프로그램 폴더가 제대로 지워지지 않은 경우, 설치된 프로그램 폴더를 삭제해야 함)
      > [도움이 된 stackoverflow](https://stackoverflow.com/questions/44690836/how-to-solve-mysql-error-which-occurs-while-loading-image)

+ Error : ER_TRUNCATED_WRONG_VALUE : Truncated incorrect DOUBLE value
  + 값의 타입이 맞지 않은 경우, 구문 오류 존재 시 발생
      + ex) 값의 타입이 INT로 정의되어있는데, VACHAR를 입력한 경우
      + ex) update 명령어 입력 시, 콤마(,)로 묶어줘야하는 부분을 and로 묶어준 경우
        

+ **ing**


# 참고 💫
+ [Mysql "too many connections" 관련 다른사람의 블로그 👍](https://plogger.tistory.com/entry/MySQL-Too-many-connections-Max-Connection-%EC%A1%B0%EC%A0%95)
