package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please input the acceleration: ")
	inputReader := bufio.NewReader(os.Stdin)
	accelerationString, err := inputReader.ReadString('\n')
	acceleration := strings.Trim(accelerationString, "\n")
	acc, err := strconv.ParseFloat(acceleration, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please input the  initial velocity: ")
	velocityReader := bufio.NewReader(os.Stdin)
	velString, err := velocityReader.ReadString('\n')
	velocity := strings.Trim(velString, "\n")
	vel, err := strconv.ParseFloat(velocity, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please input the initial distance: ")
	distanceReader := bufio.NewReader(os.Stdin)
	distanceString, err := distanceReader.ReadString('\n')
	distance := strings.Trim(distanceString, "\n")
	dist, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please input the time: ")
	timeReader := bufio.NewReader(os.Stdin)
	timeString, err := timeReader.ReadString('\n')
	timing := strings.Trim(timeString, "\n")
	time, err := strconv.ParseFloat(timing, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Results: ")
	FunctionHolder := GenDisplaceFn(acc, vel, dist)
	fmt.Println(FunctionHolder(time))
}

func GenDisplaceFn(a, Vo, So float64) func(float64) float64 {
	fn := func(t float64) float64 {
		first := 0.5 * a * math.Pow(t, 2)
		second := Vo * t
		return first + second + So
	}
	return fn
}
