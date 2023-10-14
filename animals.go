package main

import (
	"fmt"
)

func MainAnimals() {

	for {
		fmt.Println("Enter animal + action:")
		data := input([]string{}, nil)
		choice(data)
	}

}

type Animal struct {
	footEaten        string
	locomotionMethod string
	spokenSound      string
}

var cow = Animal{footEaten: "grass", locomotionMethod: "walk", spokenSound: "moo"}
var bird = Animal{footEaten: "worms", locomotionMethod: "fly", spokenSound: "peep"}
var snake = Animal{footEaten: "mice", locomotionMethod: "slither", spokenSound: "hsss"}

func (a *Animal) Eat() {
	fmt.Println(a.footEaten)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotionMethod)
}

func (a *Animal) Speak() {
	fmt.Println(a.spokenSound)
}

func input(x []string, err error) []string {
	if err != nil {
		return x
	}
	var s string
	n, err := fmt.Scanf("%s", &s)
	if n == 1 {
		x = append(x, s)
	}
	return input(x, err)
}

func choice(data []string) {
	switch animal := data[0]; animal {
	case "cow":
		switch action := data[1]; action {
		case "eat":
			cow.Eat()
		case "move":
			cow.Move()
		case "speak":
			cow.Speak()
		default:
			fmt.Println("YOUR ANIMAL DOESN'T DO IT")
		}
	case "bird":
		switch action := data[1]; action {
		case "eat":
			bird.Eat()
		case "move":
			bird.Move()
		case "speak":
			bird.Speak()
		default:
			fmt.Println("YOUR ANIMAL DOESN'T DO IT")
		}
	case "snake":
		switch action := data[1]; action {
		case "eat":
			snake.Eat()
		case "move":
			snake.Move()
		case "speak":
			snake.Speak()
		default:
			fmt.Println("YOUR ANIMAL DOESN'T DO IT")
		}
	default:
		fmt.Println("YOUR ANIMAL DOESN'T EXIST")
	}
}
