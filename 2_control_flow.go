package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/*
If, else, switch, loop (for), defer (finally).
Go control flow no need "()"
*/
func main() {
	// usual for loop
	for i := 0; i < 5; i++ {
		fmt.Print("idx:", i, " ")
	}

	fmt.Println("")
	// init and increment could be omitted, it is a while loop
	sum := 1
	for sum <= 10 {
		sum++
	}

	// if else
	if sum >= 10 {
		fmt.Println("Sum = ", sum)
	}

	fmt.Printf("pow of 2 =%f", powOrDefault10(2))
	fmt.Printf("pow of 10 =%f", powOrDefault10(10))

	// switch
	// go switch only run in 1 case, so no break required
	switch os:= runtime.GOOS; os {
	case "darwin":
		fmt.Println("Darwin")
	case "linux":
		fmt.Println("Linux")
	default :
		fmt.Printf("Default case: %s", os)
	}

	// switch var scope ended after switch block
	//	fmt.Println("verify switch variable scope, %s", os)

	// alternative way of multiple if is using "switch true"
	currentTime := time.Now()
	switch {
	case currentTime.Hour() < 12:
		fmt.Println("Good Morning")
	case currentTime.Hour() < 18:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}

	fmt.Println()
	fmt.Println("Before function")
	deferTest()
	fmt.Println("After function")


	fmt.Println()
	fmt.Println("Before loop")

	// when multiple defer is called, it is pushed to a stack
	for i:=1 ; i < 10; i++ {
		// defer would run AFTER function returned, not after a block closed.
		defer fmt.Printf("I am a deferred num: %d\n", i)
	}
	fmt.Println("After loop")

}

func deferTest(){
	// defer is like "finally" in java, it run after the function return. (but before the next statement from caller run)
	fmt.Println("start of function")
	defer fmt.Println("I am a deferred called")
	fmt.Println("End of function statement");
}

func powOrDefault10(num int) float64{
	min := float64(10)
	// Go allow if statement to execute something first before evalation
	if p := math.Pow(float64(num), 2); p > 10 {
		return p
	}

	// var p scope is ended after if blocl
	//	return p
	return min
}
