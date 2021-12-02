# Errorlist
> mysql 사용하면서 발생했던 에러 리스트들..😂

+ "too many connections"
   + 너무 많은 connection이 유지되어서 발생하는 에러 (초기설정값은 100로 설정되어있다고 한다.)
   + 또한, wait_timeout의 초기 설정값은 288000(s)로 너무 긴 편
   + 따라서 max_connections 값을 변경해줘야되며, wait_timeout 설정 값도 낮춰줘야 함


# 참고 💫
+ [Mysql "too many connections" 관련 다른사람의 블로그 👍](https://plogger.tistory.com/entry/MySQL-Too-many-connections-Max-Connection-%EC%A1%B0%EC%A0%95)
+ 
