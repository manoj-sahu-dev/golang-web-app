package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	HasDog    bool   `json:"hasDog"`
}

func main() {
	myJson := `[
		{
			"firstName": "Manoj",
			"lastName": "Sahu",
			"hasDog":true
		},
		{
			"firstName": "Saroj",
			"lastName": "Sahu",
			"hasDog":false
		}
	]`
	var data []Person
	err := json.Unmarshal([]byte(myJson), &data)
	if err != nil {
		log.Println("got error wihile converting: " + err.Error())
	}
	log.Println(data)
	log.Printf("data: %v", data)

	person1 := Person{FirstName: "Ramesh", LastName: "Singh", HasDog: false}
	person2 := Person{FirstName: "Sumesh", LastName: "Yadav", HasDog: true}

	var myslice []Person
	myslice = append(myslice, person1)
	myslice = append(myslice, person2)

	mjson, jerr := json.Marshal(myslice)
	if jerr != nil {
		log.Fatalln("error marshalling: " + jerr.Error())
	}
	log.Println(string(mjson))
}
