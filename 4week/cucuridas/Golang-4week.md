# Go lang 스터디 - 4주차

# 고루틴과 동시성 프로그래밍

## Context switching 비용 - p.508

CPU 코어가 여러 스레드를 전환하면서 수행하면 더 많은 비용이 발생 → 해당 비용을 컨텍스트 스위칭이라고 함

이 비용은 저장하는 요소에 따라 크게 증가하는 데 증가하는 요소는 **명령어 포인터, 스택 메모리 등을 저장 이를 컨텍스트라 칭함**

<aside>
💡 **스레드의 적정개수
코어 개수의 두 배 이상 스레드를 만들면 스위칭 비용이 많이 발생
→ GO 언어에서는 코어 마다 OS 스레드를 하나만 할당해서 사용**

</aside>

## 고루틴 사용 방법 - p.508

```go
go 함수 _호출
```

<aside>
💡 **고루틴은 메인함수에서 시작해서 메인함수에서 종료
메인 함수가 종료되면 아무리 많은 고루틴이 생성되어 있더라도 모두 즉시 종료되고 프로그램이 종료**

</aside>

## WaitGroup 객체

고루틴이 완료될 때 까지 메인함수를 실행 시키기위한 객체

```go
var wg sync.WaitGroup

wg.Add(3) // 완료해야하는 작업 개수 설정
wg.Done() // 작업이 완료될 때 마다 작업 개수를 하나씩 줄여주는 함수
wg.Wait() // 전체 작업이 모두 완료될 때까지 대기
```

## 고루틴의 동작 방법 - p. 513

고루틴은 명령을 수행하는 단일 흐름으로 OS 스레드를 사용하는 경량 스레드

<aside>
💡 OS 스레드란?

OS 스레드는 운영체제에 의해 스케줄되고 관리되며, CPU 코어의 쓰레드 스케줄러에 의해 실행 순서가 결정됩니다. 스레드는 프로세스의 자원을 공유하기 때문에 스레드 간의 통신과 동기화를 위해 별도의 기법이 필요할 수 있습니다.

OS 스레드는 멀티스레딩을 지원하고, 동시에 여러 스레드가 실행될 수 있어 CPU 활용도를 높일 수 있습니다. 또한, I/O 작업이나 병렬 계산 등 다양한 작업을 동시에 처리할 수 있어 응답성과 성능 개선에 도움을 줍니다.

</aside>

**2개의 코어를 가진 컴퓨터에서 고루틴이 어떻게 동작하는지**

- 고루틴이 하나 일 때
    
    코어 1 ↔ OS 스레드 1 ↔ 고루틴1
    
- 고루틴이 두개 일 때
    
    코어 1 ↔ OS 스레드 1 ↔ 고루틴1
    
    코어 2 ↔ OS 스레드 2 ↔ 고루틴3
    
- 고루틴이 3개 일때 (**끝나는 고루틴 작업의 코어에서 남은 고루틴 작업이 실행됨**)
    
    코어 1 ↔ OS 스레드 1 ↔ 고루틴1
    
    코어 2 ↔ OS 스레드 2 ↔ 고루틴2
    
    대기열: 고루틴3
    

## 동시성 프로그래밍 주의점 - p.516

**동이한 메모리 자원에 고루틴이 접근할 때 발생**

### 뮤텍스를 사용한 문제 해결 - p.518

한 고루틴에서 사용중인 값을 다른 고루틴이 건들지 못하게 하는 것

```go
var mutex sync.Mutex

func DepositAndWithdraw(account *Account){
	mutext.Lock()
	...
	...
	defer mutex.Unlock()
}
```

- 단점
    - 동시성 프로그래밍을 통해 얻는 성능 향상을 얻을 수 없음(하나의 자원에 하나의 고루틴만 사용이 가능기에)
    - 데드락 발생 가능: 한 쪽이 자원을 점유함에 따라 프로그램 실행을 위한 최소 조건의 자원을 만족하지 못하는 것
    
