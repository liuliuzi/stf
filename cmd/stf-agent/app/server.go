package app

import (
	"fmt"
	//"net/http"
	_ "net/http/pprof"
	"github.com/liuliuzi/stf/cmd/stf-agent/app/options"
	//"github.com/liuliuzi/stf/pkg"

	//"github.com/emicklei/go-restful"
	//"github.com/emicklei/go-restful/swagger"
	//"github.com/liuliuzi/stf/pkg/node"

)


//var apiRuntime,_=pkg.NewApiRuntime()

func Run(s *options.AgentServer) error {
	fmt.Println(" service in port =",s.Port)
	fmt.Println(" service in ServerIP =",s.ServerIP)
	fmt.Println(" service in LocalIP =",s.LocalIP)
/*
	apiRuntime.Etcdclient=pkg.Etcdclient("10.140.58.130",2379)
	u := node.NodeService{map[string]node.Node{},apiRuntime,"/stf/nodes/"}
	u.Register()
	config := swagger.Config{
		WebServices:    restful.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8090",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/Users/emicklei/Projects/swagger-ui/dist"}
	swagger.InstallSwaggerService(config)

	return http.ListenAndServe(":8090", nil)
*/
	return nil
}
