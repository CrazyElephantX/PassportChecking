package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {

	currentTime := time.Now()

	/*type passport struct {
		series string
		number string
	}*/

	//Способ 2, чтение файла CSV
	file, err := os.Open("list_of_expired_passports.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	password := ""
	connStr := "user=postgres password=" + password + " dbname=passport_DB sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for {
		CurrentPassport, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		//newPassport := passport{series: CurrentPassport[0], number: CurrentPassport[1]}
		//passports = append(passports, newPassport)
		//fmt.Println(newPassport.series + " " + newPassport.number)
		if CurrentPassport[0] != "PASSP_SERIES" {
			//result, err := db.Exec("insert into docs (pas_series, pas_number) values ($1,$2)",CurrentPassport[0], CurrentPassport[1])
			db.Exec("insert into docs (pas_series, pas_number) values ($1,$2)", CurrentPassport[0], CurrentPassport[1])
			/* if err != nil{
				panic(err)
			}
			fmt.Println(result) */

		}

	}

	currentTime2 := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime2)
	//4 минуты 30 секунд
}
