### 181p

- iota

    - db에서의 sequence처럼 소괄호 내에서 값이 0부터 1씩 증가하는 키워드이다.
    - const 상수 선언시에만 사용 가능하다. 변수 선언 시 아래와 같이 에러 발생

        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/ea80b64b-38f9-4389-aa4c-6265f53d7083)
        
        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/c27f5f00-8ca0-4693-a704-fa7c63e602d8)
    
    - 아래와 같이 첫번째만 연산구문을 작성하면 나머지 상수는 자동적으로 동일한 연산을 차례로 적용하여 상수를 만들 수 있다.
    
        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/a57cfef9-8705-4fed-82a7-c2c1b1ca7f40)
    
### 182p

- 타입 없는 상수

    - 상수는 타입을 선언하지 않아도 되며 이 경우 변수에 복사될 때 타입이 정해진다. 따라서 실제로는 실수인 상수라도 타입이 선언되지 않은 경우 정수와의 연산이 가능하다.

### 191p

- 쇼트서킷

    - 쇼트서킷이란 조건문에서 좌변에 대한 true, false 처리를 하는 현상으로 && 구문에서는 좌변이 false일 경우, || 구문에서는 true일 경우 우변의 조건을 건너뛴다.
    - 책에서는 이 구문에 대해 조심해야 하는 경우를 살폈지만, 선행적으로 필터링 해야 하는 조건을 좌변에 넣으면 효율적인 코드를 작성할 수 있지 않을까 생각된다. 좌변에 오는 식은 중첩 if 문에서의 외부 조건절, 우측은 내부 조건절이라고 볼 수 있다.
    
        ```go
        if 조건 A && 조건 B {}

        if 조건 A {
            if 조건 B {

            }
        }
        ```

### 194p

- if 초기문 : 조건문에서 변수 초기화 시 사용

    ![image](https://github.com/haseungyeon/fastapi/assets/59682268/cdb0cca1-1e47-40ae-9078-aeac749575c1)

    - 초기문에서 사용된 변수는 if 문 블록범위에서만 유용하다.
    - 초기문을 사용하면 필요한 조건에서만 변수를 사용하므로 공간이 절약되지 않을까 싶다.

### 208p ~

- go 언어에서의 switch 문 특징

    - 한 번에 여러 값을 비교 할 수도 있다.
    
        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/4cbcfcbc-208d-4022-85d2-474e159f6aa3)
    
    - if 문과 마찬가지로 초기문을 사용할 수 있다. (동일하게 switch 문 내부에서만 유효하다.)
    
        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/8aadf807-755f-4d84-977c-e0c315476a57)
    
    - 다른 언어의 switch와 다르게 case 문 내부에 조건문을 넣을 수 있다.
    
        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/5fc42eb7-57af-450c-90b1-cb236eaee7dc)

    - fallthrough 키워드 : 일치 또는 만족하는 조건에 해당하는 case 문 바로 다음의 case 문까지 실행이 된다.

        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/292b206c-56d6-4f4f-a7a0-011262798014)

        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/27f5bd70-74f7-450c-a673-e47468a3d5c0)

    **산탄총 수술 문제**

        => 보통 switch 문은 case 문 비교값에 상수 열거값을 사용하는데 열거값에 관련된 새로운 상수값이 추가되면 그와 관련된 case도 계속해서 추가해줘야 하는 작업이 발생한다. 따라서 case를 조건식으로 표현하여 가짓수를 줄여주는 것이 좋다.

### 224p

