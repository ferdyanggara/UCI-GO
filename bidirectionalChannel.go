package main

import "fmt"

func main() {
	n := 3
	in := make(chan int)
	out := make(chan int)

	go multiplyByTwo(in, out)
	in <- n
	fmt.Println(<-out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	number := <-in
	result := number * 2
	out <- result
}
