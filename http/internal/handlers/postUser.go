package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func PostUser(w http.ResponseWriter, r *http.Request) {

	files := []string{"../ui/html/Post.html"}

	tmpl, err1 := template.ParseFiles(files...)

	if r.Method == http.MethodGet {
		err2 := tmpl.ExecuteTemplate(w, "post", nil)
		if err2 != nil {
			return
		}
		http.Redirect(w, r, "/post", http.StatusSeeOther)
		return
	}

	if err1 != nil {
		http.Error(w, "template Not Found", http.StatusBadRequest)
		log.Println(err1)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")

	Alluser[id] = User{id, name}
	fmt.Fprintf(w, "User %v updated Successfully", name)
	return
}
