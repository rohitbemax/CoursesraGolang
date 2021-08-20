package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

func BubbleSort(iSlice []int) {
	length := len(iSlice)
	for i := 0; i < length-1; i++ {
		for j := 0; j < (length - i - 1); j++ {
			if iSlice[j] > iSlice[j+1] {
				Swap(iSlice, j)
			}
		}
	}
}

func Swap(tSlice []int, i int) {
	//We swap the elements
	tSlice[i], tSlice[i+1] = tSlice[i+1], tSlice[i]
}

func main() {
	var intSequence string
	fmt.Println("Enter a sequence of 10 integers (separated by space)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	intSequence = scanner.Text()
	sArr := strings.Split(intSequence, " ")
	iArr := make([]int, len(sArr))
	for idx, val := range sArr {
		t, _ := strconv.Atoi(val)
		iArr[idx] = t
	}

	BubbleSort(iArr)
	for _, val := range iArr {
		fmt.Printf("%d ", val)
	}
}
