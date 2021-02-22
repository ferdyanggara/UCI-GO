package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	in := make(chan int)
	out := make(chan int)

	go multipleByTwo(in, out)
	go multipleByTwo(in, out)
	go multipleByTwo(in, out)

	in <- a
	in <- b
	in <- c

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func multipleByTwo(in <-chan int, out chan<- int) {
	number := <-in
	result := number * 2
	out <- result
}
