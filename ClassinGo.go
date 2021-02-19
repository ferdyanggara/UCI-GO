package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	// "strconv"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (cow *Animal) InitCow() {
	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"
}

func (bird *Animal) InitBird() {
	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"
}

func (snake *Animal) InitSnake() {
	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{}
	bird := Animal{}
	snake := Animal{}
	cow.InitCow()
	bird.InitBird()
	snake.InitSnake()
	for {
		fmt.Printf(">")
		inputReader := bufio.NewReader(os.Stdin)
		animalType, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		animals := strings.Trim(animalType, "\n")
		animalArray := strings.Split(animals, " ")
		// fmt.Println(len(animalArray))
		// fmt.Println(animalArray[2])
		var selectedAnimal string
		for i, v := range animalArray {
			if i == 0 {
				selectedAnimal = v
			}
			if i == 1 {
				if selectedAnimal == "cow" {
					if v == "Eat" {
						cow.Eat()
					} else if v == "Move" {
						cow.Move()
					} else if v == "Speak" {
						cow.Speak()
					}
				} else if selectedAnimal == "bird" {
					if v == "Eat" {
						bird.Eat()
					} else if v == "Move" {
						bird.Move()
					} else if v == "Speak" {
						bird.Speak()
					}
				} else if selectedAnimal == "snake" {
					if v == "Eat" {
						snake.Eat()
					} else if v == "Move" {
						snake.Move()
					} else if v == "Speak" {
						snake.Speak()
					}
				}
			}
		}
	}
}
