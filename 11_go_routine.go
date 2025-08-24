package main

import (
	"fmt"
	"time"
)

func main() {
	// go routine is as simple as "go function()"
	goRoutineSample()

	channelDemo()

	bufferedChannelDemo()
}

func goRoutineSample() {
	go test("Call by Go Routine")
	test("Normal Call")

	// We sleep for a second to allow the goroutine to finish.
	// Otherwise, the main function would exit immediately.
	time.Sleep(1 * time.Second)
}

func test(name string) {
	fmt.Printf("%s Prepare to Sleep\n", name)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s value : %d\n", name, i)
	}
}

// ch <- v // send v to channel 'ch'
// v:= <- ch // receive value from channel and assign to V

func sum(s []int, c chan int) {
	fmt.Println("Sum called")
	sum := 0
	for _, v := range s {
		sum += v
	}

	fmt.Println("Sending to Sum")
	// sending and receiving is blocking operation (synchronize)
	c <- sum // send sum to c
}

func channelDemo() {
	fmt.Println("-------------")
	fmt.Println("Channel DEMO")
	fmt.Println("-------------")
	s := []int{1, 23, 3, 5, 7, 8, 9, 55}

	// create a int channel
	c := make(chan int)

	// go routine call/ async call and store to channel
	fmt.Println("Before Go routine")
	go sum(s, c)
	go sum(s[:4], c)
	fmt.Println("After Go routine")
	x, y := <-c, <-c
	fmt.Printf("x: %d, y:%d\n", x, y)
}

func bufferedChannelDemo() {
	fmt.Println("-------------")
	fmt.Println("Buffered Channel DEMO")
	fmt.Println("-------------")
	//  Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

	// buffer size =2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// full, will hit "all goroutines are asleep - deadlock!"
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
