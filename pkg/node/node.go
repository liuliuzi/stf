package node

import (
    //"fmt"
    "encoding/json"
)
type Node struct {
	Id, Name string
}

func (u Node) String() string {
	j, _ := json.Marshal(u)
	return string(j)
}
