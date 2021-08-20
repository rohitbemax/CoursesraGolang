package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	personMap := make(map[string]string)
	var name string
	var address string

	fmt.Println("Enter a name")
	fmt.Scan(&name)

	fmt.Println("Enter an address")
	fmt.Scan(&address)

	//Add the values to the map
	personMap["name"] = name
	personMap["address"] = address

	fmt.Println(personMap)

	jString, err := json.Marshal(personMap)
	if err == nil {
		fmt.Println(string(jString))
	}
}
