package user
import (
	"net/http"
	"github.com/emicklei/go-restful"
	"fmt"
	"github.com/liuliuzi/stf/pkg"
	"golang.org/x/net/context"
	"strings"
	//"github.com/emicklei/go-restful/swagger"
)

type UserService struct {
	// normally one would use DAO (data access object)
	Users map[string]User
	ApiRuntime   *pkg.ApiRuntime
	Prefix       string	
}
func (u UserService) Init(){
	u.Prefix="/stf/node/"
	fmt.Println(u.Prefix)
}	

func (u UserService) Register() {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	ws.Route(ws.GET("/").To(u.findAllUsers))
	ws.Route(ws.GET("/{user-id}").To(u.findUser)) // on the response
	ws.Route(ws.PUT("/{user-id}").To(u.updateUser)) // from the request
	ws.Route(ws.POST("/").To(u.createUser)) // from the request
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))
	restful.Add(ws)
}

func (u UserService) findAllUsers(request *restful.Request, response *restful.Response) {
	fmt.Println("findAllUsers")
	value, _ := u.storageGetKey()
	ret:="["
	for _, key := range value {
            fmt.Println( key)
            ret=ret+key
            key=strings.Replace(key,u.Prefix,"",20)
            value, _ := u.storageGet(key)
            //fmt.Println (value)
            ret=ret+value

        }
    ret=ret+"]"

	response.WriteEntity(ret)
}


func (u UserService) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	//usr := new(User)
	value, err := u.storageGet(id)
	if err != nil {
	    fmt.Println(err)
	}else{
	    fmt.Println(value)
	}
	
	if len(value) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		fmt.Println("findUser")
		response.WriteEntity(value)
	}
}

func (u *UserService) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.storageUpdate(usr.Id,usr.String())
		fmt.Println("updateUser")
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}


func (u *UserService) createUser(request *restful.Request, response *restful.Response) {

	usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		
		u.storageSet(usr.Id,usr.String())		
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (u *UserService) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	u.storageDelete(id)
	fmt.Println("removeUser")
}

func (u *UserService) storageSet(key string , value string ) error {
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

func (u *UserService) storageGetKey() ([]string ,error) {
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

func (u *UserService) storageGet(key string ) (string ,error) {
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
func (u *UserService) storageUpdate(key string , value string ) (string ,error) {
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
func (u *UserService) storageDelete(key string ) error {
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

