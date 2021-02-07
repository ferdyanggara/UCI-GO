package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// var stringToCheck string
	// fmt.Scanln(&stringToCheck)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if strings.Contains(input, "i") || strings.Contains(input, "a") || strings.Contains(input, "n") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
