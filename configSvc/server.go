package configsvc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	cstore "github.com/pramineni01/madr/configstore"
	"github.com/pramineni01/madr/mockGrpc"
)

// type apiCall struct {
// 	seq int
// 	request proto.MessageV1
// 	response proto.MessageV1
// 	waitTime int
// }

const (
	ip_port = "127.0.0.1:5000"
)

type Server struct {
}

func NewServer() *Server {
	return new(Server)
}

var grpcSvr *mockgrpc.Server

func (s *Server) Start() {
	fmt.Println("Reached Server Start()")
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)
	e.POST("/mock/config", configMockService)
	e.GET("/mock/stop", stopGrpcSvc)

	// Start server
	fmt.Println("Starting server")
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func configMockService(eCtx echo.Context) error {
	// parse the json
	// jsonpb to Protos
	// insert data in apiCalls
	// var svc string
	// if svc = c.QueryParam("svc"); len(svc) < 1 {
	// 	return c.String(http.StatusBadRequest, fmt.Sprintf("Query Param 'svc' is missing"))
	// }
	fmt.Println("Service to mock: , svc")
	cs := cstore.NewConfigStore(eCtx)

	fmt.Println(cs.Dump())

	// stop the running service, if any
	if grpcSvr != nil {
		grpcSvr.Stop()
	}

	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	grpcSvr := mockSvr.NewServer(cancelCtx, cs)

	retCh, err := grpcSvr.Start(ip_port)
	time.Sleep(100 * time.Millisecond) // wait till grpc service is up

	go func() {
		done := <-retCh
		defer cancelFunc()
		grpcSvr.GracefulStop()
	}
	
	// TODO: check health and if started then return started, else error
	return eCtx.JSON(http.StatusCreated, cfg)
	
	
	// u := new(User)
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }
	// fmt.Println(u)
	// return c.JSON(http.StatusCreated, u)

	// return c.String(http.StatusOK, fmt.Sprintf("Hello World: %s", c.QueryParam("svc")))
}