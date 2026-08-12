package main

import (
	"fmt"
	"net/url"
)

const urlPath = "https://www.cricbuzz.com/profiles/13213/rahmanullah-gurbaz-wk?name=dhoni&team=CSK"

func main() {

	U, err := url.Parse(urlPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(U.Scheme)
	fmt.Println(U.Path)
	fmt.Println(U.RawPath)
	fmt.Println(U.Port())
	fmt.Println(U.Host)
	fmt.Println(U.User)
	fmt.Println(U.Fragment)
	fmt.Println(U.RawQuery)

	queryValues := U.Query()

	fmt.Println(queryValues["name"])

	for idx, value := range queryValues {
		fmt.Println(idx, " : ", value)
	}
}
