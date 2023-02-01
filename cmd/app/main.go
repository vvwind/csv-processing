package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"yadro/internal/services"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Filename not provided")
	}

	argsWithoutProg := os.Args[1]

	if len(argsWithoutProg) < 4 {
		log.Fatal("Incorrect data file")
	}

	file, err := os.Open(argsWithoutProg)
	if err != nil {
		log.Fatal(err)
	}
	var data = [][]string{}
	dct := make(map[string]int)
	reader := csv.NewReader(file)
	records, errReading := reader.ReadAll()

	if errReading != nil {
		log.Fatal("Got error when reading csv!: ", errReading)
	}
	if len(records) == 0 {
		log.Fatal("Empty file provided!")
	}
	for i, v := range records[0] {
		dct[v] = i
	}
	for i, v := range records {

		if i > 0 {
			index, iErr := strconv.Atoi(v[0])
			if iErr != nil {
				log.Panicln("Invalid index at row: ", v)
			}
			if index != i {
				log.Panicln("Invalid index at row: ", v)
			}
		}
		data = append(data, v)
	}

	dataStruct := services.Service{}
	dataStruct.InitService(dct, data)
	dataStruct.Iterator()
	for _, v := range dataStruct.Data {
		fmt.Println(strings.Join(v, ","))
	}

}
