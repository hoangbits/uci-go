package main

import (
	"fmt"
)

// Animal represent animal characteristic
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat smt
func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

// Move somewhere
func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

// Speak smt
func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	for {
		var s1, s2 string
		fmt.Println("Please input request as 2 strings seperate by a space")
		fmt.Print(">")
		fmt.Scan(&s1, &s2)
		switch s1 {
		case "cow":
			switch s2 {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		case "bird":
			switch s2 {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		case "snake":
			switch s2 {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}

		}
	}
}
