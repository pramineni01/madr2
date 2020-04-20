module mockGrpc

go 1.14

replace github.com/pramineni01/madr/proto/keyvaluestore => ../proto/keyvaluestore
require (
    google.golang.org/protobuf v1.21.0
    github.com/pramineni01/proto/keyvaluestore v1.0.0
)
