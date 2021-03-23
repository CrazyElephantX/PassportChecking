package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {

	currentTime := time.Now()

	type passport struct {
		series string
		number string
	}

	//passports := make([]passport, 0)
	file, err := ioutil.ReadFile("list_of_expired_passports_1.csv")

	if err != nil {
		log.Fatal(err)
	}

	passportline := strings.Split(string(file), "\n")

	for i := 0; i < len(passportline); i++ {

		if passportline[i] != "" {

			CurrentPassport := strings.Split(string(passportline[i]), ",")

			newPassport := passport{series: CurrentPassport[0], number: CurrentPassport[1]}
			//passports = append(passports, newPassport)
			fmt.Println(newPassport.series + newPassport.number)
		}
	}

	currentTime2 := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime2)

}
