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

	rangeAndCloseDemo()
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

func rangeAndCloseDemo() {
	fmt.Println("-------------")
	fmt.Println("Range and Close Channel DEMO")
	fmt.Println("-------------")

	// ch can return a flag indicate whether there is next val to get or not
	// we can use `range` to iterate available value
	// v, ok := <-ch
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	c2 := make(chan int, 10)
	c2 <- 1
	c2 <- 2
	// if a channel without close, range reqad will cause error
	//fatal error: all goroutines are asleep - deadlock!
	//goroutine 1 [chan receive]
	close(c2)

	fmt.Println("try to range read from channel")
	for i := range c2 {
		fmt.Println(i)
	}
	fmt.Println("Range read ended")

	c3 := make(chan int, 10)
	c3 <- 1
	c3 <- 2

	<-c3
	//<-c3

	// if channel is already empty, the following code fails
	nextVal, haveValue := <-c3
	if haveValue {
		fmt.Println("hasValue")
		fmt.Println(nextVal)
	} else {
		fmt.Println("c3 channel is empty")
	}

	// the following would not work, as value from channel c2 already read from range previously
	//fmt.Println("try to read directly from channel")
	//fmt.Println(<-c2)
	//fmt.Println(<-c2)

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// read is suppose to be closed by sender, not reader
	close(c)
}
