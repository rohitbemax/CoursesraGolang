package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Noise() {
	fmt.Println(animal.noise)
}

//We can have some Ploymophism here since all are Animals
func (animal *Animal) Action(action string) {
	switch action {
	case "eat":
		animal.Eat()

	case "move":
		animal.Move()

	case "speak":
		animal.Noise()
	}
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var animalName string
	var animalAction string

	printUsageInfo()

	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		strVal := scanner.Text()
		tokens := strings.Split(strVal, " ")
		if len(tokens) < 2 {
			printUsageSyntax()
			continue
		}

		animalName = tokens[0]
		animalAction = tokens[1]

		switch animalName {
		case "cow":
			cow.Action(animalAction)

		case "bird":
			bird.Action(animalAction)

		case "snake":
			snake.Action(animalAction)
		}
	}
}

func printUsageInfo() {
	fmt.Println("Enter animal(cow, bird, snake) followed by action request(eat, move, speak)")
	fmt.Println("Example1: > cow eat")
	fmt.Println("Example2: > snake speak")
}

func printUsageSyntax() {
	fmt.Println("You must enter two words > animal_name animal_action")
}
