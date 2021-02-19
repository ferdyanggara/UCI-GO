package main

import (
	"fmt"
	"time"
)

func gobOne() {
	fmt.Println("I never print")
	time.Sleep(time.Second * 2)
	fmt.Println("I never print")
}

func main() {
	go gobOne()
	fmt.Println("I print 1st")
}
