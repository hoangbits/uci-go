package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* 1.gather input;
	sorting
	display
	*/
	var input string
	fmt.Printf("Input a list of numbers which seperates by a ',':\n")
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}

	sliStr := strings.Split(input, ",")
	var sliInt []int
	// var sliInt = []int{}
	// sliInt := []int{}

	for _, char := range sliStr {
		tempInt, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		sliInt = append(sliInt, tempInt)
	}
	fmt.Printf("Before sorting:%v\n", sliInt)
	sliInt = BubbleSort(sliInt)
	fmt.Printf("After sorting:%v\n", sliInt)
}

// Swap swap value between 2 andress
func Swap(pos *int, nextPos *int) {
	temp := *pos
	*pos = *nextPos
	*nextPos = temp
}

// BubbleSort implements sorting algorithm over a slice of
// integer
func BubbleSort(sli []int) []int {
	for i := range sli {
		for j := range sli[:len(sli)-1-i] {
			if sli[j] > sli[j+1] {
				Swap(&sli[j], &sli[j+1])
			}
		}
	}
	return sli
}
