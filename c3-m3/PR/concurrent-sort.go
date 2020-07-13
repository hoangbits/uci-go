package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var sortedNumbers = make(chan []int, 4)

	numbers := getNumbers()

	if len(numbers) < 4 {
		fmt.Printf("More numbers are needed")
		return
	}

	s := float32(len(numbers)) / 4
	b1, b2, b3 := int(s*1), int(s*2), int(s*3)

	go sortSlice(numbers[:b1], sortedNumbers)
	go sortSlice(numbers[b1:b2], sortedNumbers)
	go sortSlice(numbers[b2:b3], sortedNumbers)
	go sortSlice(numbers[b3:], sortedNumbers)

	wg.Add(4)

	wg.Wait()

	var result []int
	for i := 0; i < 4; i++ {
		sorted := <-sortedNumbers
		result = append(result, sorted...)
	}

	sort.Ints(result)
	fmt.Printf("\n # Final Result: %+v\n", result)

}

func getNumbers() []int {
	var numbers []int

	fmt.Println("Insert at least 4 numbers")
	fmt.Println("To random choose specify 2, lengh and maximum positive value")

	inputReader := bufio.NewReader(os.Stdin)

	input, _ := inputReader.ReadString('\n')

	rawNumbers := strings.Split(input, ",")

	for _, rawNumber := range rawNumbers {
		rawNumber = strings.TrimSpace(rawNumber)
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			continue
		}

		numbers = append(numbers, number)
	}

	if len(numbers) == 2 {
		length, limit := numbers[0], numbers[1]
		var randomNumbers []int

		for i := 0; i < length; i++ {
			randomNumbers = append(randomNumbers, rand.Intn(limit))
		}
		return randomNumbers
	}

	return numbers
}

func sortSlice(numbers []int, sorted chan<- []int) {
	fmt.Printf("%+v\n", numbers)
	sort.Ints(numbers)
	sorted <- numbers
	wg.Done()
}
