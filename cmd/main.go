package main

import (
	"fmt"
	"github.com/pramineni01/madr/configSvc"
)

func main() {
	fmt.Println("Reached Main")
	restSrv := configApi.NewServer()
	restSrv.Start()
}
