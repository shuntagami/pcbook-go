#!make

gen:
	protoc --proto_path=proto proto/*.proto --go_out=./pb --go-grpc_out=./pb

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8080

client1:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...

run:
	go run main.go
