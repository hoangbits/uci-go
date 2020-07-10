package main

import (
	"fmt"
)

func main() {
	// https://stackoverflow.com/questions/39565055/read-a-set-of-integers-separated-by-space-in-golang
	fmt.Println("Please enter moret than 4 integers")
	var n int
	if m, err := fmt.Scan(&n); m != 1 {
		panic(err)
	}
	fmt.Println("Enter the integerss")
	sli := make([]int, n)
	ReadN(sli, 0, n)

}

// ReadN dfd
func ReadN(sli []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := fmt.Scan(&sli[i]); m != 1 {
		panic(err)
	}
	// 1,2,3, 4- 3,2,1, 0
	// e.g subsequent call when n=4 as initial value
	ReadN(sli, i+1, n-1)
}
