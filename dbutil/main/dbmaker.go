package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3",
		"./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		Id          int
		Name        string
		WeightClass string
		Status      string
		Wins        int
		Losses      int
		Events      string
	)
	rows, err := db.Query("select * from robots where wins is not null")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Id, &Name, &WeightClass, &Status, &Wins, &Losses, &Events)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(Id, Name, WeightClass, Status, Wins, Losses, Events)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
