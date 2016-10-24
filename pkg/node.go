package pkg
import (
    "fmt"
    //"net/url"
    "github.com/coreos/etcd/client"
    "golang.org/x/net/context"
    "net/http"
    "github.com/bitly/go-simplejson"
    "io/ioutil"
    //"github.com/golang/glog"

)

type NodeManager struct {
    Root           string
    Etcdclient     client.KeysAPI


}
func (nm *NodeManager)Init() {
    nm.Etcdclient=Etcdclient("127.0.0.1",2379)
    fmt.Println("NodeManager init")
}

func (nm *NodeManager)AddNode(r *http.Request, srv *Server) (error,[]byte) {    
    defer r.Body.Close()
    body, _ := ioutil.ReadAll(r.Body)
    js, _ := simplejson.NewJson(body)
    nodeName,err := js.Get("metadata").Get("name").String()
    if err != nil {
        return err,nil
    } 
    etcdNodeName:="/stf/nodes/"+nodeName
    fmt.Println(etcdNodeName)
    resp, err := nm.Etcdclient.Set(context.Background(), etcdNodeName, string(body), nil)
    if err != nil {
        fmt.Println(resp)
        return err,nil
    }
    return nil,[]byte("node add success")
}

func (nm *NodeManager) GetNodeList(r *http.Request, srv *Server) (error,[]byte) {
    resp, err := nm.Etcdclient.Get(context.Background(), "/stf/nodes/", nil)
    if err != nil {
        fmt.Printf("errr %q\n", resp)
        return nil,nil
    } else {
        for _,node:=range  resp.Node.Nodes{
            fmt.Println(node.Key)
        }
        
    }
    return nil,[]byte("test233333")
}


func (nm *NodeManager) GetNodeListDetail(r *http.Request, srv *Server) (error,[]byte) {
    return nil,[]byte("test5555555552")
}

func (nm *NodeManager)GetNode(r *http.Request, srv *Server) (error,[]byte) {
    
    resp, err := nm.Etcdclient.Get(context.Background(), "/stf/nodes/"+"10.254.177.223", nil)
    if err != nil {
        fmt.Printf("errr %q\n", resp)
    } else {
        fmt.Println(resp.Node.Value)
    }
    return nil,[]byte(resp.Node.Value)
}
func (nm *NodeManager)UpdateNode(r *http.Request, srv *Server) (error,[]byte) {
    return nil,[]byte("test333")
}
func (nm *NodeManager)RemoveNode(r *http.Request, srv *Server) (error,[]byte) {
    return nil,[]byte("test444")
}
