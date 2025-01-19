module gobase_om

go 1.22.4

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.69.4 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1 // indirect
	google.golang.org/protobuf v1.36.3 // indirect
)

// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative ./cities.proto
