package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {

	currentTime := time.Now()

	password := ""
	connStr := "user=postgres password=" + password + " dbname=passport_DB sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("COPY docs FROM 'C:/12/ft.csv' WITH CSV HEADER")
	if err != nil {
		panic(err)
	}
	currentTime2 := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime2)
	//4 минуты 30 секунд
}
