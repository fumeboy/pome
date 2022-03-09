pb:
	 protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative ./proto/A.proto
	 protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative ./proto/B.proto

initA:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./A/*.go
	docker build -t pome_service_a:v1 .
initB:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build ./B/*.go
	docker build -t pome_service_b:v1 .
init:
	make initA
	make initB

prepare:
	docker-compose up pome_etcd

test:
	docker-compose start pome_service_a
	docker-compose start pome_service_b
	go run ./client.go
	docker-compose stop pome_service_a
	docker-compose stop pome_service_b