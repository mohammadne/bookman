# generate proto-buf files

## GO
- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest (protoc-gen-go)
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest (protoc-grpc-gen-go)

- protoc --proto_path=. --go_out=. service.proto
- protoc --proto_path=. --go-grpc_out=. service.proto
<!-- OR -->
- protoc --proto_path=. --go_out=. --go-grpc_out=. service.proto
