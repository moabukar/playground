# gRPC demo

- create proto file

- `protoc --go_out=.--go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    hello_world.proto`

- protoc-gen-go-grpc --go_out=.--go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    hello_world.proto


```bash

protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       /Users/mohameda/Documents/Learning/playground/gRPC/demo.proto
```

```bash
go run server.go

go run client.go
```
