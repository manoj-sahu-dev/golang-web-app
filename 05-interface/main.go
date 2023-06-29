package main

import "log"

type Animal interface {
	says() string
	numberOfLegs() int
}

type Dog struct {
	name string
	age  int
}

type Gorilla struct {
	name string
	food string
}

func main() {
	dog := Dog{
		name: "Doggo",
		age:  5,
	}

	gorilla := Gorilla{
		name: "Gorilla",
		food: "Fishes",
	}

	log.Println(dog)
	log.Println(gorilla)

	printName(dog)
}

func printName(A Animal) {
	log.Println("A has ", A.numberOfLegs(), " Legs and says "+A.says())
}

func (d Dog) says() string {
	return "i am doggo!"
}
func (d Dog) numberOfLegs() int {
	return 4
}
