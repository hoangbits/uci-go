package main

import (
	"fmt"
	"strings"
)

/*
input a string and check
i character at first
n at last
a in between
*/

func main() {
	for {
		var findString string
		fmt.Printf("Please input a string(input ':q' to quit): ")
		num, err := fmt.Scan(&findString)
		if err != nil {
			fmt.Printf(err.Error() + string(num))
		}

		if findian(findString) {
			fmt.Printf("Found!\n")
		} else {
			if findString == ":q" {
				fmt.Printf("Bye!\n")
				break
			}
			fmt.Printf("Not Found!\n")
		}
	}

}

func findian(s string) bool {
	if (strings.HasPrefix(s, "i") || strings.HasPrefix(s, "I")) &&
		strings.ContainsAny(s, "aA") &&
		strings.HasSuffix(s, "n") || strings.HasSuffix(s, "N") {
		return true
	}
	return false
}
