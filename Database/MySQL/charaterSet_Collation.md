# Character Set
> mysql에서 schema 생성 시 설정하는 Charset과 Collation 설정에 관해 정리   

+ Character Set : 문자와 인코딩의 집합
+ 대부분 사용하는 Character set
   + `utf8`
   + `euckr`

+ **...ing**

# Collation
> mysql은 다른 DB와 다르게, collcation 속성을 가지고 있다고 함   

+ Collation : character set 안에서 character들을 비교하기 위한 규칙(rule)들의 정의

+ 대부분 사용하는 collation
   + `binary collation`
   + `case-insensitive collation`   


+ 다른 나라 언어들을 사용하는 경우, 문자집합을 어떻게 사용할것인지 collcation 속성을 통해 설정함

+ **...ing**

## 참고 ✨
+ [잘 정리하신 다른 사람의 블로그 ⭐](https://mysqldba.tistory.com/154)
+ [인코딩에 관련 다른 사람의 블로그](https://kamang-it.tistory.com/245)
