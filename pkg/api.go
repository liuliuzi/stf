package pkg
import (
    "fmt"
    "net/url"
    "errors"
    "net/http"

    
)
type Cmd func(r *http.Request,srv *Server) (error,[]byte)

func ApiSubHandler(r *http.Request, srv *Server) (error,Cmd){
    //fmt.Println("tset_ApiSubHandler")
    queryValues, _ := url.ParseQuery(r.URL.RawQuery)
    fmt.Println(r.Method,r.URL.Path,queryValues)
    Path:=r.URL.Path
    Method:=r.Method
    
    if  Path=="/api/node"{
        if Method=="GET" {
            return nil,srv.Apiruntime.NodeManager.GetNodeList
        }
        if Method=="POST" {
            return nil,srv.Apiruntime.NodeManager.AddNode///////////
        }
    }
    if  Path=="/api/node/detail"{
        if Method=="GET" {
            return nil,srv.Apiruntime.NodeManager.GetNodeListDetail
        }
    }
    if  Path=="/api/node/nodeid"{
        if Method=="GET" {
            return nil,srv.Apiruntime.NodeManager.GetNode/////////////
        }
        if Method=="PUT" {
            return nil,srv.Apiruntime.NodeManager.UpdateNode
        }
        if Method=="DELETE" {
            return nil,srv.Apiruntime.NodeManager.RemoveNode
        }
    }
    
    
    return errors.New("cannot find handle"),nil

}

type ApiRuntime struct {
    Root           string
    
    NodeManager    *NodeManager

}
func NewApiRuntime() (*ApiRuntime, error) {
    fmt.Println("NewApiRuntime123")
    apiRuntime := &ApiRuntime{
        Root:           "root2",
        NodeManager:    &NodeManager{"root1",nil},
    }
    apiRuntime.NodeManager.Init()
    fmt.Println(apiRuntime.Root)
    return apiRuntime, nil
}

func NewServer() (*Server, error) {

    apiruntime, err := NewApiRuntime()

    if err != nil {
        return nil, err
    }
    srv := &Server{
        Apiruntime: apiruntime,
        
    }
    return srv, nil
}

type Server struct {
    Apiruntime *ApiRuntime
}
