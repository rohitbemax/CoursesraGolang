package main

import (
	"fmt"
	"sync"
	"time"
)
/*Implement the dining philosopher’s problem with the following constraints/modifications.
1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher. */

const (
	NumPhilos = 5
	TimesToEat = 3
)

type chopS struct { sync.Mutex }

type philo struct {
	id int
	leftCS, rightCS *chopS
}

func philoEat(p philo, ch chan bool, wa *sync.WaitGroup) {

	//Each philosopher can eat only three time
	for i:=0; i<TimesToEat; i++ {
		<- ch //If the host allows philosopher to eat (any two at a time)

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.id)

		//We will spend some time eating (simulating)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("finishing eating", p.id)
		p.rightCS.Unlock()
		p.leftCS.Unlock()

		//We will spend some time before we eat again (simulating via sleep)
		time.Sleep(20 * time.Millisecond)
	}

	//Indicate that you have finished eating 3 times
	wa.Done()
}

func host(eat chan bool) {
	for {
		//Since it's a buffered channel of size 2 only 2 go routines can run at a time concurrently
		eat <- true
	}
}

func main() {

	philosophers := make([]philo, 0, 5)
	chopSticks := make([]chopS, 0, 5)

	//We will use buffered channels for synchronisation
	eat := make(chan bool, 2)

	var wa sync.WaitGroup
	wa.Add(5)

	//Create 5 chops sticks
	for i:=0; i<cap(chopSticks); i++ {
		var chop chopS
		chopSticks = append(chopSticks, chop)
	}

	//Create 5 philosophers
	for i:=0; i< cap(philosophers); i++ {
		philosophers = append(philosophers, philo{id: i, leftCS: &chopSticks[i], rightCS: &chopSticks[(i+1)%5]})
	}

	for i:=0; i< NumPhilos; i++ {
		go philoEat(philosophers[i], eat, &wa)
	}

	go host(eat)

	//We will wait for all the 5 philosophers to finish eating
	wa.Wait()

	//fmt.Println("Number of philosopher:", len(philosophers), cap(philosophers))
	fmt.Println("All philosophers have eaten 3 times, exiting application")
}