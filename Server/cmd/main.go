package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	port := flag.String("port", "3000", "server port")
	flag.Parse()
	mux.HandleFunc("/home", Home)
	mux.HandleFunc("/user/Manoj", User)
	mux.HandleFunc("/user", Userdetails)
	server := http.ListenAndServe(":"+*port, mux)
	log.Println("Server started on Port ", *port)
	if server != nil {
		log.Fatal(server)
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Home Page"))
	if err != nil {
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	name := strings.TrimPrefix(path, "/user/")

	_, err := w.Write([]byte(fmt.Sprintf("Hello %v", name)))
	if err != nil {
		return
	}
}

func Userdetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")

		_, err := w.Write([]byte(fmt.Sprintf("Hello , %v", name)))
		if err != nil {
			return
		}
		return
	}

	http.Error(w, "Method Not Followed", http.StatusMethodNotAllowed)
	return
}
