package main

import (
	"flag"
	"fmt"
	"http/internal/handlers"
	"log"
	"net/http"
)

func main() {
	//	creating multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/user/{id}", handlers.Users)
	mux.HandleFunc("/user", handlers.GetAllUser)
	mux.HandleFunc("/post", handlers.PostUser)
	mux.HandleFunc("/delete/{id}", handlers.DeleteUser)

	port := flag.String("port", "8000", "Server port")
	flag.Parse()

	log.Println("Server Started on ", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), mux)

	if err != nil {
		log.Fatal("something wrong")
	}
}
