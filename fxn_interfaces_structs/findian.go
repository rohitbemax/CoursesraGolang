package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a string:")
	inputReader := bufio.NewReader(os.Stdin)
	sentence, _ := inputReader.ReadString('\n')
	sentence = sentence[:len(sentence)-1] //Get rid of '\n'
	sentence = strings.ToLower(sentence)
	chars := []rune(sentence)

	if len(chars) < 3 {
		fmt.Print("Not found!")
	}

	switch {

	//Check if the first char is 'i'
	case chars[0] != 'i':
		fmt.Print("Not found!")

	//Check if the last character is 'n'
	case chars[len(chars)-1] != 'n':
		fmt.Print("Not found!")

	//Check if array has more than 2 characters 'i'-->chars[0] and 'n'-->chars[end]
	//and then check for 'a'-->chars[1-(end-1)]
	case len(chars) >= 3:
		foundA := false
		for i := 1; i <= len(chars)-2; i++ {
			if chars[i] == 'a' {
				fmt.Print("Found!")
				foundA = true
				break
			}
		}
		if !foundA {
			fmt.Println("Not Found!")
		}
	}
}
