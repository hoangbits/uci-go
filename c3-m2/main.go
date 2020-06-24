package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	a := 1
	go foo(&a, &wg)
	go bar(&a, &wg)
	wg.Wait()
}

func foo(a *int, wg *sync.WaitGroup) {
	*a = *a + 2
	fmt.Println(*a)
	wg.Done()
}
func bar(a *int, wg *sync.WaitGroup) {
	*a = *a + 6
	fmt.Println(*a)
	wg.Done()
}
