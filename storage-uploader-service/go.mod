module github.com/Condition17/fleet-services/storage-uploader-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/storage-uploader-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	cloud.google.com/go v0.71.0 // indirect
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201105193453-e2fb352f53cd
	github.com/Condition17/fleet-services/lib v0.0.0-20201105193453-e2fb352f53cd
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201105193453-e2fb352f53cd // indirect
	github.com/Condition17/fleet-services/user-service v0.0.0-20201105193453-e2fb352f53cd // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1 // indirect
	golang.org/x/tools v0.0.0-20201105173854-bc9fc8d8c4bc // indirect
	google.golang.org/genproto v0.0.0-20201105153401-9d023cd09d72 // indirect
	google.golang.org/protobuf v1.25.0
)
