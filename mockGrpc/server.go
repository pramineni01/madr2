package mockServer

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/golang/protobuf/jsonpb"
	kvs "github.com/pramineni01/madr/proto/keyvaluestore"
	proto "github.com/golang/protobuf/proto"
)

type IConfig interface {
	GetNextApiCallInfo() string, error
}

type configData struct {
	svc string			`json: "svc"`
	ApiCalls []apiCall	`json: "apiCalls"`
	index int
}

type apiCall struct {
	Seq int			`json: "seq"`
	Request string	`json: "request"`
	Response string	`json: "response"`
}

func NewConfigData(cfg string) IConfig {
	return &configData{cfg: cfg, index: 0}
}

func(c config) GetNextApiCallInfo() string, error {
	// return next API Call
	if index >= len(apiCalls) {
		return nil, errors.New("API Calls exceeded than actual configuration provided")
	}
	
	defer c.index++
	return c.cfg[index]
}

type Container struct {
	Option Options `protobuf:"varint,1,opt,name=option,enum=main.Options" json:"option,omitempty"`
}
type Server struct {
	server *kvs.KeyValueStoreServer
	cfg configData
}

func NewServer(ctx context.Context) *Server {
	var c configData
	if err := ctx.Bind(&c); err != nil {
		return err
	}
	
	fmt.Println(c.ApiCalls)
	return &Server{cfg: c}
}

func (ms *Server) GetValues(ctx context.Context, req *kvs.Request) (res *kvs.Response, err error) {

	// marshal request
	marshaler := jsonpb.JSONPBMarshaler{}
	actualReqString := marshaler.MarshalToString(req)
	
	// get the apiCall data and compare the apiCall request with marshaled request
	currApiCall := ms.cfg.GetNextApiCallInfo()
	if currApiCall.Request != actualReqString {
		return nil, errors.New("Request does not match. Expected: %s; Actual: %s", currApiCall.Request, actualReqString)
	}

	// if matches unmarshal the response and return it
	p := proto.Message{}
	if err := jsonpb.UnmarshalString(currApiCall.Response, p); err != nil {
		return nil, err // error here
	}

	return &p, nil
}

func (ms *Server) Start(ip_port string) <-chan bool, error {
	done := make(chan bool)
	defer func() {
		done <- true
	}

	ms.cfg := NewConfigData(cfg)
	
	lis, err := net.Listen("tcp", ip_port)
	if err != nil {
		return done, errors.New("Error opening port to listen. Err: %v", err)
	}

	s := grpc.NewServer()
	kvs.RegisterKeyValueStoreServer(s, &MockServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			return done, errors.New("Error starting the server. Err: %v", err)
		}
	}
}