package main

import (
	"fmt"
	"sort"
	"unicode"
)

func main() {
	var sli []int
	for {
		var inputValue rune
		fmt.Println("Please enter an integer: ")
		// duplicated line: https://stackoverflow.com/questions/57734634/why-i-get-a-duplicate-output-when-the-loop-if-there-is-a-scanf
		num, err := fmt.Scanf("%c\n", &inputValue)
		if err != nil {
			fmt.Print(err.Error() + "\n")
			continue
		}

		if unicode.IsNumber(inputValue) {
			// get integer out of ASCII: https://stackoverflow.com/questions/21322173/convert-rune-to-int
			number := int(inputValue - '0')
			sli = append(sli, number)
			sort.SliceStable(sli, func(i, j int) bool { return sli[i] < sli[j] })

			fmt.Printf("Scanned: %d len=%d cap=%d %v \n\n", num, len(sli), cap(sli), sli)

		} else if unicode.IsLetter(inputValue) && inputValue == rune('w') {
			fmt.Print("bye!\n")
			break
		}

	}
}
