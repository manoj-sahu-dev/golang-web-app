package main

import "log"

type User struct {
	firstName string
	lastName  string
}

func main() {
	mymap := make(map[string]string)
	mymap["dog"] = "cat"
	log.Println(mymap["dog"])

	aMap := make(map[string]User)
	aMap["dog"] = User{firstName: "wow dog", lastName: "catcat"}
	log.Println(aMap)
}
