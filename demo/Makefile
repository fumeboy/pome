init:
	 protoc --go_out=plugins=grpc:./A/guestbook/ ./A/guestbook/guestbook.proto
	 protoc --go_out=plugins=grpc:./B/guestbook/ ./B/guestbook/guestbook.proto

run1:
	go run ./B/*.go
run2:
	go run ./P/*.go
run3:
	go run ./A/*.go