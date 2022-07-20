package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:areyh817@tcp(127.0.0.1:3308)/goproject")
	if err != nil {
		log.Fatal(err)
	} else {
		print("성공입니다.")
	}

	var name string

	err = db.QueryRow("SELECT name FROM test WHERE id=1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	defer db.Close()
}
