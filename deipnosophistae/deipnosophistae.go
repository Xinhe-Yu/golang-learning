package main

import (
	"fmt"
	"sync"
	"time"
)

const philosopherCount = 5
const maxConcurrentEaters = 2

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id              int
	leftCS, rightCS *Chopstick
}

var wg sync.WaitGroup

func host(permit chan struct{}) {
	for {
		if len(permit) < maxConcurrentEaters {
			permit <- struct{}{}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (p *Philosopher) eat(permit chan struct{}) {
	for i := 0; i < 3; i++ {
		<-permit

		p.rightCS.Lock()
		p.leftCS.Lock()

		fmt.Printf("starting to eat %d\n", p.id+1)
		time.Sleep(time.Second)
		fmt.Printf("finishing eating %d\n", p.id+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		wg.Done()
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	chopsticks := make([]*Chopstick, philosopherCount)
	for i := 0; i < philosopherCount; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, philosopherCount)
	for i := 0; i < philosopherCount; i++ {
		philosophers[i] = &Philosopher{
			id:      i,
			leftCS:  chopsticks[i],
			rightCS: chopsticks[(i+1)%philosopherCount],
		}
	}

	permit := make(chan struct{}, maxConcurrentEaters)

	wg.Add(philosopherCount * 3)

	go host(permit)

	for _, philosopher := range philosophers {
		go philosopher.eat(permit)
	}

	wg.Wait()
}
