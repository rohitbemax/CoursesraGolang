package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Println("Enter the filename:")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)

	//Create a slice of type name
	sPersons := make([]name, 0, 10)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	//Read one line at a time from the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		//Split the string into nameArr[0]=fname & nameArr[1]=lname
		nameArr := strings.Split(scanner.Text(), " ")

		//Create a new name struct object and add to the slice
		sPersons = append(sPersons, name{fname: nameArr[0], lname: nameArr[1]})
	}

	file.Close()

	//Iterate the slice and print the fname and lname
	for _, v := range sPersons {
		fmt.Println(v)
	}
}
