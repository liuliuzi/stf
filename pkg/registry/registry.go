package registry

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
	"io/ioutil"
	"github.com/liuliuzi/stf/pkg/util"
)


type registrar struct{
	name string
}

func NewRegistrar() registrar {
	return registrar{name:"test"}
}


func (r *registrar) AddNode(dstIp string, dstPort string,srcIp string) error{
	path:="/nodes"
	hostname,err:=util.GetHostName()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(hostname)
	nowTime:=time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(nowTime)
	nodeInfo := `{"Id": "`+srcIp+`","Name": "test","HostName": "` + hostname + `","Status": "active","Timestamp": "` + nowTime + `"}`
	fmt.Println(nodeInfo)
	//nodeInfo:=`{"Id": "123","Name": "test","HostName": "ssssss","Status": "active","Timestamp": "20000111"}`
	resp, err := http.Post("http://"+dstIp+":"+dstPort+path,
		"application/json",
		strings.NewReader(nodeInfo))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(body))
	return err
}