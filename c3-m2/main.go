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

/**
Race Condition
  Interleave: If I have 2 task and Order of execution between concurrent tasks
  is not known, it's not determined, meaning it's not deterministic
  It means these tasks can be interleaved in different ways.
  Interleave is happening at the machine code level rather than the source code

  Race condition is a problem where the outcome of the program depends on the interleaving
  And Interleaving is non-deterministic.
  It's determined by the operating system and the Go-runtime.
  So, the interleaving can change every time you run it.
**/
