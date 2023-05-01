# Channel
> **고루틴(goroutine)간의 통신 채널**로, 고루틴끼리 메시지를 전달할 수 있는 메시지 큐(Thread-safe Queue)🌟

## 개념
+ 여러 개의 고루틴을 실행했을 때, 서로 간의 정보를 공유할 수 있도록 해줌   
   <img src="https://user-images.githubusercontent.com/72974863/234153214-7de0bbc2-0015-43bd-9e93-421acc039168.png">   
+ 채널은 type이 있는데 int, string, err 등이 있음(주고 받을 데이터의 타입에 따라 채널의 타입도 달라짐)

## 채널 생성
### 1. 채널 인스턴스 생성
+ 슬라이스, 맵 등과 같이 ```make()```함수를 사용하여 채널 인스턴스 생성, 채널을 의미하는 ```chan```키워드 사용해야 함
   + ```channel_first := make(chan int)``` : int 타입의 데이터를 송수신할 채널 인스턴스 생성
   + ```var channel_seconde chan string = make(chan string)``` : string 타입의 데이터를 손수신할 채널 인스턴스 생성   
+ 일반적으로 채널 생성하면, 크기가 0인 채널이 만들어 짐. 채널에 들어온 데이터를 담아둘 곳이 없다는 것을 나타내고, 데이터를 빼갈 때까지 대기함.    
  --> 채널에서 데이터를 가져가지 않아서 프로그램이 멈추는 경우 발생할 수 있음.   
  
### 2. 버퍼를 가진 채널 인스턴스 생성
+ 기존에 슬라이스, 맵을 할당할때 ```make()``` 함수를 사용한 것처럼, 두번째 파라미터에 버퍼 크기를 적어주면 됨!
   + ```var chan string channel_third = make(chan string, 2)``` : string 타입의 데이터를 2개를 담아낼 수 있는 채널 인스턴스 생성   
+ 1과 마찬가지로, 버퍼가 다 차면 빈자리가 생길 때까지 대기함. 따라서 데이터를 빼내지 않으면 버퍼가 없을 때처럼 고루틴이 멈추는 경우 발생함.   


## 채널 송수신
### 1. 채널에 데이터 넣기(송신)
+ ```<-``` 연산자를 이용하여 채널에 데이터 넣음
   + ```channel_first <- 1000``` : channel_first라는 채널에 1000 데이터를 넣음   
   + ```channel_second <- "this is a message"``` : channel_second라는 채널에 "this is a message" 데이터 넣음  


### 2. 채널에 데이터 빼기(수신)
+ ```<-``` 연산자를 이용하여 채널로부터 데이터를 뺌
   + ```var checkInt int = <- channel_first``` : channel_first에 담긴 데이터를 checkInt 변수로 빼냄
   + ```checkString := <- channel_second``` : channel_seconde에 담긴 데이터를 checkString 변수로 빼냄

## Select문
+ 채널에서 데이터가 들어오는 것을 대기할 때, 데이터가 들어오지 않으면 다른 작업을 하도록 처리하거나 또는 여러 채널을 동시에 대기하고 싶은 경우 에 select문 사용❗
+ 여러 채널을 동시에 대기 가능함. 하나의 채널이라도 데이터를 읽어오면 해당 구문을 실행하고 select문 종료됨. 
+ 하나의 case에 대해서 처리되기 때문에, 반복해서 데이터를 처리하고 싶다면 for문도 같이 사용해야 함

## 채널 닫기
+ ```close()```함수를 사용하여 채널 닫음. 채널을 닫으면 해당 채널로 더 이상 송신 불가능 하나, 채널에 데이터가 있다면 수신은 가능함  


   
   
## 예제

#### [👉Select문 활용한 예제 1 - channel_example1.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/Channel/channel_exmple1.go)
#### [👉Select문 활용한 예제 2 - channel_example2.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/Channel/channel_exmple2.go)
#### [👉Send, Receive 예제 - channel_example3.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/Channel/channel_exmple3.go)      




## 참고
+ [예제로 배우는 Go 프로그래밍💫](http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90)

