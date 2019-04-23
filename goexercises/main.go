package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	csvfile, _ := os.Open("result.csv")
	reader := csv.NewReader(bufio.NewReader(csvfile))
	var people []Person

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		people = append(people, Person{
			FirstName: line[0],
			LastName:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}

	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))

	problemsfile, _ := os.Open("problems.csv")
	problemreader := csv.NewReader(bufio.NewReader(problemsfile))

}
