package main

import "fmt"

func main() {
	fmt.Println(getMax(1, 2, 5, 6))
	vSlice := []int{1, 3, 6, 8}
	fmt.Println(getMaxSlice(vSlice))
}

func getMax(vals ...int) int {
	vMax := -1
	for _, v := range vals {
		if v > vMax {
			vMax = v
		}
	}
	return vMax
}

func getMaxSlice(vals []int) int {
	vMax := -1
	for _, v := range vals {
		if v > vMax {
			vMax = v
		}
	}
	return vMax
}
