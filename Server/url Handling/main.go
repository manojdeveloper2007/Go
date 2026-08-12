package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

const url = "http://localhost:8000/user"

func main() {

	//url details
	det, _ := url2.Parse(url)

	fmt.Println(det.User)
	fmt.Println(det.Host)
	fmt.Println(det.Hostname())
	fmt.Println(det.Port())
	fmt.Println(det.Scheme)
	fmt.Println(det.RawPath)
	fmt.Println(det.Path)
	fmt.Println(det.Fragment)

	data := url2.Values{}

	data.Set("name", "Manoj")
	response, err := http.PostForm(url, data)

	if err != nil {
		log.Fatal(err)
	}
	content, _ := io.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Println(string(content))
}
