genv2: 
	make clean
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

gen:
	make clean
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:. \

clean:
	rm -rf pb
	mkdir pb

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

run:
	go run cmd/server/main.go && go run cmd/client/main.go

test:
	go test ./serializer
	go test ./pb
	go test ./sample

coverage:
	go test -cover -race ./serializer
	go test -cover -race ./pb
	go test -cover -race ./sample