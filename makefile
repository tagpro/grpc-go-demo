
.PHONY: generate
generate:
	protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/maths.proto

.PHONY: build
build:
	go build -o bin/server ./server && echo "generated ./bin/server"
	go build -o bin/client ./client && echo "generated ./bin/client"

.PHONY: run
run: build
	./bin/server

.PHONY: client
client: build
	./bin/client -num1 20 -num2 5
