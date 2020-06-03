package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sli := []int{}
	for i := 0; true; i++ {
		fmt.Print("Enter a number/x to exit: ")
		var input string
		fmt.Scan(&input)
		el, error := strconv.Atoi(input)

		if error != nil {
			fmt.Println("Program exited")
			break
		}

		sli = append(sli, el)
		fmt.Println("Appended:", sli)
		sort.Ints(sli)
		fmt.Println("Sorted:", sli)

	}

}
