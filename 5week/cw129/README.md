# gRPC

## RPC
![image](https://github.com/CW129/go_lang_study/assets/104714337/24cff63c-47f9-4764-a5e9-bed2fb91523e)
    stub : 서버와 클라이언트는 서로 다른 메모리 주소 공간을 사용, 즉 함수 호출에 사용되는 매개변수를 변환이 필요, stub가 변환을 처리해줌
    skeleton :
    marshalling :
    unmarshalling : 
## gRPC 통신

![image](https://github.com/CW129/go_lang_study/assets/104714337/7efb3ea3-55fc-4439-aaca-135788a13bd7)

        
    
    gRPC = 통신을 위해 만들어진 RPC 기반 프레임워크
    

## gRPC vs RESTful
    RESTful
        HTTP 프로토콜 기반
        payload에 json 형식의 데이터를 담음 -> 평문을(텍스트) 전송함으로 속도가 느린편
        
    gRPC : 
        HTTP/2(HTTP 개정판) 프로토콜 기반
        payload에 protobuf를 담음 -> 이진 데이터 전송함으로 속도가 빠름
        서버-서버, 클라이언트-서버, 양방향 통신을 지원하지만 브라우저에선 gRPC 통신을 지원하지않음

## protobuf
