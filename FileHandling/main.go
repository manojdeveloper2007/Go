package main

import (
	"fmt"
	"log"
	"os"
)

func createFile() {
	file, err := os.Create("./sample.txt")
	if err != nil {
		log.Fatal("Something wrong on creating a file")
	}
	log.Println("file created successfully")
	defer file.Close()
}

func readFile(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))
}

func main() {
	file := "./sample.txt"
	content := []byte("Hello,Iam Manoj.Iam a Software developer")
	err := os.WriteFile(file, content, 0644)
	if err != nil {
		panic(err)
	}
	readFile(file)
}
