package main

import (
	"github.com/pramineni01/madr/configSvc"
)

func main() {
	restSrv := configApi.NewServer()
	restSrv.Start()
}
