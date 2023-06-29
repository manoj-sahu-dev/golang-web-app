package main

import (
	"log"

	"github.com/tsawler/myniceprogram/helpers"
)

const numPool = 10

func main() {
	/*
		var myVar helpers.SomeType
		myVar.TypeName = "Some name"

		log.Println(myVar.TypeName)
	*/

	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan

	log.Println(num)
}

func calculateValue(myChan chan int) {
	rnum := helpers.RandomNumber(numPool)
	myChan <- rnum
}
