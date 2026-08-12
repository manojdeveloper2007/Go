package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.cricbuzz.com"

	res, err := http.Get(url)

	if err != nil {
		panic("url failed to load")
	}

	ans, err1 := io.ReadAll(res.Body)

	if err1 != nil {
		panic("Something wrong")
	}

	fmt.Println(string(ans))

	defer res.Body.Close()
}
