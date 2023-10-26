// Here is an easy example where the program is running multiple goroutines using an iterator, and the race conditions happen
// when multiple goroutines read the value of the counter and increase it, some routines read the same value and update the same value,
// this actually causes a false addition and in the end can yields a wrong result

// Note that the use of the wg variable allows the program to know that a new goroutine has started and when it finished,
// before continuing to the next step of the program, this let us finish all the goroutines before printing the result,
// BUT you can erase the wg variable and the program still works, the big difference is that this time the program will print
// a more wrong number, because it's printing BEFORE finishing all the tasks running!

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
