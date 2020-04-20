module github.com/pramineni01/madr/mockgrpc

go 1.14

replace github.com/pramineni01/madr/configstore => ../configstore
replace github.com/pramineni01/madr/proto/keyvaluestore => ../proto/keyvaluestore

require (
	github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0
	github.com/pramineni01/madr/configstore v1.0.0
	github.com/pramineni01/madr/proto/keyvaluestore v1.0.0
	google.golang.org/grpc v1.28.1
	google.golang.org/protobuf v1.21.0
)
