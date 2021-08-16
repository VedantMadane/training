// Sample program to show how maps are not safe for concurrent use by default.
// The runtime will detect concurrent writes and panic.
package main

import (
	"fmt"
	"sync"
	"time"
)

// scores holds values incremented by multiple goroutines.
var scores = make(map[string]int)

func main() {
	var wg sync.WaitGroup
	var locker sync.Mutex

	wg.Add(2)
	start := time.Now()

	go func() {
		locker.Lock()
		for i := 0; i < 1000; i++ {

			scores["A"]++

		}
		locker.Unlock()
		wg.Done()
	}()

	end := time.Since(start).Seconds()

	go func() {
		for i := 0; i < 1000; i++ {
			locker.Lock()
			scores["B"]++
			locker.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final scores:", scores, "\n Time:", end)
}
