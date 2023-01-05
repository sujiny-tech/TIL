# HTTP/3
> HTTP 3번째 메이저 업데이트 버전으로 22.06.06 IETF RFC9114 표준화됨

# HTTP/3 특징
+ 이전버전에서는 TCP로 통신하는 반면, HTTP/3은 UDP 기반의 QUIC(Quick UDP Internet Connection)프로토콜을 사용하여 통신함
   > OSI 7계층 중 4번째 전송계층에서 사용되는 프로토콜인 TCP, UDP의 차이는 데이터 전송 보장유무   
   > ✔️ TCP는 연속성보다 신뢰성 있는 데이터 전송을 보장하는 프로토콜   
   > ✔️ UDP는 TCP보다 빠르고 네트워크 부하가 적다는 장점이 있지만, 신뢰성있는 데이터 전송을 보장하지 않음(신뢰성보다 연속성이 중요한 실시간 스트리밍과 같은 서비스에 자주 사용됨)   

+ 참고에서 wiki 링크의 그림(HTTP버전별 비교한 그림)참고하면 좋을 듯함   


**... ing**


# 참고
+ [Wiki: HTTP3](https://en.wikipedia.org/wiki/HTTP/3)
+ [golang quic-go](https://pkg.go.dev/github.com/lucas-clemente/quic-go#section-readme)
