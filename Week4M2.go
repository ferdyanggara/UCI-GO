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

type AnimalInit struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (cow *AnimalInit) InitCow(animalName string) {
	cow.name = animalName
	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"
}

func (bird *AnimalInit) InitBird(animalName string) {
	bird.name = animalName
	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"
}

func (snake *AnimalInit) InitSnake(animalName string) {
	snake.name = animalName
	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"
}

func (a AnimalInit) Eat() {
	fmt.Println(a.food)
}
func (a AnimalInit) Move() {
	fmt.Println(a.locomotion)
}
func (a AnimalInit) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// FORMAT INPUT :
	// CREATED AN ANIMAL
	// 1st "newanimal"
	// 2nd name of the animal
	// 3rd type of the animal
	// CREATED A QUERY
	// 1st "query"
	// 2nd name of the animal
	// 3rd action (eat or move)
	animalList := make([]AnimalInit, 0, 5)
	for {
		fmt.Printf(">")
		inputReader := bufio.NewReader(os.Stdin)
		command, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cleanedCommand := strings.Trim(command, " \n")
		singleCommand := strings.Split(cleanedCommand, " ")
		// fmt.Println(singleCommand)
		if singleCommand[0] == "newanimal" {
			newAnimal := AnimalInit{}
			if singleCommand[2] == "cow" {
				newAnimal.InitCow(singleCommand[1])
			} else if singleCommand[2] == "bird" {
				newAnimal.InitBird(singleCommand[1])
			} else {
				newAnimal.InitSnake(singleCommand[1])
			}
			animalList = append(animalList, newAnimal)
			fmt.Println("Created it!")
		} else {
			for _, v := range animalList {
				if v.name == singleCommand[1] {
					if singleCommand[2] == "Eat" {
						v.Eat()
					} else if singleCommand[2] == "Move" {
						v.Move()
					} else if singleCommand[2] == "Speak" {
						v.Speak()
					}
				}
			}
		}
	}
}