- for 문의 여러 형태

    ![image](https://github.com/haseungyeon/fastapi/assets/59682268/ac474db7-4223-4a98-9a48-5cb2e5cdf657)

    ![image](https://github.com/haseungyeon/fastapi/assets/59682268/389869a9-40de-4386-b3cc-b4ef8f415827)

    => for 문 하나로 while문, 무한루프를 구현할 수 있음.

### 230p

- 레이블
    - 레이블 키워드를 지정하여 for 문 내에 어떤 부분에서 전체 루프를 종료할 지 지정할 수 있다.
    
        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/4c1cda97-0810-4cdc-9fcf-c5407a3a791a)
    
    - 레이블을 많이 쓰면 가독성이 떨어져 함수로 break를 구현
    - 함수도 활용하고 기능 단위로 쪼개진 함수 내부에 레이블을 사용하면 어떨까?

### 238p

- 배열
    - 사용법
    
        ```go
        var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

        var s = [5]int{1:10, 3:30}

        x := [...]int{10, 20, 30} // 개수 생략 시 선언 대입 시의 요소 수로 지정됨
        ```
### 247p

- 다중배열
    - 사용법
    
        ```go
        var tt [2][5]int = [2][5]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5, 6}}
        ```

        - 일반 배열과 동일함.
        - 초기화 시 요소의 중괄호가 가장 바깥 쪽의 중괄호와 동일한 줄에 위치하지 않을 경우 `,`를 찍어주어야 한다.
        
            ![image](https://github.com/haseungyeon/fastapi/assets/59682268/ce0f6ee0-83f0-484b-a4de-3650a4fb06d6)

### 256p

- 구조체
    - 여러 타입의 변수들을 하나의 변수로 묶어서 사용
    - 사용법
    
        ```go
        type person struct {
            var Name string
            var Age int
            var Height float64
            var Weight float64 
        }
        ```
    - 타입명의 첫 글자가 대문자일 경우 외부로 공개되는 타입.

### 257p

- 구조체의 초기화
    - 배열과 마찬가지로 일부 또는 전체 초기화가 가능.

### 258p

- 구조체를 포함하는 구조체
    - 클래스 상속과 유사하게 이미 정의된 구조체의 필드를 재사용 하는 구조.

        ```go
        type User struct {
            Name string
            ID   string
            Age  int
        }

        type VIPUser struct {
            UserInfo User
            VIPLever int
            Price    int
        }
        ```
    - 필드명을 생략한다면 "."을 이중으로 찍지 않고 구조체 내부의 구조체의 필드에 접근할 수 있다.
    - 만약 구조체 내부에 여러 구조체를 사용할 경우에는 내부에 사용되는 구조체 타입을 명시해주어야 한다.

        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/a4216787-0055-47c5-9d4a-4de32d5c70a2)

        ![image](https://github.com/haseungyeon/fastapi/assets/59682268/0f028cc9-5aef-467b-a7e0-f63b82475227)

    ### 266p

    - 메모리 정렬과 패딩
        - 컴퓨터는 변수를 메모리에 저장할 때 주소 탐색 시간을 줄이기 위해 변수 크기의 배수에 해당하는 주소값에 저장한다.
        
            ex) var number int = 1
            => 8바이트인 int의 배수에 해당하는 104번지에 변수 저장

        - 구조체 변수를 저장할 경우 먼저 선언된 변수의 용량이 후순위 선언된 변수보다 용량이 클 경우 용량이 큰 변수 크기에 맞춰 적은 용량의 변수에 공간을 추가시키는 동작을 **패딩** 이라고 한다.
        
        패딩 현상을 방지하고 메모리 공간을 효율적으로 사용하기 위해서는 구조체 필드값 선언 시 공간이 작은 변수부터 지정해주어야 한다.

### 277p

- 포인터
    - 변수의 메모리 주소를 값으로 가지는 타입
    - 포인터를 사용하는 이유
        - 구조체나 배열 같은 메모리 공간을 많이 차지하는 변수들을 입력받거나 복사할 경우 전체 값을 복사하기에는 성능과 공간낭비가 심해지기 때문.
        - go 에서는 인수가 함수 내부에 입력될 경우 매개변수에 값이 복사가 되므로 직접적인 외부인수로 입력된 변수값을 변경할 수 없는데 포인터를 사용하면 가능하다.
    - 포인터 변수 사용법
    
        ```go
        var a int
        var p *int
        p = &a
        ```
    - 구조체 포인터 변수 사용법
    
        ```go
        // 변수 선언 및 포인터 변수 선언
        var data Data
        var p *Data = &data

        // 즉시 선언, 이 경우 구조체 변수 명은 미정
        var p *Data = &Data{}

        // new 키워드 사용, 이 경우 초기화는 불가능
        var p = new(Data)
        ```
### 286p

- 인스턴스
    - 인스턴스란 메모리에 할당된 변수의 실체이다.
    - 여러 개의 포인터 변수를 생성하여 동일한 변수를 가리킨다 하더라도 인스턴스가 새로 생성되는 것은 아니다.
    - 인스턴스는 go 언어의 가비지 콜렉터가 자동으로 소멸시키며, 함수 내부에 선언된 변수가 블록 스코프를 벗어날 경우 삭제하는 식으로 동작한다. 단, 함수 외부로 공개되는 인스턴스는(대문자로 선언됨) 외부 패키지에서 사용되므로 소멸되지 않는다.

### 290p

- 스택 메모리
    - go 언어에서 스택 메모리는 공간이 동적으로 결정되며 함수 호출 시 사용된다. 동적으로 할당되기 때문에 타 언어와 달리 재귀호출 등으로 인한 메모리가 고갈 될 걱정이 없다.
- 힙 메모리
    - 마찬가지로 동적으로 크기가 결정되며 함수 외부에서 사용되는 데이터와 외부 공개 패키지 레벨의 데이터가 저장된다.

- go 에서는 탈출 분석을 통해 해당 데이터의 외부사용여부를 따져 힙에 할당할지 스택에 할당할지와 인스턴스의 소멸여부를 결정된다.

### 296p

- 문자열의 출력과 크기, 길이
    - 문자열 출력 시 슬라이스 Rune 타입으로 캐스팅 후 반환하는 방법도 있으나 타입캐스팅 시 발생되는 새로운 메모리 공간에 대한 낭비로 인해 아래와 같은 방법을 사용하는 것이 좋다.
    
        ```go
        str := "Hello는 한국말로 안녕"
        for _, v := range str {
            fmt.Print("%c", v)
        }
        ```
    - 문자열 크기
    
        ```go
        str := "Hello는 한국말로 안녕"
        len(str)
        ```
    - 문자열 길이
    
        ```go
        str := "Hello는 한국말로 안녕"
        len([]rune(str))
        ```
### 313p

- 문자열 합치기
    - 문자열은 불변형이므로 새로운 문자열을 합치거나 변형할 경우 새로운 메모리 공간을 할당하게 된다. 이는 메모리 공간의 낭비로 이어지므로 strings 패키지의 builder.WriteRune 사용이 권장된다.
    - strings.builder는 슬라이스를 사용하므로 새로운 메모리 공간을 생성하지 않고 기존 공간을 그대로 사용한다.
    
        ```go
        package main

        import (
            "strings"
            "fmt"
        )

        func ToUpper(str string) string {
            var builder strings.Builder
            for _, c := range str {
                if c >= 'a' && c <= 'z' {
                    builder.WriteRune('A' + (c-'a'))
                } else {
                    builder.WriteRune(c)
                }
            }
            return builder.String()
        }

        func main() {
            va
        }
        ```

### 322p

- 패키지
    - 기본 패키지
    
        ```go
        import (
            "fmt"
            "math/rand"
            "text/template"
            htemplate "html/template" // 패키지 이름이 겹칠 경우 별칭 사용
        )
        ```
    - 모듈 패키지
        - 직접 만든 모듈 내부에 생성된 패키지
        - go mod init [module_name]/[package_name] 으로 생성한 모듈 경로를 그대로 입력
        
            ```go.mod
            module ex13.2/ex13.2

            go 1.20
            ```
            ```go
            import "ex13.2/ex13.2"
            ```
    - 외부 패키지
        - 로컬 폴더에 저장된 패키지가 아닌 외부로부터 불러오는 패키지
        
            ```go
            "database/sql"
            _ "github.com/mattn/go-sqlite3" // 부가 효과
            ```
        - 외부 패키지에 노출시킬 변수나 함수들은 첫글자를 대문자로 표기한다.
        - 외부 패키지 사용 시 아래 명령어를 실행하여 mod 파일을 갱신시키고 패키지를 다운로드 한다.

            ```bash
            go mod tidy
            ```
