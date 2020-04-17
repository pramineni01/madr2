package configApi

import (
	"net/http"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	"github.com/pramineni01/madr/madr"
	proto "google.golang.org/protobuf/runtime/protoiface"
)

struct apiCall {
	seq number
	request proto.MessageV1
	response proto.MessageV1
	waitTime number
}

apiCalls := make([]apiCall)

struct Server {
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) Start() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.GET("/", sayHello)
	e.POST("/adr/config_mock_service", configMockService)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func configMockService() error {
	// parse the json
	// jsonpb to Protos
	// insert data in apiCalls
}