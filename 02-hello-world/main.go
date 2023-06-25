package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	birthday    time.Time
}

func main() {
	/*
		mystring := "green"
		log.Println("my string is set to " + mystring)
		changeUsingPointer(&mystring)
		log.Println("my string is changed to " + mystring)

		word1, word2 := saySomething()
		fmt.Println(word1)
		fmt.Println(word2)
	*/

	user := User{
		firstName: "Manoj",
		lastName:  "Sahu",
	}
	fmt.Println(user.firstName + " : " + user.lastName + user.phoneNumber + " : ")
	fmt.Println(user.birthday)
}

func changeUsingPointer(s *string) {
	fmt.Println("location of s is ", s)
	newValue := "red"
	*s = newValue
}

func saySomething() (string, string) {
	return "hello", "world"
}
