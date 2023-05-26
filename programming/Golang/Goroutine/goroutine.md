# Goroutine
> 비동기작업을 효율적으로 처리할 수 있는 스레드 라이브러리

## 개념
+ 스레드(Thread)는 **프로그램내에서 실행 흐름**을 나타냄.
+ 고루틴(Goroutine)은 Go에서 사용하는 **경량 스레드**를 의미함.
 
+ 일반적으로 다른언어에서는 스레드 생성 시 1MB가 필요하고, 커널레벨 스레드와 연결되어 운영체제에서 스케줄링하여 돌아가는 방식인 반면,    
  Goroutine은 **2KB만 필요**하고 커널에서 직접 스케줄링을 받지 않고 **GO 스케줄러가 직접 스케줄링해주므로 훨씬 가볍고 빠르다고 함**
  
## 사용 방법
+ 적용하고 싶은 함수 호출 부분 앞에 간단하게 "go"만 붙여주면 고루틴으로 실행됨
+ 익명함수에서도 고루틴으로 실행 가능함

+ 고루틴도 자식 스레드의 작업이 남아있어도 부모 스레드/프로세스가 종료되면 자동 종료됨
   + 일반적인 운영체제에서 자식 스레드는 부모 스레드가 종료되면 남은 작업에 상관 없이 종료됨. 이처럼 고루틴도 동일함.
   + 자식 스레드의 작업이 모두 다 처리될 때까지 부모 스레드를 대기시키려면, WaitGroup 라이브러리 사용하면 됨🌟
      + Mutex Lock과 사용용도는 비슷함

## 그외
+ ```func (*Once) Do(f func())``` : 함수 한번만 실행
   + Once를 통해, goroutine안에서 특정함수를 한번만 실행 가능함   

   + 따라서, **복잡한 반복문 안에서 초기화할때 유용**함   

  ```
  //example. sync.Once
  func printhello() {
     fmt.Println("Hello, world!")
  }
  func main() {
  
     once :=new(sync.Once)
  
     for i:=0; i<5; i++ {
        go func(n int) {
           fmt.Println("iter : ", n)
           once.Do(printhello) //한번만 실행할 함수명 입력하면 됨.
        }(i)
     }
  }
  ```
  
   
[👉Goroutine 예제 - goroutine_example.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/Goroutine/goroutine_example.go)   
