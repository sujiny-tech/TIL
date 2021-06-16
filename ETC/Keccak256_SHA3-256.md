# Keccak256 vs SHA3-256
+ **Keccak**
   + 이전 SHA1, SHA2의 충돌 가능성을 피하기 위해 전혀 다른 알고리즘의 개발이 진행되었음.
   + NIST에서 SHA3 해시 알고리즘으로 선정된 알고리즘.   
   
+ **차이점**   
   + SHA3-256(M) := Keccak256(M'), **∵ M'= M||01**
   + 즉, '01'은 2-bit suffix   
   
✨ [참고 : 해시넷 - SHA3](http://wiki.hash.kr/index.php/SHA3)
✨ [참고 : 유튜브](https://www.youtube.com/watch?v=mBw8kVaPngc)
