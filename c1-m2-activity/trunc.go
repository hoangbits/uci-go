package main

import "fmt"

func main() {
	var number float64
	fmt.Printf("Input a floating number for truncating: ")
	num, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Scanned %d number(s) with Truncated value is: %v \n", int(num), int(number))
}
