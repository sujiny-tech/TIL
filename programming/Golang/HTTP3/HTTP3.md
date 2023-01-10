# HTTP/3
> HTTP 3번째 메이저 업데이트 버전으로 22.06.06 IETF RFC9114 표준화됨

# HTTP/3 특징
+ 이전버전에서는 TCP로 통신하는 반면, HTTP/3은 UDP 기반의 QUIC(Quick UDP Internet Connection)프로토콜을 사용하여 통신함
   > OSI 7계층 중 4번째 전송계층에서 사용되는 프로토콜인 TCP, UDP의 차이는 데이터 전송 보장유무   
   > ✔️ TCP는 연속성보다 신뢰성 있는 데이터 전송을 보장하는 프로토콜   
   > ✔️ UDP는 TCP보다 빠르고 네트워크 부하가 적다는 장점이 있지만, 신뢰성있는 데이터 전송을 보장하지 않음(신뢰성보다 연속성이 중요한 실시간 스트리밍과 같은 서비스에 자주 사용됨)   

+ HTTP/3과 이전버전의 차이(하단 '참고'의 'Wiki HTTP/3'에서의 그림)
   > <img src="https://user-images.githubusercontent.com/72974863/210688335-c6ce67ff-0b82-459b-9e0e-53e95f26f3ce.png" width=50% height=50%>  
   > ✔️ HTTP/2보다 가장 큰 장점은 Zero RTT(Round Trip Time), 패킷 손실에 대한 빠른 대응, 사용자 IP가 바뀌어도 연결 유지된다는 것

# HTTP/3 지원

### 웹 브라우저
+ Chrome, Edge, Firefox, Safari 등

### 서비스
+ CDN(Content Delivery Network)서비스 관련된 부분에서 적용하고 있다고 함

### 웹서버
+ Ngnix에서 개발 중인 버전 체험가능하다고 함
+ Caddy 웹서버, LiteSpeed 웹서버

### HTTP/3가 적용된 사이트
+ 구글(유튜브를 포함한 대부분의 서비스에 적용됨), 네이버, 페이스북(인스타그램), 넷플릭스, 줌 비디오 커뮤니케이션   


### Chrom DevTools Protocol
+ 크롬 최신버전에서 '도구 더보기' → '개발자 도구'를 눌르면 웹에서 발생하는 통신에 대한 프로토콜 확인 가능함
   + 단축키 : Ctrl + Shift + I 누르면 개발자 도구 뜸
   + 프로토콜이 안보인다면, 상태 필드에서 마우스 오른쪽 클릭하여 '프로토콜' 항목을 체크해주면 확인 가능
   + 프로토콜 값이 'h3'가 HTTP3임을 나타냄
      + [Chrom Platform Status 웹사이트에서 h3 언급 (2022-10-18)](https://chromestatus.com/feature/5154357283651584_
      + 위의 링크 내에 'Specification'을 참고하면, 해당 워킹 그룹에서 작성한 스펙에서 HTTP3를 h3라고 지칭하는 것을 확인 가능함   



**... ing**


# 참고
+ [Wiki: HTTP/3](https://en.wikipedia.org/wiki/HTTP/3)
+ [golang quic-go](https://pkg.go.dev/github.com/lucas-clemente/quic-go#section-readme)
