package main

import (
	"fmt"
	"math"
)

/*
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

	fmtPrintln("pow of 2 =", powOrDefault10(2))
	fmtPrintln("pow of 10 =", powOrDefault10(10))
}

func powOrDefault10(num int) {
	min := 10
	// Go allow if statement to execute something first before evalation
	if p := math.Pow(num, 2); p < 10 {
		return min
	}

	return p
}
