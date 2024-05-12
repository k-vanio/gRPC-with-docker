install:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/ktr0731/evans@latest

protoc:
	protoc --go_out=./internal --go-grpc_out=require_unimplemented_servers=false:./internal ./proto/**/*.proto
	go mod tidy

run:
	go build -o ./bin/server ./cmd/main.go
	./bin/server