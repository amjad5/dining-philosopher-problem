package main

/*	Implement the dining philosopher’s problem with the following constraints/modifications.
	There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
	The host allows no more than 2 philosophers to eat concurrently.
	Each philosopher is numbered, 1 through 5.
	When a philosopher starts eating (after it has obtained necessary locks)
	it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
	When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself,
	where <number> is the number of the philosopher.
*/

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	timesEaten      int
	id              int
}

var hostChanel = make(chan bool, 2)
var philWaitGroup sync.WaitGroup

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		hostChanel <- true
		fmt.Println("starting to eat ", p.id)
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("finishing eating ", p.id)
		p.timesEaten += 1
		time.Sleep(200 * time.Millisecond)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		<-hostChanel
	}
	philWaitGroup.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], 1, i + 1}
	}

	for i := 0; i < 5; i++ {
		philWaitGroup.Add(1)
		go philos[i].eat()
	}

	philWaitGroup.Wait()
}
