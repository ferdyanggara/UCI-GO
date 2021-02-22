package main

import "fmt"

func main() {
	a := make(chan int)
	go timesTwo(5, a)
	fmt.Println(<-a)
}

func timesTwo(b int, a chan<- int) {
	result := b * 2
	a <- result
}
