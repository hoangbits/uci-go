package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *ChopS
	name string
	nbrTimeEat int
}

func main() {

	var wg sync.WaitGroup
	wg.Add(15)
	c := make(chan *Philosopher,2)

	Csticks := make([] *ChopS, 5)
	for i := 0; i<5; i++ {
		Csticks[i] = new(ChopS)
	}

	philos := make([] *Philosopher, 5)
	for i := 0; i<5; i++ {
		philos[i] = &Philosopher{Csticks[(i+1)%5],Csticks[i], strconv.Itoa(i+1),0}
	}


	go host(c)
	for i := 0; i<5; i++ {
		go philos[i].eat(c,&wg)
	}

	wg.Wait()

}

func (p Philosopher) eat(channel chan *Philosopher, wait *sync.WaitGroup) {
	for {
		channel <- &p
		if (p.nbrTimeEat<3) {
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Printf("starting to eat %v\n", p.name)
			p.nbrTimeEat = p.nbrTimeEat+1
			//Eating take time
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("finishing eating %v\n", p.name)

			p.rightCS.Unlock()
			p.leftCS.Unlock()
			wait.Done()
		}
	}
}

func host(c chan *Philosopher) {
	for {
		if  (len(c)==2) {
			<- c
			<- c
		}
	}
}
