module github.com/Condition17/fleet-services/lib

go 1.14

// this replaces help testing the implemented funcitonality
replace (
	github.com/Condition17/fleet-services/lib => ./.
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/Condition17/fleet-services/run-controller-service v0.0.0-20201207213217-0df36076c6cf
	github.com/Condition17/fleet-services/user-service v0.0.0-20201206222112-6d307be30861
	github.com/ghodss/yaml v1.0.0
	github.com/micro/go-micro/v2 v2.9.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
