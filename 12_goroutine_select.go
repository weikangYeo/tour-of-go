package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan int)

	// go routine call a anonymous function
	go func() {
		for i := 0; i < 10; i++ {
			// in a go routine process, it will pause when reading a not closed channel and it is empty
			// Go will unblock when the channel contain value
			// if it is called in main thread (main goroutine), and IF No other thread able to unblock it
			// then it will cause a "deadlock" panic
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// main thread continue to execute this, as the previous called was async call
	fibonacciWithSelect(c, quit)

	// the following code will cause deadlock panic
	//c := make(chan int)
	//fmt.Println("Waiting on channel c...")
	//<-c // The main goroutine blocks here forever
	//fmt.Println("This will never be printed.")

	testCh := make(chan int)
	go func() {
		time.Sleep(10 * time.Second)
		testCh <- 1
	}()
	fmt.Println("Try to read testCh value, where it has been sleep for 10 seconds")
	fmt.Println(<-testCh)
	fmt.Println("Value read and no deadlock panic")

	defaultSelectDemo()
}

func defaultSelectDemo() {
	fmt.Println("-------------")
	fmt.Println("Default Select Demo")
	fmt.Println("-------------")

	start := time.Now()
	// time.Tick() Returns a channel that will repeatedly receive values at a fixed frequency.
	tick := time.Tick(100 * time.Millisecond)
	// time.After()  Returns a channel that receives a value only once, after the duration has elapsed.
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			// default case, run when non of case valid
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	// infinite loop
	for {
		//  The select statement lets a goroutine wait on multiple communication operations.
		// thread blocked until one of cases run
		select {
		// this can run at first iteration
		case c <- x:
			fmt.Println("c <- X case")
			x, y = y, x+y
		// this can't be run at first 10 iteration in main, as `quit` always empty
		case <-quit:
			fmt.Println("quit case")
			return
		}
	}
}
