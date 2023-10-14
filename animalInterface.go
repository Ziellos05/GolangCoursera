package main

import "fmt"

func MainAnimalsInterface() {

	for {
		fmt.Println("Enter animal + action:")
		data := input([]string{}, nil)
		choiceInterface(data)
	}

}

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

type cow2 struct {
	footEaten        string
	locomotionMethod string
	spokenSound      string
}

type bird2 struct {
	footEaten        string
	locomotionMethod string
	spokenSound      string
}

type snake2 struct {
	footEaten        string
	locomotionMethod string
	spokenSound      string
}

func (c *cow2) Eat() {
	fmt.Println(c.footEaten)
}

func (c *cow2) Move() {
	fmt.Println(c.locomotionMethod)
}

func (c *bird2) Speak() {
	fmt.Println(c.spokenSound)
}

func (c *bird2) Eat() {
	fmt.Println(c.footEaten)
}

func (c *bird2) Move() {
	fmt.Println(c.locomotionMethod)
}

func (c *cow2) Speak() {
	fmt.Println(c.spokenSound)
}

func (c *snake2) Eat() {
	fmt.Println(c.footEaten)
}

func (c *snake2) Move() {
	fmt.Println(c.locomotionMethod)
}

func (c *snake2) Speak() {
	fmt.Println(c.spokenSound)
}

func choiceInterface(data []string) {

	var myAnimal AnimalInterface
	var myCow = cow2{"Grass", "walk", "moo"}
	var myBird = bird2{"worms", "fly", "peep"}
	var mySnake = snake2{"mice", "slither", "hsss"}

	switch animal := data[0]; animal {
	case "cow":
		myAnimal = &myCow
		switch action := data[1]; action {
		case "eat":
			myAnimal.Eat()
		case "move":
			myAnimal.Move()
		case "speak":
			myAnimal.Speak()
		default:
			fmt.Printf("YOUR ANIMAL DOESN'T DO IT")
		}
	case "bird":
		myAnimal = &myBird
		switch action := data[1]; action {
		case "eat":
			myAnimal.Eat()
		case "move":
			myAnimal.Move()
		case "speak":
			myAnimal.Speak()
		default:
			fmt.Printf("YOUR ANIMAL DOESN'T DO IT")
		}
	case "snake":
		myAnimal = &mySnake
		switch action := data[1]; action {
		case "eat":
			myAnimal.Eat()
		case "move":
			myAnimal.Move()
		case "speak":
			myAnimal.Speak()
		default:
			fmt.Printf("YOUR ANIMAL DOESN'T DO IT")
		}
	default:
		fmt.Printf("YOUR ANIMAL DOESN'T EXIST")
	}
}
