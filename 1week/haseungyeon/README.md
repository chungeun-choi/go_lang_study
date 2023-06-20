### 52p

- go mod init golang\hello 실행 시
현재 디렉토리가 C:\Users\haseu\Desktop\Practice\golang\hello 명령어 중 path 입력방식이 상대경로도, 절대경로도 아닌데 어떻게 인식을 하는건지??

    ![image](https://github.com/haseungyeon/study/assets/59682268/31197275-1e3b-4468-9519-cc87036828b9)

    ![image](https://github.com/haseungyeon/study/assets/59682268/35e532a7-752c-458a-83ee-1ed4e7e85935)

-> 실제 path와 다르게 지어도 생성은 됨

**Go 파일 실행 순서(윈도우)**

1. go mod init [project_name]/[package_name]
2. go build
3. .\[exe_file]

### 76p

- 동적 컴파일 언어는 프로그램 실행 시점에 컴파일이 된다고 한다.
정적 컴파일 언어의 단점이 실행환경이 다를 경우, 그에 맞는 실행 파일을 빌드해야 하는 점인데
golang은 정적 컴파일 언어임에도 환경변수 설정을 통해 유연한 빌드가 가능하여 이런 단점을 극복하면서도 정적 컴파일 언어의 장점인 빠른 속도를 갖출 수 있다고 한다.
그렇다면 그럼에도 불구하고 동적 컴파일 언어를 써야 하는 이유가 뭘까?

=> 빌드를 매번 하지 않아도 되기 때문에 개발 생산성이 향상되는 효과가 있다.

### 89p

![image](https://github.com/haseungyeon/study/assets/59682268/7f4daf12-4f2f-466b-ba6d-83e2a1e4d759)

- 고랭도 C언어처럼 main 함수가 필수적으로 존재해야 한다. 따라서 main을 사용하기 위해 package main을 작성해야 한다.
=> main()함수가 없는 패키지는 패키지 이름으로 main을 쓸 수 없다. 바꿔 말하면, main()함수를 써야지만 package main으로 .go 파일을 main 패키지로 선언할 수 있다?

### 98p

- 고랭에서는 카멜 케이스 표기법을 사용한다. ex) goLanguage

### 99p

- 고랭의 타입 : 숫자, 불리언, 문자열, 배열, 슬라이스, 구조체, 포인터, 함수, 인터페이스, 맵, 채널등의 타입이 있다.

- 타입을 선언해야 하는 이유? => 내 생각에 고랭이나 C언어와 같은 강타입(타입 선언 강제성이 강함)언어들은 약타입 언어에 비해 메모리(공간)를 낭비할 수 밖에 없다. 다만, 타입선언을 통해 메모리 공간이 정해지면 시작과 끝 주소값 또한 정해지므로 검색시간이 짧아져 시간적인 측면에서 성능이 좋아진다.

### 102p

- 타입별 기본값 : 무의 상태가 기본값 이라고 생각하면 외우기 쉽다.
 
    - int : 0
    - float64 : 0.0
    - bool : false
    - string : ""
    - 그 외 : nil

### 103p

- 선언 대입문 := ex) c := 365

=> 변수 선언, 타입 키워드를 생략하고 우측의 타입을 그대로 적용. 간단해서 편하겠지만, 가독성은 조금 떨어질수도 있다. 타입을 유연하게 받아들일 경우 `:=`를 사용하고 그 외 목적이 확실한 경우는 타입을 확실히 선언해주는게 좋을 듯?

- 고랭 = 융통성 없는 강타입 언어

    ![image](https://github.com/haseungyeon/study/assets/59682268/2e3cee83-221d-4ee5-bb7a-91836fdf39c8)

- question : 내 컴퓨터는 64bit 컴퓨터이고 따라서 int로 선언 시 int64와 같다라고 100p 표에 언급되어 있다. 위 이미지에서는 타입이 달라서 연산 불가라고 나오는데 같은 공간을 차지한다라는 뜻일 뿐 타입이 달라 컴파일 시 에러가 날까?

![image](https://github.com/haseungyeon/study/assets/59682268/d695e9e1-4afe-4820-8dcf-aaacf7605396)

=> 안됨. 그냥 64bit 컴퓨터일 땐 굳이 복잡하게 int64 쓰지 말고 int 쓰자.

### 111p

- 기본적으로 정수와 실수형을 선언 대입문 `:=`을 사용하여 선언할 경우 우선적으로

    - 정수 : int
    - 실수 : float64

    타입이 지정된다.

### 112p

- f1의 경우 float32는 7자리까지 표현 가능하니까 42707.40까지만 표현되어야 하는게 아닌가??

    ![image](https://github.com/haseungyeon/study/assets/59682268/edafbe9e-4d3d-4c75-932f-1d12c9505ded)


### 128p

```go
stdin := bufio.NewReader(os.Stdin)
stdin.ReadString('\n')
```
- 위 구문이 없으면 에러가 발생했을 경우 a, b값을 다시 입력받지 않고 첫 번째 입력받은 a, b값이 그대로 두 번째 출력구문에서도 출력됨.
