package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//1. Integer slice of size 3
	var iSlice = make([]int, 0, 3)

	//In a loop ask user to enter integers till they enter "X"
	temp := ""
	for i := 0; ; i++ {
		fmt.Print("Enter an interger (enter X to terminate):")
		fmt.Scan(&temp)

		val, err := strconv.Atoi(temp)
		if err == nil {
			iSlice = append(iSlice, val)
			sort.Ints(iSlice)
			printSlice(iSlice)

		} else if temp == "X" {
			break
		}
	}
}

func printSlice(s []int) {
	fmt.Println(s)
}
