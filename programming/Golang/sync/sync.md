# Sync
> go에서 동기화를 위한 표준 라이브러리, goroutine의 동기화처리를 위해선 꼭 사용되는 패키지임!

## 동기화(Synchronization) 객체 
+ **Mutex** : 뮤텍스, 여러 스레드에서 공유되는 데이터를 보호할 때 사용
+ **RWMutex** : 읽기/쓰기 뮤텍스, 읽기 동작과 쓰기 동작을 나누어 Lock 걸 수 있음
+ **Cond** : 조건 변수, 대기하고 있는 하나의 객체를 깨우거나 여러 개를 동시에 깨울 수 있음
+ **Once** : 특정 함수를 딱 한번만 실행할 때 사용
+ **Pool** : 멀티 스레드(goroutine)에서 사용할 수 있는 객체 풀, 자주 사용하는 객체를 보관했다가 다시 사용
+ **WaitGroup** : goroutine이 모두 끝날 때까지 기다릴 때 사용
+ **Atomic** : 더 이상 쪼갤 수 없는 연산(원자적 연산)을 뜻하며, 멀티 스레드(goroutine)/멀티 코어 환경에서 안전하게 값을 연산해줌


### Mutex ⭐
> 상호배제(mutual exclusion)라고도 함   
> **여러 곳에서 공유되는 데이터를 특정 부분에서만 사용하도록 보호하기 위해 사용** 함

+ func (m * Mutex) Lock() : 뮤텍스 잠금, 즉 **공유되는 데이터 보호 시작할 때 사용**
+ func (m * Mutex) Unlock() : 뮤텍스 잠금 해제, 즉 **공유되는 데이터 보호를 마칠 때 사용**
   + Lock, Unlock **짝이 안맞으면 데드락 발생**하므로, 잠그면 무조건 해제시켜줘야 함


### RWMutex
> 읽기/쓰기 뮤텍스, 읽기와 쓰기 동작을 나누어 잠금(Lock)을 걸 수 있음   
> **쓰기 작업을 할 때, 다른 곳에서 이전데이터를 읽지 못하도록 방지**하거나,   
> **읽기 작업을 할 때, 데이터가 변경되는 상황을 방지**할 때 사용됨   

+ func (rw * RWMutex) Lock() : 쓰기 뮤텍스 잠금
   + 공유되는 데이터에 **쓰기 작업을 할 수 있도록 보장**시켜줌
   + 즉, 다른 곳(gorutine)에서 해당하는 공유된 데이터에 대해 읽기/쓰기 작업 불가능
+ func (rw * RWMutex) Unlock() : 쓰기 뮤텍스 잠금 해제
+ func (rw * RWMutex) RLock() : 읽기 뮤텍스 잠금
   + 읽기 작업에 한해서 **공유되는 데이터가 변경되지 않음을 보장**시켜줌
+ func (rw * RWMutex) RUnlock() : 읽기 뮤텍스 잠금 해제
   + 마찬가지로 Lock, Unlock **짝이 안맞으면 데드락 발생**하므로 잠그면 무조건 해제시켜줘야 함


### Cond
> 대기중인 객체 하나만 깨우거나, 여러 개를 동시에 깨울 때 사용

+ * func NewCond(l sync.Locker) * sync.Cond : **조건 변수 생성**
   + **뮤텍스를 먼저 생성한 다음, sync.NewCond 함수로 조건 변수** 생성
+ func (c * Cond) Wait() : **goroutine 실행을 멈추고 대기**하는 함수
+ func (c * Cond) Signal() : **대기 중인 goroutine 하나만 깨우는** 함수
+ func (c * Cond) Broadcast() : 대기 중인 **모든 goroutine을 깨우는** 함수
   
   
### Once ⭐
> 특정 함수를 한번만 수행해야될 때 사용됨

+ func (o * Once) Do(f func()) : **입력받은 함수에 대해 한번만 수행**
   + sync.Once 할당 뒤, Do 함수를 통해 해당하는 함수 한번만 수행 가능
   + 복잡한 반목문에서 초기화할 때 유용한 함수임


### Pool
> 일종의 캐시와 같은 역할을 하며, 메모리할당과 해제 횟수를 감소시켜 성능을 높일 때 사용   
> 객체(메모리)를 사용한 후 보관해놓다가 다시 사용할 수 있게 해주는 기능   

+ func (p * Pool) Get() interface{} : pool에 **보관된 객체(데이터)를 가져오는 함수**
+ func (p * Pool) Put(x interface{}) : pool에 **객체(데이터) 보관하는 함수**   
   + 이를 통해 메모리를 효율적으로 관리할 수 있음
   + 수명 주기가 짧은 객체에 대해서는 pool이 적합하지는 않음  


### WaitGroup ⭐
> goroutine만 사용한다면 main이 종료될 때 대기없이 바로 같이 종료하게 됨   
> 이를 방지하기 위해, **지정한 개수의 goroutine이 완료될 때까지 기다릴 때 사용**함

+ func (wg * WaitGroup) Add(delta int) : 대기 그룹에 **고루틴 개수 추가**해주는 함수
+ func (wg * WaitGroup) Done() : **goroutine이 끝났음을 알려줄 때 사용**하는 함수
   + Add 함수에 설정한 값과 Done 함수가 호출되는 횟수는 동일해야 함
   + 동일하지 않을 시, 패닉 발생함
   + defer() 함수와 함께 사용해서 goroutine이 끝나기 직전에 wg.Done() 함수 호출하도록 사용 가능
+ func (wg * WaitGroup) Wait() : **모든 goroutine이 끝날 때까지 기다리도록** 하는 함수
  
