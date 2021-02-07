package main

import (
	"fmt"
)

func main() {
	var yourMap map[string]string
	yourMap = make(map[string]string)
	yourMap["Joe"] = "Johnson"

	// another way to create it is using map literal
	myMap := map[string]string{
		"Joe": "Johnson"}

	// fmt.Println(myMap["Joe"])
	// fmt.Println("his name is ", yourMap["Joe"])
	// delete(myMap, "Joe")
	for key, value := range myMap {
		fmt.Println("key: ", key, "value: ", value)

	}
}
