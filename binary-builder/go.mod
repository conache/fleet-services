module github.com/Condition17/fleet-services/binary-builder

go 1.14

replace (
	github.com/Condition17/fleet-services/binary-builder => ./.
)

require (
	cloud.google.com/go v0.72.0 // indirect
	cloud.google.com/go/storage v1.12.0
	github.com/Condition17/fleet-services/file-service v0.0.0-20201124082821-bbccf128fb9b
	github.com/Condition17/fleet-services/lib v0.0.0-20201120170506-cd29a9cac69c
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4 // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)