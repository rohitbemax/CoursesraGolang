package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	NumCoroutines = 4
)

func sortArr(arr []int, tname string, ca chan []int){
	fmt.Println("Coroutine: ", tname, "is going to sort arr", arr)
	sort.Ints(arr)
	ca <- arr
}

func main() {

	printUsage()

	//Receive user input
	var commandString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandString = scanner.Text()
	tArr := strings.Split(commandString, " ")
	lengthArr := len(tArr)
	if lengthArr < 4 {
		fmt.Println("We need at least", NumCoroutines, "numbers to spread the work b/w", NumCoroutines,  "co-routines")
	}

	//Create an array of size equal to number of tokens
	intArr := make([]int, 0, lengthArr)

	//Convert the string tokens to integers and add to the integer array
	for _, val := range tArr {
		i, _ := strconv.Atoi(val)
		intArr = append(intArr, i)
	}

	//Create a channel which can take an array of integers equal to number of token provided by user
	ca := make(chan []int, lengthArr)

	//Create an array which can take sorted arrays from the co-routines
	collectorArr := make([]int, 0, lengthArr)

	//We will pass the data in chunks of i/NumCoroutines b/w NumCoroutines
	idx := 0
	start := 0
	end := 0
	//We will make NumCoroutines iterations and try to divide the array b/w NumCoroutines coroutines
	for iter:=1; iter<= NumCoroutines; iter++ {
		start = idx
		//During the 4th iteration if we have any extra elements we will assign it 4th co-routine
		if iter == NumCoroutines {
			end = idx+(lengthArr/ NumCoroutines) + lengthArr%NumCoroutines
		} else {
			end = idx + (lengthArr/ NumCoroutines)
		}
		go sortArr(intArr[start:end], "c" + strconv.Itoa(iter), ca)
		idx += lengthArr/ NumCoroutines
	}

	for {
		collectorArr = append(collectorArr, <- ca...)
		if len(collectorArr) == lengthArr {
			break
		}
	}

	sort.Ints(collectorArr)
	fmt.Println("Final sorted array in main: ", collectorArr)
}

func printUsage() {
	fmt.Println("### Usage (we need minimum", NumCoroutines,  "numbers to spread the data b/w ", NumCoroutines, "co-routines")
	fmt.Println("Enter numbers separated by space as delimiter and once you have entered numbers press enter")
	fmt.Println("Example: >1 2 0 -99 11 -45 55 17")
	fmt.Print("Input numbers(hit enter to finish)>")
}