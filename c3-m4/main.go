package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ChopS cm
type ChopS struct {
	sync.Mutex
}

// Philo cm
type Philo struct {
	//6,Each philosopher is numbered, 1 through 5.
	name            int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(semanphore chan int, wg *sync.WaitGroup) {
	// 2.Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		// 7.When a philosopher starts eating
		fmt.Println("starting to eat ", p.name)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		// 8.When a philosopher finishes eating
		fmt.Println("finishing eating ", p.name)
	}
	<-semanphore // removes an int from semanphore, allowing another to proceed
	wg.Done()
}

// MaxConcurrency Maximum number of gorountine run concurrency
const MaxConcurrency = 2

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {

		// 3.The philosophers pick up the chopsticks in any order,
		// not lowest-numbered first (which we did in lecture).
		// philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]} // lecture
		if RandBool() {
			philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
		}
		philos[i] = &Philo{i + 1, CSticks[(i+1)%5], CSticks[i]}
	}
	semanphore := make(chan int, MaxConcurrency)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// 5.The host allows no more than 2 philosophers to eat concurrently.
		semanphore <- 1 // will block if there is MaxConccurrency in semanphore
		// 4.In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		go philos[i].eat(semanphore, &wg)
	}
	wg.Wait()
}

// RandBool This function returns a random boolean value based on the current time
func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return true
	}
	return false
}
