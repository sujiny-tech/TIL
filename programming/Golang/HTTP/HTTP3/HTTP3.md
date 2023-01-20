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


+ HTTP/3 - UDP를 사용하여 개선된 점
   > ✔️ UDP의 커스텀 기능을 이용해서 기존 TCP 기능을 모두 구현할 수 있다고 함(TCP와 비슷한 수준의 기능을 가질 수 있고, 핸드쉐이크를 사용하지 않을 수 있다는 것을 뜻함)   
   > ✔️ 내가 이해하기로는 UDP을 장점(속도)을 살리되 기존 TCP의 기능을 구현하여 (안정성을) 보완했다는 것으로 이해했음
   + **[개선점1]** 연결 설정 시, 레이턴시 감소되는 점
      + '클라이언트 요청 -> 서버처리 -> 클라이언트로 응답'의 한 사이클을 1RTT(Round Trip Time)   
   + **[개선점2]** 패킷 손실 감지에 걸리는 시간을 단축한 점
      + TCP에서 '패킷 전송 -> 타임아웃 -> 패킷 재전송 -> ACK 받음'으로 통신 진행하는데, QUIC에서는 ACK가 어떤 패킷에서 온건지 확인하기 위해, 헤더에 패킷번호에 대한 공간을 부여했다고 함

   + **[개선점3]** 멀티플렉싱 지원
      + 멀티플렉싱 : 하나의 통신 채널을 통해 다중적으로 모두 처리 하는 것 (즉, 하나의 서버에서 여러 클라이언트를 모두 처리한다는 것)
      + 중간에 어느 스트림의 패킷 손실되어도, 다른 스트림에 지장없다는 점에서 TCP의 단점인 HOLB(Head of Line Blocking) 방지 가능하다고 함

   + **[개선점4]** 클라이언트 IP와 무관하게 연결 유지
      + 랜덤한 ConnectionID를 사용하므로, 클라이언트 IP와는 상관없이 기존 연결을 유지할 수 있음(핸드쉐이트 과정 생략 가능하다는 말과 동일함)



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


## Chrome DevTools Protocol
+ 크롬 최신버전에서 '도구 더보기' → '개발자 도구'를 눌르면 웹에서 발생하는 통신에 대한 프로토콜 확인 가능함
   + 단축키 : Ctrl + Shift + I 누르면 개발자 도구 뜸
   + 프로토콜이 안보인다면, 상태 필드에서 마우스 오른쪽 클릭하여 '프로토콜' 항목을 체크해주면 확인 가능
   + 프로토콜 값이 'h3'가 HTTP3임을 나타냄
      + [Chrom Platform Status 웹사이트에서 h3 언급 (2022-10-18)](https://chromestatus.com/feature/5154357283651584)
      + 위의 링크 내에 'Specification'을 참고하면, 해당 워킹 그룹에서 작성한 스펙에서 HTTP3를 h3라고 지칭하는 것을 확인 가능함   



**... ing**


# 참고
+ [Wiki: HTTP/3](https://en.wikipedia.org/wiki/HTTP/3)
+ [golang quic-go](https://pkg.go.dev/github.com/lucas-clemente/quic-go#section-readme)
