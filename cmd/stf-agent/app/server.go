package app

import (
	"fmt"
	//"net/http"
	_ "net/http/pprof"
	"github.com/liuliuzi/stf/cmd/stf-agent/app/options"
	"github.com/liuliuzi/stf/pkg/registry"
	"strconv"

)


//var apiRuntime,_=pkg.NewApiRuntime()

func Run(s *options.AgentServer) error {
	fmt.Println(" service in port =",strconv.Itoa(s.Port))
	fmt.Println(" service in ServerIP =",s.ServerIP)
	fmt.Println(" service in LocalIP =",s.LocalIP)
	register:=registry.NewRegistrar()
	err:=register.AddNode(s.ServerIP,strconv.Itoa(s.Port),s.LocalIP)
	if err!=nil {
		fmt.Println("eeeeeeeeee")
	}

	return nil
}
