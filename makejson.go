package main

import (
	"bufio"
	"fmt"
	"os"
	// "sort"
	"encoding/json"
	// "strconv"
)

type Person struct {
	name    string
	address string
}

func main() {
	var p1 Person
	fmt.Println("What is your name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputName := scanner.Text()
	p1.name = inputName
	fmt.Println("What is your address?")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	inputAdd := scanner2.Text()
	p1.address = inputAdd
	fmt.Println("what is p1 : ", p1)

	credenMap := make(map[string]string)
	credenMap["address"] = p1.address
	credenMap["name"] = p1.name

	barr, err := json.Marshal(credenMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(barr))
}
