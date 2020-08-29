package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://vttkxspt:sjA5CdRG1tepOQye8KB1ZMsPjQZ273V9@lallah.db.elephantsql.com:5432/vttkxspt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	log.Println("Okay")

	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}
	if _, err := stmt.Exec(1, "inactive"); err != nil {
		log.Fatal("error execute update ", err)
	}
	fmt.Println("update success")

}
