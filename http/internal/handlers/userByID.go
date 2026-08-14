package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Users(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	us, exists := Alluser[id]

	if !exists {
		http.Error(w, "user not Found", http.StatusBadRequest)
		return
	}

	_, err := fmt.Fprintf(w, "User ID : %v\nName : %v", us.userid, us.name)
	if err != nil {
		panic("Something Wrong")
	}
}
