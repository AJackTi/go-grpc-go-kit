### Go + gRPC with Go kit

.
├── cmd
│   └── main.go          # main entrypoint file
├── endpoints
│   └── endpoints.go     # contains the endpoint definition
├── pb
│   ├── math.pb.go       # our gRPC generated code
│   └── math.proto       # our protobuf definitions
├── service
│   └── api.go           # contains the service's core business logic
├── transports
│   └── grpc.go          # contains the gRPC transport

```
protoc --go_out=. --go-grpc_out=. pb/math.proto
```

```
go run cmd/main.go
```
