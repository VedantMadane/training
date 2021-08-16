// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

// Add imports.
import (
	"fmt"
	"runtime"
	"sync"
	//	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declare a wait group and set the count to two.
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		for i := 100; i >= 0; i-- {
			fmt.Println("--", i)
		}

		// Tell main we are done.
		waitgroup.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts up from 0 to 100 and
		// display each value.
		for i := 0; i <= 100; i++ {
			fmt.Println("++", i)
		}

		// Tell main we are done.
		waitgroup.Done()
	}()
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	waitgroup.Wait()
	// Display "Terminating Program".
	fmt.Println("Terminating program")
}
