# Create gRPC files from .proto 

### Generate files
*Command:*  
```
protoc --proto_path=./proto --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false proto/*.proto
```

### Testing
[Instruction](https://gist.github.com/p12s/2298bb21e1a53d9a1cbbf0cd54b90404)
Command:  
```sh
evans -p PORT -r
```

### Problem solving
1. Wrong or undefined files saving path:  
  
*ERROR:*  
```
WARNING: Malformed 'go_package' option in "proto/logger.proto", please specify:
option go_package = "./;__";
```
  
**Solution option:**  
```
option go_package = "./pb";
```

2. Not found program ([source](https://grpc.io/docs/languages/go/quickstart/)):   
  
*ERROR:*  
```
protoc-gen-go-grpc: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.
```
  
**Solution option:**  
- Install the protocol compiler plugins for Go using the following commands:  
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```
- Update your PATH:  
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
