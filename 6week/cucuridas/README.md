# Go lang 스터디 - 6주차

# Cobra 사용하기

## Cobra? 무엇인가?

코브라는 바둑의 CLI 프레임워크이다. 강력한 최신 CLI 응용프로그램을 생성하기 위한 라이브러리와 Cobra 기반 응용프로그램 및 명령 파일을 신속하게 생성하기 위한 도구가 포함되어 있으며

- 장점
    1. 간편한 명령어 구조: Cobra는 명령어와 하위 명령어들을 쉽게 정의할 수 있는 구조를 제공합니다. Command, Subcommand, Flag 등을 쉽게 추가하고 구성할 수 있어서 명령어의 계층 구조를 명확하게 표현할 수 있습니다.
    2. 자동 완성 기능: Cobra는 명령어와 플래그의 자동 완성 기능을 내장하고 있어서 사용자가 탭을 누를 때 가능한 옵션들을 보여주어 CLI 사용의 편의성을 높여줍니다.
    3. 다양한 Flag 유형: Cobra는 문자열, 불리언, 정수 등 다양한 유형의 플래그를 지원하며, 사용자가 입력한 값들을 쉽게 추출하고 사용할 수 있습니다.
    4. Help 커맨드 지원: Cobra는 기본적으로 **`-help`** 플래그를 지원하며, 자체적인 Help 커맨드를 통해 사용자에게 명령어와 플래그의 사용법과 설명을 자세하게 보여줄 수 있습니다.
    5. Plugin 지원: Cobra는 외부 플러그인을 지원하여 기능을 확장할 수 있습니다. 이를 통해 개발자는 필요한 기능을 쉽게 추가하고 확장할 수 있습니다.
    6. 풍부한 기능과 커뮤니티: Cobra는 많은 사용자들에게 사용되고 있으며, 풍부한 기능과 다양한 커뮤니티 지원을 받고 있습니다. 따라서 문제가 발생했을 때 도움을 받기 쉽고, 많은 예제와 자료를 찾아볼 수 있습니다.
    7. 명령어의 실행 제어: Cobra는 각 명령어의 실행을 제어하는 PreRun, Run, PostRun 등의 Hook 메서드를 제공하여 명령어 실행 전후에 원하는 동작을 추가할 수 있습니다.

## **Getting Started ‘Cobra’**

- **Go version 관리를 위한 GVM**

[Mac에서 GVM으로 Golang 버전 관리하기](https://gurumee92.tistory.com/308)

- 파일 구조
    
    ```go
    ▾ appName/
        ▾ cmd/
            add.go
            your.go
            commands.go
            here.go
          main.go
    ```
    

- **command 생성 구조**
    
    1) `cobra.Command` 구조체를 통해 명령어 객체 생성
    
    ```go
    var cmdPrint = &cobra.Command{
        Use:   "print [string to print]",
        Short: "Print anything to the screen",
        Long: `print is for printing anything back to the screen.
    For many years people have printed back to the screen.`,
        Args: cobra.MinimumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
          fmt.Println("Print: " + strings.Join(args, " "))
        },
      }
    ```
    
    구조체 필드 정의(필요항목만 정리)
    
    | 필드 명 | 용도 |
    | --- | --- |
    | Use | 서브 커맨드 이름 |
    | Short | help 플래그를 통해 보여지는 간단한 명령어 설명 |
    | Long | help 플래그를 통해 보여지는 상세한 명령어 설명 |
    | Args | 사용자에 의해 입력되는 값 |
    | Run | 해당 커맨드를 통해 실행되는 비즈니스 로직 |
    
    <aside>
    💡 **Args validatation check를 위한 함수 구현**
    
    ```go
    Args: func(cmd *cobra.Command, args []string) error {
    		if len(args) != 1 {
    			return errors.New("enter the URL")
    		}
    
    		_, err := url.ParseRequestURI(args[0])
    		if err != nil {
    			return errors.New("invalid URL")
    		}
    
    		return nil
    	},
    ```
    
    </aside>
    
    2) 정의 후 rootCmd정의
    
    ```go
    var rootCmd = &cobra.Command{Use: "app"}
    ```
    

3) `AddCommand()` 함수를 통해 명령어 구조 생성

```go
rootCmd.AddCommand(cmdEcho)
cmdEcho.AddCommand(cmdTimes)
```

4) `execute()` 명령어를 통해 명령어 실행

**해당 method는 root command에서 호출되어 사용됩니다**

```go
rootCmd.Execute()
```

# kubectl - apply 코드 분석

### 구조체

| 이름 | 설명 | 모듈 이름 |
| --- | --- | --- |
| ApplyFlags | CLI 플래그가 수집하느 ㄴ정보를 반영하여 옵션으로 변환, 이 옵션은 명령어 런타임 요구 사항을 반영하고 구조체의 변환을 줄이기 위해 사용 | apply.go |
| ApplyOptions | apply 명령어를 위한 플래그나 다른 설정 입력 값을 정의한 구조체 | apply.go |
| ApplyDeleteOptions | apply 명령시 삭제해야할 요소를 설정하는 옵션 구조체 | applyset_pruner.go |
| PruneObject | prune 객체는 prune 중 일부분 삭제해야하는 api 객체 입니다 | applyset_pruner.go |
| ApplySet | ApplySet은 ApplySet Apply/Prune에 대한 정보를 추적 | applyset.go |
| ApplySetParentRef | ApplySetParentRef는 적용 세트를 추적하는 데 사용되는 상위 개체의 개체 및 유형 메타를 저장합니다. | applyset.go |
| ApplySetTooling | Apply가 호출된 이름과 version 정보를 담아두는 구조체 | applyset.go |
| Patcher | Patcher는 OpenAPI 개체를 패치하기 위한 옵션을 정의합니다. | patcher.go |
| Pruner | 메니페스트로 정의되어지지 않은 자원을 삭제하기 위한 정보가 담겨진 구조체 | prune.go |
| ViewLastAppliedOptions | ‘apply view-last-applied’ 명령어를 통해 출력되는 구조체 | apply_view_last_applied.go |
| SetLastAppliedOptions | ‘apply set-last-applied’ 명령어를 통해 출력되는 구조체 | apply_set_last_applied.go |
| PatchBuffer | 적용할 변경 사항을 캐시에 담아두는 구조체 | apply_set_last_applied.go |