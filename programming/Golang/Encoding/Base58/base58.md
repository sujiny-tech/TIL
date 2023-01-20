# Base58
> 기존 base64의 경우, 혼동되는 문자들(0, O, I, l) 그리고 query를 방해하는 문자(=,+...)가 있다는 단점을 보완하기 위해 나온 인코딩 기법 📝

+ 처리방식
   + base64
      + 64진수 체계의 인코딩 기법
      + 입력된 값을 6bit씩 구분하여, 지정된 table의 charater로 치환   
   + base64url
      + base64에서 62,63번째에 해당하는 +, / 가 데이터 컨트롤시 오류를 일으킬 가능성이 있음
      + 따라서, url 사용시 문제 발생을 막기 위해 -, _로 대체된 것
   + **base58**
      + 입력된 값을 big number로 변경하여 58로 나눠, 지정된 table의 character로 치환



# 참고 ✨
+ [해시넷 Base58](http://wiki.hash.kr/index.php/%EB%B2%A0%EC%9D%B4%EC%8A%A458)   
+ [github 다른 사람의 소스 참고⭐](https://github.com/Akers9560/java-base58-prefix)   


