# JOSE(JSON Object Sigining and Encryption)
> JSON 데이터 포맷을 이용해 메시지에 대한 암호화를 지원하는 표준 규격

### JWT(JSON Web Token)
+ **stateless한 서비스**(상태정보가 유지되지 않은 서비스, http 프로토콜)에서 식별/인증/인가를 위한 모델
+ 속성정보(claim)을 JSON 데이터 포맷으로 표현한 토큰
+ 기술적으로는 서명되지 않은 토큰을 의미하지만, **일반적인 상황에서 JWS/JWE를 의미**한다고 함 (주로 JWS)   

### JWS(JSON Web Signature)
+ **인증 정보를 서버의 private key로 서명한 것**을 토큰화한 것
+ [JavaScript node-jws 모듈](https://github.com/auth0/node-jws) 참고 ✨   
+ [JavaScript node-jsonwebtoken 모듈](https://github.com/auth0/node-jsonwebtoken) 참고 ✨   

### JWE(JSON Web Encryption)
+ 서버와 클라이언트 간 **암호화된 데이터**를 토큰화한 것   

### JWK(JSON Web Key)
+ **JWE**에서 **payload encryption에 사용된 key**를 토큰화한 것 (key 자체 암호화 되어 있음)   

✨ [참고하기 좋은 사이트1 - jwt.io ](https://jwt.io/)
✨ [참고하기 좋은 블로그](https://ehdvudee.tistory.com/14)   
