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
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		var s1, s2 string
		fmt.Println("Please input request as 2 strings seperate by a space")
		fmt.Print(">")
		fmt.Scan(&s1, &s2)
		animal := animals[s1]
		switch s2 {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
}
