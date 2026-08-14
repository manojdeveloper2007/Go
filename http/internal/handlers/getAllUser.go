package handlers

import (
	"fmt"
	"net/http"
)

type User struct {
	userid int
	name   string
}

var Alluser = make(map[int]User)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	Alluser[15] = User{15, "Manoj"}
	Alluser[49] = User{49, "Sanjay"}

	for _, value := range Alluser {
		_, err := fmt.Fprintf(w, "User ID : %v\nName : %v\n\n", value.userid, value.name)
		if err != nil {
			return
		}
	}
}
