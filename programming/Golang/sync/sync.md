# Sync
> go에서 동기화를 위한 표준 라이브러리, goroutine의 동기화처리를 위해선 꼭 사용되는 패키지임!

## WaitGroup
> goroutine만 사용한다면 main이 종료될 때 대기없이 바로 같이 종료하게 됨   
> 이를 방지하기 위해, **지정한 개수의 goroutine이 완료될 때까지 기다릴 때 사용**함

+ func (wg * WaitGroup) Add(delta int) : 대기 그룹에 **고루틴 개수 추가**해주는 함수
+ func (wg * WaitGroup) Done() : **goroutine이 끝났음을 알려줄 때 사용**하는 함수
   + Add 함수에 설정한 값과 Done 함수가 호출되는 횟수는 동일해야 함
   + 동일하지 않을 시, 패닉 발생함
   + defer() 함수와 함께 사용해서 goroutine이 끝나기 직전에 wg.Done() 함수 호출하도록 사용 가능
+ func (wg * WaitGroup) Wait() : **모든 goroutine이 끝날 때까지 기다리도록** 하는 함수
   
... ing

## Mutex

## RWMutex

## Once

## Pool

