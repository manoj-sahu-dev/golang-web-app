package main

import "fmt"

type myStruct struct {
	firstName string
}

func (s *myStruct) printName() {
	fmt.Println(s.firstName)
}

func main() {
	v1 := new(myStruct)
	v1.firstName = "Manoj"
	fmt.Println(v1.firstName)

	v2 := myStruct{
		firstName: "Saroj",
	}
	fmt.Println(v2.firstName)
	v2.printName()
}
