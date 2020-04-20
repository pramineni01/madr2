package configstore

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

// Interface is an interface for configdata
type Interface interface {
	GetNextAPICallInfo() (*APICall, error)
	GetServiceName() string
}

type configStore struct {
	Svc      string    `json: "svc"`
	APICalls []APICall `json: "apiCalls"`
	Index    int
}

// APICall holds mock info for an expected api call
type APICall struct {
	Seq      int    `json: "seq"`
	Request  string `json: "request"`
	Response string `json: "response"`
}

// NewConfigStore returns next expected api call data
func NewConfigStore(ectx echo.Context) (Interface, error) {
	var cfg configStore
	if err := ectx.Bind(&cfg); err != nil {
		return nil, err
	}
	cfg.Index = 0

	return &cfg, nil
}

// GetServiceName returns grpc service name
func (c *configStore) GetServiceName() string {
	return c.Svc
}

// GetNextAPICallInfo() interface function to return next api call info
func (c *configStore) GetNextAPICallInfo() (*APICall, error) {
	// return next API Call
	if c.Index >= len(c.APICalls) {
		return nil, errors.New("API Calls exceeded than actual configuration provided")
	}

	defer func() {
		c.Index++
	}()
	return &c.APICalls[c.Index], nil
}

func (c *configStore) Dump() {
	fmt.Println(u)
}
