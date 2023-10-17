package main

import (
	"fmt"
	"sync"
)

const numPhilosophers = 5
const numMeals = 3

var (
	chopsticks  [numPhilosophers]sync.Mutex
	permissions = make(chan struct{}, numPhilosophers-1)
)

func philosopher(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for meal := 0; meal < numMeals; meal++ {
		think(id)

		pickupChopsticks(id)
		eat(id)
		putdownChopsticks(id)
	}
}

func think(id int) {
	fmt.Printf("Philosopher %d is thinking\n", id)
}

func pickupChopsticks(id int) {
	fmt.Printf("Philosopher %d is trying to pick up chopsticks\n", id)

	chopsticks[id].Lock()
	chopsticks[(id+1)%numPhilosophers].Lock()

	fmt.Printf("Philosopher %d picked up chopsticks\n", id)
}

func putdownChopsticks(id int) {
	fmt.Printf("Philosopher %d is putting down chopsticks\n", id)

	chopsticks[id].Unlock()
	chopsticks[(id+1)%numPhilosophers].Unlock()
}

func eat(id int) {
	fmt.Printf("Philosopher %d is starting to eat\n", id)
	permissions <- struct{}{}
	defer func() { <-permissions }()
	fmt.Printf("Philosopher %d is finishing eating\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosopher(i, &wg)
	}

	wg.Wait()
}
