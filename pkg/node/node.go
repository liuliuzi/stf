package node

import (
    //"fmt"
    "encoding/json"
)
type Node struct {
	Id, Name ,HostName, Status, Timestamp string

}


func (u Node) String() string {
	j, _ := json.Marshal(u)
	return string(j)
}
