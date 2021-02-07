package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "Joe", age: 15}
	fmt.Println(p1)
}
