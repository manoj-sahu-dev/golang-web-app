package main

import (
	"errors"
	"log"
)

func main() {

	var one float32 = 10.0
	var two float32 = 0.0

	result, err := divide(one, two)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(result)
}

func divide(a float32, b float32) (float32, error) {
	var result float32
	if b == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = a / b
	return result, nil
}
