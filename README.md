## Notes

### Folder structure

    - grpc-go-demo
        - README.md
        - server.go
        - client.go
        - proto
            - maths.proto
            - maths.pb.go
            - maths_grpc.pb.go


### Dependencies

    - [go](https://formulae.brew.sh/formula/go)
    - [protoc](https://formulae.brew.sh/formula/protobuf)
    - [grpc and go proto plugins](https://grpc.io/docs/languages/go/quickstart/#prerequisites) (Required to compile the proto files into go code)

### Installation

    - `make generate` - generate go code from proto files
    - `make run` - run the server
    - `make client` - run the client


## Development

### Setup steps

1. `go mod init github.com/tagpro/grpc-go-demo` to initialise this repo as a go project
2. `go mod tidy` to update `go.mod` file using inbuilt go tools.

### Generate code from proto files

    - `make generate` does 2 things.
      - It generates go code from the proto files and it generates a `maths.pb.go` file that contains the generated code.
      - It generate grpc code from the proto files and it generates a `maths_grpc.pb.go` file that contains the generated code.
    - `make run` builds `server.go` and runs the server.
    - `make client` builds `client.go` and runs the client.

### Calling the server

    - `grpc_cli ls localhost:8000 -l` lists the services available on the server.
    - `grpc_cli call localhost:8000 Add "num1: 1, num2: 2"` calls the Add service on the server.