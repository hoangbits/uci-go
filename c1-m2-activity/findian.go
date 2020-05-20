package main

import (
	"fmt"
	"strings"
)

func main() {
	var findString string
	fmt.Printf("Please input a string: ")
	num, err := fmt.Scan(&findString)
	if err != nil {
		fmt.Printf(err.Error() + string(num))
	}

	if findian(findString) {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}

func findian(s string) bool {
	if (strings.HasPrefix(s, "i") || strings.HasPrefix(s, "I")) && strings.ContainsAny(s, "aA") && strings.ContainsAny(string(s[len(s)-1]), "n") {
		return true
	}
	return false
}
