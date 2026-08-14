package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	delete(Alluser, id)
	_, err := fmt.Fprintf(w, "user id %v deleted successfully", id)
	if err != nil {
		return
	}
}
