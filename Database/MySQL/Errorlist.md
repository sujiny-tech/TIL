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

+ "Error 1054 Unknown column in where clause"
   + query를 날릴 때, sql문에 조건문 부분을 ~ 'WHERE ='+<조건변수> 로 처리 했을 때 발생하는 에러
   + ~ WHERE = '%s'",<조건변수>로 처리해줘야 함 (즉, 변수 앞뒤로 작은 따옴표 '가 있어야 함, 작은 따옴표가 문자열의 구분 기호❗)
   + [stackoverflow 참고](https://stackoverflow.com/questions/61848379/error-1054-unknown-column-in-where-clause)


+ **ing**


# 참고 💫
+ [Mysql "too many connections" 관련 다른사람의 블로그 👍](https://plogger.tistory.com/entry/MySQL-Too-many-connections-Max-Connection-%EC%A1%B0%EC%A0%95)
