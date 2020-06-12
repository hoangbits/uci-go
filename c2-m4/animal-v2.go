package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var firstCmd, secondCmd, thirdCmd string
	var animals map[string]Animal
	animals = make(map[string]Animal)
	for {
		fmt.Println("Enter command (enter :q :q :q to quit!):")
		fmt.Print(">")
		fmt.Scan(&firstCmd, &secondCmd, &thirdCmd)
		/**
		case 1: newanimal-name-type(cow-bird-snake)
		case 2: query-name-information_request(eat-move-speak)
		*/
		if firstCmd == "newanimal" {
			if len(secondCmd) > 0 && len(thirdCmd) > 0 &&
				(thirdCmd == "cow" || thirdCmd == "bird" || thirdCmd == "snake") {
				var tmpAnimal Animal
				switch thirdCmd {
				case "cow":
					tmpAnimal = Cow{}
				case "bird":
					tmpAnimal = Bird{}
				case "snake":
					tmpAnimal = Snake{}
				}
				animals[secondCmd] = tmpAnimal
				fmt.Println("Created it")
			}
		} else if firstCmd == "query" {
			if len(secondCmd) > 0 && len(thirdCmd) > 0 {
				PrintAnimalByQuery(animals[secondCmd], thirdCmd)
			}
		} else if firstCmd == ":q :q :q" {
			fmt.Println("Bye!")
			break
		} else {
			fmt.Println("Not found!")
		}
	}
}

// PrintAnimalByQuery ,.
func PrintAnimalByQuery(animal Animal, method string) {
	switch ani := animal.(type) {
	case Cow:
		PrintByName(ani, method)
	case Bird:
		PrintByName(ani, method)
	case Snake:
		PrintByName(ani, method)
	}
}

// PrintByName ..
func PrintByName(ani Animal, method string) {
	switch method {
	case "eat":
		ani.Eat()
	case "move":
		ani.Move()
	case "speak":
		ani.Speak()
	}
}
