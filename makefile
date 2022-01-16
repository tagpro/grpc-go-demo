
.PHONY: generate
generate:
	protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/maths.proto

.PHONY: build
build:
	go build -o bin/server server.go
	go build -o bin/client client.go

.PHONY: run
run: build
	./bin/server

.PHONY: client
client: build
	./bin/client
