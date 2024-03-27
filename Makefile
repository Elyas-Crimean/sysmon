build: sysmon client

sysmon:
	go build -v -o bin/sysmon ./cmd/sysmon
sysmon.exe:
	GOOS=windows go build -v -o bin/sysmon.exe ./cmd/sysmon
client:
	go build -v -o bin/client ./cmd/client
grpc:
	protoc -I./api:/usr/include/google --go_out=api --go-grpc_out=require_unimplemented_servers=false:api api/sysmon.proto