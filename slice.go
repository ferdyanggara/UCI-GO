package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	// "strconv"
)

func main() {
	intSlice := make([]int, 0, 3)
	for {
		fmt.Println("Enter an integer: ('X' for exit)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "X" {
			fmt.Println("Exiting the program")
			break
		} else {
			check, err := strconv.Atoi(input)
			if err == nil {
				intSlice = append(intSlice, check)
				sort.Ints(intSlice)
				fmt.Println(intSlice)
			} else {
				fmt.Println("Input a valid number")
				continue
			}
		}

	}
}
