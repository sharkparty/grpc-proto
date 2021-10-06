genv2: 
	make clean
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

gen:
	make clean
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:. \

clean:
	rm -rf pb
	mkdir pb
run:
	go run main.go