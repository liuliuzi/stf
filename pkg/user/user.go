package user

import (
    //"fmt"
    "encoding/json"
)
type User struct {
	Id, Name string
}

func (u User) String() string {
	j, _ := json.Marshal(u)
	return string(j)
}
