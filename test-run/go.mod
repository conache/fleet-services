module test-run

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/Condition17/fleet-services/common v0.0.0-20200919205711-d2a60ae78150
	github.com/Condition17/fleet-services/file-service v0.0.0-20200902192511-2fa60a36c63c
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.8.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	google.golang.org/protobuf v1.25.0
)
