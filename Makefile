#!make

gen:
	protoc --proto_path=proto proto/*.proto --go_out=./pb --go-grpc_out=./pb --grpc-gateway_out=./pb

clean:
	rm pb/*.go

server1:
	go run cmd/server/main.go -port 8080

rest1:
	go run cmd/server/main.go -port 8081 -type rest -endpoint 0.0.0.0:8080

server2:
	go run cmd/server/main.go -port 50051

server2-tls:
	go run cmd/server/main.go -port 50051 -tls

server3:
	go run cmd/server/main.go -port 50052

server3-tls:
	go run cmd/server/main.go -port 50052 -tls

client1:
	go run cmd/client/main.go -address 0.0.0.0:8080

client1-tls:
	go run cmd/client/main.go -address 0.0.0.0:8080 -tls

test:
	go test -cover -race ./...

run:
	go run main.go

cert:
	cd cert; ./gen.sh; cd ..

.PHONY: gen clean server client test cert
