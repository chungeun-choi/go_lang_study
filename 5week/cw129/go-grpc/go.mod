module go-grpc

go 1.20

// replace (
// 	go-grpc/data => ./data
// )

require (
	github.com/golang/glog v1.1.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.56.1
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230626202813-9b080da550b3 // indirect
)
