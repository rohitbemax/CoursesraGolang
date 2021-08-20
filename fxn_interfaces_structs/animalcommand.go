package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct{ food, locomotion, sound string }
type bird struct{ food, locomotion, sound string }
type snake struct{ food, locomotion, sound string }

//All the methods for cow
func (a cow) Eat() {
	fmt.Println(a.food)
}

func (a cow) Speak() {
	fmt.Println(a.sound)
}

func (a cow) Move() {
	fmt.Println(a.locomotion)
}

//All the methods for bird
func (a bird) Eat() {
	fmt.Println(a.food)
}

func (a bird) Speak() {
	fmt.Println(a.sound)
}

func (a bird) Move() {
	fmt.Println(a.locomotion)
}

//All the methods for snake
func (a snake) Eat() {
	fmt.Println(a.food)
}

func (a snake) Speak() {
	fmt.Println(a.sound)
}

func (a snake) Move() {
	fmt.Println(a.locomotion)
}

var animalsMap map[string]Animal

//newanimal animal_name type(cow/bird/snake)
func newAnimalCommand(animalName, animalType string) Animal {
	var animal Animal
	fmt.Println("AnimalType: ", animalType)
	switch animalType {
	case "cow":
		animal = cow{"grass", "walk", "moo"}

	case "bird":
		animal = bird{"worms", "fly", "peep"}

	case "snake":
		animal = snake{"mice", "slither", "hsss"}

	default:

	}
	return animal
}

func queryAnimalCommand(animalName, animalInformation string) {
	animal := animalsMap[animalName]
	if animal != nil {
		switch animalInformation {
		case "eat":
			animal.Eat()

		case "speak":
			animal.Speak()

		case "move":
			animal.Move()
		}
	} else {
		fmt.Printf("No animal [%s] exists\n", animalName)
	}
}

func scanAndProcessUserCommand() {

	var commandString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandString = scanner.Text()
	tokens := strings.Split(commandString, " ")

	if len(tokens) == 3 {
		switch tokens[0] {
		case "newanimal":
			if isValidNewAnimalCommand(tokens) {
				animalsMap[tokens[1]] = newAnimalCommand(tokens[1], tokens[2])
			} else {
				fmt.Printf("No animalType %s instantional detail found, cannot create animal object\n", tokens[2])
			}

		case "query":
			if isValidQueryAnimalCommand(tokens) {
				queryAnimalCommand(tokens[1], tokens[2])
			} else {
				fmt.Printf("No action %s found\n", tokens[2])
			}

		default:
			fmt.Println("Not a valid command")
		}

	} else if tokens[0] == "X" {
		os.Exit(0)
		fmt.Println("Will exit application")
	} else {
		fmt.Println("Should have 3 tokens <command> <param1> <param2>")
	}

}

//Type of animal can only be of type cow, bird or snake
func isValidNewAnimalCommand(tokens []string) bool {
	if (tokens[2] == "cow") || (tokens[2] == "bird") || (tokens[2] == "snake") {
		return true
	} else {
		return false
	}
}

//Animal can only have information related to eat, speak, move
func isValidQueryAnimalCommand(tokens []string) bool {
	if (tokens[2] == "eat") || (tokens[2] == "speak") || (tokens[2] == "move") {
		return true
	} else {
		return false
	}
}

func printUsage() {
	fmt.Println("*** Usage information (to exit type X and enter) ***")
	fmt.Println("### newanimal <animal_name> <animal_type either cow, bird or snake>")
	fmt.Println("### query <animal_name> <animal_information either eat, speak or move>")
	fmt.Println("### Example: 1. [>newanimal tweety bird] 2. [>query tweety speak]")
	fmt.Println("*** -------------------------------------------- ***")
	fmt.Println("")
}

func main() {
	animalsMap = make(map[string]Animal)

	for {
		printUsage()
		fmt.Print(">")
		scanAndProcessUserCommand()
	}
}
