module github.com/Condition17/fleet-services/resource-manager-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/Condition17/fleet-services/resource-manager-service => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0
)

require (
	cloud.google.com/go v0.72.0 // indirect
	cloud.google.com/go/pubsub v1.8.3 // indirect
	github.com/Condition17/fleet-services/lib v0.0.0-20201115141257-e152a37ff788
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201115141257-e152a37ff788
	github.com/Condition17/fleet-services/test-run-service v0.0.0-20201115141257-e152a37ff788
	github.com/Condition17/fleet-services/user-service v0.0.0-20201115141257-e152a37ff788
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.3 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/googlepubsub/v2 v2.9.1
	github.com/ozonru/etcd/v3 v3.3.0-rc.0-grpc1.30.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	golang.org/x/sys v0.0.0-20201117170446-d9b008d0a637 // indirect
	golang.org/x/tools v0.0.0-20201117152513-9036a0f9af11 // indirect
	google.golang.org/api v0.35.0
	google.golang.org/genproto v0.0.0-20201117123952-62d171c70ae1 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
)
