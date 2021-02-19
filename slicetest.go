package main

import "fmt"

func main() {
	readers := [4]string{"Alice", "Bob", "Charlie", "Dytto"}

	a := readers[0:2]
	b := readers[1:3]
	fmt.Println("Before Modification")
	fmt.Println("Readers", readers)
	fmt.Println("a:", a, "b:", b)

	b[0] = "Common"
	fmt.Println("After Modification")
	fmt.Println("Readers", readers)
	fmt.Println("a:", a, "b:", b)
}
