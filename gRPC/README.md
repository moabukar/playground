# gRPC demo

- create proto file

```bash
❯ protoc --proto_path=/Users/mohameda/Documents/Learning/playground/gRPC/ \
       --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       <change_this path to the path where your proto file exists>

```

```bash
go run server.go

go run client.go
```
