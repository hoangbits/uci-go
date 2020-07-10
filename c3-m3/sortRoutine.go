package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
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
	fmt.Println("Unsorted Main Slice:", sli)
	// each goroutine sort a subarray and print
	subSize := int(math.Ceil(float64(len(sli)) / 4.0))
	var wg sync.WaitGroup
	wg.Add(4)
	var mainSli = [][]int{}
	for i := 0; i < len(sli); i += subSize {
		var subSli = []int{}
		if i+subSize >= len(sli) {
			subSli = sli[i:]
		} else {
			subSli = sli[i : i+subSize]
		}
		mainSli = append(mainSli, subSli)
	}
	for _, subSli := range mainSli {
		go SortSubArray(subSli, &wg)
	}
	wg.Wait()
	// main goroutine sort data again and print
	sort.Ints(sli)
	fmt.Println("Sorted Main Slice:", sli)
}

// ReadN cm
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

// SortSubArray cm
func SortSubArray(sli []int, wg *sync.WaitGroup) {
	fmt.Println("Unsorted Subarray:", sli)
	sort.Ints(sli)
	fmt.Println("Sorted Subarray:", sli)
	wg.Done()
}
