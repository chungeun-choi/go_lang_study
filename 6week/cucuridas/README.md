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
    | Use | 한 줄로 되어진 간단한 명령어 설명 |
    | Short | help 플래그를 통해 보여지는 간단한 명령어 설명 |
    | Long | help 플래그를 통해 보여지는 상세한 명령어 설명 |
    | Args | 사용자에 의해 입력되는 값 |
    | Run | 해당 커맨드를 통해 실행되는 비즈니스 로직 |
    
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