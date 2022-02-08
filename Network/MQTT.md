# MQTT
> Message Queuing Telemetry Transport의 약자로 IoT처럼 제한된 또는 대규모 트래픽 전송을 위해 만들어진 프로토콜임

## 동작 원리
+ **Producer, Broker, Consumer** 대상이 존재함.
+ 메시지를 해당 Topic에 배포하고, 그 topic에 Consumer에게 topic에 대한 data를 전달함.
   + Producer는 Broker에게 특정 Topic에 대해 해당하는 data와 같이 메세지를 전송하고,   
   + Broker는 Producer가 배포(발행)한 topic을 가지게 되며,
   + Consumer는 원하는 topic에 대한 data를 전달받게 됨.   

+ **Broker**는 **나머지 컴포넌트 간의 통신을 조정하는 역할**을 수행함 ✨

> <img src="https://user-images.githubusercontent.com/72974863/152911115-35427e47-3377-400e-984b-ae5c5da9140b.png">   

