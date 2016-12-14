package node
import (
	"net/http"
	"github.com/emicklei/go-restful"
	"fmt"
	"github.com/liuliuzi/stf/pkg"
	"golang.org/x/net/context"
	"strings"
	//"encoding/json"
	//"github.com/bitly/go-simplejson"
	//"github.com/emicklei/go-restful/swagger"
)

type NodeService struct {
	// normally one would use DAO (data access object)
	Nodes map[string]Node
	ApiRuntime   *pkg.ApiRuntime
	Prefix       string
}

func (u NodeService) Init(){
	u.Prefix="/stf/node/"
	fmt.Println(u.Prefix)
}

func (u NodeService) Register() {
	ws := new(restful.WebService)
	ws.
		Path("/nodes").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	ws.Route(ws.GET("/").To(u.findAllNodes))
	ws.Route(ws.GET("/{Node-id}").To(u.findNode)) // on the response
	ws.Route(ws.PUT("/{Node-id}").To(u.updateNode)) // from the request
	ws.Route(ws.POST("/").To(u.createNode)) // from the request
	ws.Route(ws.DELETE("/{Node-id}").To(u.removeNode))
	restful.Add(ws)
}

func (u NodeService) findAllNodes(request *restful.Request, response *restful.Response) {
	fmt.Println("findAllNodes")
	value, _ := u.storageGetKey()
	ret:="["
	for _, key := range value {
            fmt.Println( key)
            ret=ret+`"`+key+`":`
            key=strings.Replace(key,u.Prefix,"",20)
            value, _ := u.storageGet(key)
            //fmt.Println (value)
            ret=ret+value

        }
    ret=ret+"]"

	response.WriteEntity(ret)
}


func (u NodeService) findNode(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("Node-id")
	value, err := u.storageGet(id)
	if err != nil {
	    fmt.Println(err)
	}else{
	    fmt.Println(value)
	}

	if len(value) == 0 {
		response.WriteErrorString(http.StatusNotFound, "Node could not be found.")
	} else {
		fmt.Println("findNode")
		response.WriteEntity(value)
	}
}

func (u *NodeService) updateNode(request *restful.Request, response *restful.Response) {
	nod := new(Node)
	err := request.ReadEntity(&nod)
	if err == nil {
		u.storageUpdate(nod.Id,nod.String())
		fmt.Println("updateNode")
		response.WriteEntity(nod)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}


func (u *NodeService) createNode(request *restful.Request, response *restful.Response) {

	nod := new(Node)
	err := request.ReadEntity(&nod)
	if err == nil {
		//u.storageSet(nod.Id,nod)
		u.storageSet(nod.Id,nod.String())
		//u.storageSet(nod.Id,{'Id':'123566666','Name':'liu'})

		//u.storageSet(nod.Id,kk)
		response.WriteHeaderAndEntity(http.StatusCreated, nod)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (u *NodeService) removeNode(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("Node-id")
	u.storageDelete(id)
	fmt.Println("removeNode")
}

func (u *NodeService) storageSet(key string , value string ) error {
	resp, err := u.ApiRuntime.Etcdclient.Set(context.Background(), u.Prefix+key, value, nil)
    if err != nil {
    	fmt.Println("storageSet failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageSet success")
    	fmt.Println(resp)
    }
    return err
}

func (u *NodeService) storageGetKey() ([]string ,error) {
	resp, err := u.ApiRuntime.Etcdclient.Get(context.Background(), u.Prefix, nil)
    if err != nil {
    	fmt.Println("storageGetKey failed")
        fmt.Println(err)
        return nil,err
    }else{
    	ret:=[]string{}
    	fmt.Println("storageGetKey success")
    	for _, node := range resp.Node.Nodes {
            fmt.Println( node)
            ret=append(ret,node.Key)
        }

        return ret,nil

    }
}

func (u *NodeService) storageGet(key string ) (string ,error) {
	fmt.Println("storageGet", key)
	resp, err := u.ApiRuntime.Etcdclient.Get(context.Background(), u.Prefix+key, nil)
    if err != nil {
    	fmt.Println("storageGet failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageGet success")
    	fmt.Println(resp)
    }
    return resp.Node.Value,err
}
func (u *NodeService) storageUpdate(key string , value string ) (string ,error) {
	resp, err := u.ApiRuntime.Etcdclient.Update(context.Background(), u.Prefix+key,value)
    if err != nil {
    	fmt.Println("storageUpdate failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageUpdate success")
    	fmt.Println(resp)
    }
    return resp.Node.Value,err
}
func (u *NodeService) storageDelete(key string ) error {
	resp, err := u.ApiRuntime.Etcdclient.Delete(context.Background(), u.Prefix+key, nil)
    if err != nil {
    	fmt.Println("storageDelete failed")
        fmt.Println(err)
        return err
    }else{
    	fmt.Println("storageDelete success")
    	fmt.Println(resp)
    	return nil
    }
}

