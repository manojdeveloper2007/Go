package main

import (
	"Database/database"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := flag.String("dsn", "root:root@tcp(localhost:3306)/golang?parseTime=true", "data source name")
	flag.Parse()

	db, err := database.Opendb(*dsn)

	if err != nil {
		log.Fatal("Something wrong on database ", err)
	}

	var name string
	var age int

	exec := db.QueryRow("SELECT * FROM User WHERE name = (?)", "Manoj")

	err1 := exec.Scan(&name, &age)
	if err1 != nil {
		return
	}

	log.Printf("User :\nName : %v\nAge : %v\n", name, age)
	defer db.Close()
}
