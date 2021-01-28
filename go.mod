module pome

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
replace google.golang.org/grpc/v2 => github.com/fumeboy/grpc-go/v2 v2.0.3

replace go.etcd.io/etcd => github.com/coreos/etcd v0.0.0-20201125193152-8a03d2e9614b

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/marwan-at-work/mod v0.4.1 // indirect
	go.etcd.io/etcd v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/v2 v2.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
)
