package mockgrpc

import (
	"context"
	"errors"
	"net"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	cfg "github.com/pramineni01/madr/configstore"
	kvs "github.com/pramineni01/madr/proto/keyvaluestore"
)

// Server is grpc mocker server
type Server struct {
	server *kvs.KeyValueStoreServer
	cfg    cfg.Interface
}

// NewServer gives an instance of mock grpc server
func NewServer(ctx context.Context, c cfg.Interface) *Server {
	return &Server{cfg: c}
}

// Start the mock grpc service
func (ms *Server) Start(ipPort string) (<-chan bool, error) {
	done := make(chan bool)
	defer func() {
		done <- true
	}()

	ms.cfg = NewConfigData(cfg)

	lis, err := net.Listen("tcp", ipPort)
	if err != nil {
		return done, errors.New("Error opening port to listen. Err: %v", err)
	}

	s := grpc.NewServer()
	kvs.RegisterKeyValueStoreServer(s, &MockServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			return done, errors.New("Error starting the server. Err: %v", err)
		}
	}()
}

// GetValues An api call. All the api calls looks like this.
func (ms *Server) GetValues(ctx context.Context, req *kvs.Request) (res *kvs.Response, err error) {
	res, err := handleRequest(req)
	if err == nil {
		// print and handle error
	}
	return res, err

}

// helper function
func (ms *Server) handleRequest(req *proto.Message) (res *proto.Message, err error) {

	// marshal request
	marshaler := jsonpb.Marshaler{}
	actualReqString, err := marshaler.MarshalToString(req)
	if err != nil {
		// handle error
	}

	// get the apiCall data and compare the apiCall request with marshaled request
	currAPICall, err := ms.cfg.GetNextAPICallInfo()
	if err != nil {
		// handle error
	}
	if currAPICall.Request != actualReqString {
		//return nil, errors.New("Request does not match. Expected: %s; Actual: %s", currApiCall.Request, actualReqString)
		// handle error
	}

	// if matches unmarshal the response and return it
	var p proto.Message
	if err := jsonpb.UnmarshalString(currApiCall.Response, &p); err != nil {
		return nil, err // error here
	}

	return &p, nil
}

// func (ms *Server) GetValues(ctx context.Context, req *kvs.Request) (res *kvs.Response, err error) {

// 	// marshal request
// 	marshaler := jsonpb.Marshaler{}
// 	actualReqString, err := marshaler.MarshalToString(req)
// 	if err != nil {
// 		// handle error
// 	}

// 	// get the apiCall data and compare the apiCall request with marshaled request
// 	currAPICall, err := ms.cfg.GetNextAPICallInfo()
// 	if err != nil {
// 		// handle error
// 	}
// 	if currAPICall.Request != actualReqString {
// 		//return nil, errors.New("Request does not match. Expected: %s; Actual: %s", currApiCall.Request, actualReqString)
// 		// handle error
// 	}

// 	// if matches unmarshal the response and return it
// 	var p kvs.Response
// 	if err := jsonpb.UnmarshalString(currApiCall.Response, &p); err != nil {
// 		return nil, err // error here
// 	}

// 	return &p, nil
// }
