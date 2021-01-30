proto:
	 protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./proto/A.proto
	 protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./proto/B.proto

initA:
	CGO_ENABLED=0 go build ./A/*.go
	docker build -t pome_service_a:v1 .
initB:
	CGO_ENABLED=0 go build ./B/*.go
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