package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {

	db, err := sql.Open(
		"mysql",
		"root:root@tcp(172.19.0.2:3306)/go_api")
	if err != nil {
		panic(err)
		fmt.Println("faiiiiiiiiiiiiiiil")
	}
	defer db.Close()
	fmt.Println(db)
	fmt.Println("*****************************")
}