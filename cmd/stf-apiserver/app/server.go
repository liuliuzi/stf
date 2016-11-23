package app

import (
	"fmt"
	"net/http"
	
	_ "net/http/pprof"
	"github.com/liuliuzi/stf/cmd/stf-apiserver/app/options"
	"github.com/liuliuzi/stf/pkg"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/liuliuzi/stf/pkg/user"
	

)


var apiRuntime,_=pkg.NewApiRuntime()

func Run(s *options.APIServer) error {
	fmt.Println("start service in port ",s.Port)
	
	

	apiRuntime.Etcdclient=pkg.Etcdclient("10.140.58.130",2379)
	u := user.UserService{map[string]user.User{},apiRuntime,"/stf/nodes/"}
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

}


