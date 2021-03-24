package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	currentTime := time.Now()

	type passport struct {
		series string
		number string
	}
	//Способ 1, чтение файла
	/*
		//passports := make([]passport, 0)
		file, err := ioutil.ReadFile("list_of_expired_passports.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		passportline := strings.Split(string(file), "\n")

		for i := 0; i < len(passportline); i++ {

			if passportline[i] != "" {

				CurrentPassport := strings.Split(string(passportline[i]), ",")

				newPassport := passport{series: CurrentPassport[0], number: CurrentPassport[1]}
				//passports = append(passports, newPassport)
				fmt.Println(newPassport.series + " " + newPassport.number)
			}
		}

		currentTime2 := time.Now()
		fmt.Println(currentTime)
		fmt.Println(currentTime2)
		//4 минуты 30 секунд
	*/

	//Способ 2, чтение файла CSV
	file, err := os.Open("list_of_expired_passports.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	for {
		CurrentPassport, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		newPassport := passport{series: CurrentPassport[0], number: CurrentPassport[1]}
		//passports = append(passports, newPassport)
		fmt.Println(newPassport.series + " " + newPassport.number)
	}

	currentTime2 := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime2)
	//4 минуты 30 секунд
}
