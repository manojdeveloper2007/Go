package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Home Page")
	if err != nil {
		return
	}
}
