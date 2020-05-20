package main

import "fmt"

func main() {
	var number int
	fmt.Printf("Input a floating number for truncating: ")
	num, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Scanned %d number(s) with Truncated value is: %d \n", int(num), int(number))
}
