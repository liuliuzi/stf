package app

import (
	"fmt"
	"net/http"
	
	_ "net/http/pprof"
	"liuyaoting.io/stf/cmd/stf-apiserver/app/options"
	"liuyaoting.io/stf/pkg"
	

)


func handler(w http.ResponseWriter, r *http.Request) {
	
	
	fmt.Println("-------------------------------")
	//fmt.Println(r.Method,r.URL.Path,queryValues)
	//defer r.Body.Close()
	//body, _ := ioutil.ReadAll(r.Body)
	//js, err := simplejson.NewJson(body)
	//fmt.Println(string(body))
	//arr := js.Get("kind")
	//fmt.Println(arr)
	err,ApiSubHandler_fun:=pkg.ApiSubHandler(r,srv)
	if err != nil {
		fmt.Println(err)
		return 
	}
	err,rep:=ApiSubHandler_fun(r,srv)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Fprintf(w,string(rep))
	

}
/*func http_server(port int) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)

}*/

var srv,_=pkg.NewServer()

func Run(s *options.APIServer) error {
	fmt.Println("start service in port ",s.Port)
	
	//nodeManager:=NodeManager{"root",nil}
	http.HandleFunc("/", handler)
	return http.ListenAndServe(":8090", nil)

}