- 해결 방안: 뮤텍스가 꼬이지 않도록 좁은 범위에서 데드락이 걸리지 않는 지 확인해서 사용하기
    
    (google cloud profiler와 같은 프로파일링 툴)
    
    [[GCP]Stackdriver Profiler 살펴보기](https://medium.com/@jwlee98/gcp-stackdriver-profiler-살펴보기-dfdab3d70c4b)
    

## 또 다른 자원 관리 기법

- 영역을 나누는 방법
    
    ```go
    type job interface {
    	Do()
    }
    
    type SquareJob struct {
    	index int
    }
    
    func (j *SquareJob) Do() {
    	fmt.Printf("%d 작업 시작\n", j.index)
    	time.Sleep(1 * time.Second)
    	fmt.Printf("%d 작업 완료 - 결과: %d/\n",j.index, j.index*j.index)
    }
    
    func main() {
    	var jobList [10]job
    	// 작업 구역을 리스트에 할당
    	for i := 0l i <10; i++{
    		jobList[i] = &SquareJob{i}
    	}
    
    	var wg sync.WaitGroup
    	wg.Add(10)
    
    	for i := 0; i < 10; i ++ {
    		// 앞서 작업구역할당내용을 불러와 작업 시작
    		job := jobList[i]
    		go func() {
    			job.Do()
    			wg.Done()
    		}{}
    	}
    	wg.Wait()
    }
    ```
    

# 채널과 컨텍스트 - p.530

## 채널이란?

고루틴 끼리 메세지를 전달 할 수 있는 메세지 큐

## 채널 생성 방법 및 데이터 넣고 빼기

```go
var messages cha string = make(chan string)
//.          채널 타입         채널 키워드 타입

messages <- "This is a message"

var msg string = <- messages
```

## 채널 크기 설정 - p.534

일반적으로 채널을 생성하면 크기가 0인 채널을 생성하기에 데이터를 담아 둘 공간이 존재 하지 않음 
→ 버퍼를 가진 채널을 생성하여 보관할 수 있는 영역을 구현

```go
var chan string messges = make(chan string, 2)
```

## 채널에서 데이터 대기 - p.534

아래와 같은 구문을 통해 채널에 데이터가 쌓이기를 계속해서 대기하게 됨

이에 따라 **데드락 발생**

```go
for n := range ch {
	...
}
```

close 함수를 통해 작업이 끝났다면 채널을 반납하도록 정의 

채널을 제때 닫아주지 않아 고루틴에서 데이터를 기다리며 무한히 대기하는 경우를 **좀비 루틴 또는 고루틴 릭이라고 부름**

```go
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)	
	
	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i <10 ; i++{
		ch <- i* 2
	}

	close(ch)
	wg.Wait()

}
```

## Select 문을 통한 대기 유무와 여러 채널을 동시에 확인 - p.540

```go
func square(wg *sync.WaitGroup, ch chna int, quit chan bool){
	for {
		select {
			case n := <-ch:
				fmt.Printf("Square: %d\n",n*n)
				time.Sleep(time.Second)
			case <-queit:
				wg.Done()
		}
	}
}
```

## 컨텍스트 사용하기

context 패키지에서 제공하는 기능으로 작업을 지시할 때 **작업 가능 시간, 작업 취소 등의 조건을 지시하는 작업 명세서 역할**

```go
func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go printEverySecont(ctx)
	time.Sleep(5*time.Second)
	//5초 뒤에 cancel 함수를 호출해 ctx에 채널에 종료 시그널을 보냄
	cancel()

	wg.Wati()

}

func PrintEverySecond(ctx context.Context){
	tick := time.Tick(time.Second)
	for {
		select {
		// cancel 명령어를 통해 생성된 종료 시그널 채널의 값을 확인해 고루틴을 종료
		case <-ctx.Done():
			wg.Done()
			return
		case <- tick:
			fmt.Println("TIck")
		}
	}
}
```

- 작업 시간을 설정한 컨텍스트
    
    ```go
    ctx, cancel  := context.WithTimeout(context.background(), 3*time.Second)
    ```
    
- 특정 값을 설정한 컨텍스트
    
    ```go
    ctx := context.WithValue(context.Background(), "number",9)
    ```
    

# 객체 지향 설계 원칙 SOLID

## 단일 책임 원칙 - p.578

정의: 모든 객체는 책임을 하나만 져야한다

이점: 코드 재사용성을 높여줌

## 개방-페쇄 원칙 - p.581

정의: 확장에는 열려 있고 변경에는 닫혀 있다

이점: 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 됩니다

## 리스코프 치환 원칙 -p.583

정의: ‘q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라 하자. 그렇다면 S가 T의 하위 타입이라면 q(y)는 타입 S의 객체 y에 대해 증명 할 수 있어야한다 ’

이점: 예상치 못한 작동을 예방할 수 있다

## 인터페이스 분리 원칙 -p.586

정의: 클라이언트는 자ㅇ신이 이용하지 않는 메서드에 의존하지 않아야한다

이점: 인터페이스를 분리하면 불필요한 메서들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용가능

## 의존 관계 역전 원칙 - P.587

정의: 상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전 시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다

원칙

- 상위 모듈은 하위 모듈에 의해서는 안된다. 둘다 추상 모듈에 의존해야 함
- 추상 모듈은 구체화 된 모듈에 의존해서는 안된다 구체화 된 모듈은 추상모듈에 의존해야 한다

이점 

- 구체화된 모듈이 아닌 추상 모듈에 의존함으로써 확장성이 증가합니다
- 상호 결합도가 낮아져서 다른 프로그램으로 이식성 증가

# 테스트

## 작성 규약 - p.596

- 파일 명이 ‘_test.go’ 로 끝나야 합니다
- testing 패키지를 임포트 해야합니다
- 테스트 코드는 func TestXxxx(t *testing.T) 형태이어야합니다

## 테스트를 돕는 외부 패키지

“stretchr/testify” 테스트하고 테스트실패를 알릴 수 있는 다양한 함수 제공

```go
func TestSquare1(t * testing.T){
	assert := assert.Nes(t)
	assert.Equal(81, squere(9), "Squere(9) should be 81")
}
```

- Equal(): 두 값을 비교하여 다를 경우 테스트 실패
- Greater(): 두 값을 비교하여 대소 비교
- Len(): object 항목의 length가 입력된 값(int)가 아닐경우
- NotNilIf(): object가 nil 이면 테스트 실패
- NotEqualIf(): 두 값이 같을 경우 실패

## TDD(Test Driven Development) 테스트 주도 개발

테스트 작성 시기를 코드 작성 이전으로 옮겨 개발

## 벤치 마크

벤치 마크 코드는 테스트 작성 규약 1,2번을 따르며 3번의 항목의 경우 func BenchmarkXxx(b *testing.B)을 따름