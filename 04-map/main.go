package main

import (
	"log"
	"sort"
)

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

	// slices

	var mySlice []int
	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 8)
	mySlice = append(mySlice, 1)

	log.Println(mySlice)
	sort.Ints(mySlice)
	log.Println(mySlice)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	log.Println(nums)

	log.Println(nums[0:2])

}
