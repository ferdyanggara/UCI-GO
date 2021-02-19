package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	intList := make([]int, 0, 10)
	fmt.Println("Welcome, please enter the sequence of the integer (pls put space between integer) ")
	sequence, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	trimmedSequence := strings.Trim(sequence, " \n")
	s := strings.Split(trimmedSequence, " ")
	fmt.Println("what is s :", s)
	for i, v := range s {
		if i < 10 {
			getConv, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			intList = append(intList, getConv)
			bubbleSort(intList)
		}
	}
	fmt.Println(intList)
}

func bubbleSort(intlist []int) {

	for i := 0; i < len(intlist)-1; i++ {
		for j := 0; j < len(intlist)-i-1; j++ {
			if intlist[j] > intlist[j+1] {
				// Swap(&intlist[j], &intlist[j+1])
				Swap(intlist, j)
			}
		}
	}
}

// func Swap(num1 *int, num2 *int) {
// 	temp := *num1
// 	*num1 = *num2
// 	*num2 = temp
// }

func Swap(intlist []int, j int) {
	temp := intlist[j]
	intlist[j] = intlist[j+1]
	intlist[j+1] = temp
}
