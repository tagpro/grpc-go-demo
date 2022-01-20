
.PHONY: generate
generate:
	protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/maths.proto

.PHONY: build/client
build/client:
	go build -o bin/client ./client && echo "generated ./bin/client"

.PHONY: build/server
build/server:
	go build -o bin/server ./server && echo "generated ./bin/server"

.PHONY: build
build: build/client build/server

.PHONY: run
run: build/server
	./bin/server

.PHONY: client
client: build/client
	./bin/client -num1 20 -num2 5
