/*
use `go run <program_name>` to start program
*/
// mandatory, main function always in main package
package main

import "fmt"

func main() {
	fmt.Println("test")
	fmt.Println("1 + 1 = ", add(1, 1))
	fmt.Println("4 - 2 = ", substract(4, 2))

	a := "AAA"
	b := "BBB"
	fmt.Println("Before swap string, a: ", a, " b: ", b)
	a, b = swap(a, b)
	fmt.Println("After swap string, a: ", a, "b: ", b)
	powerResult, halfResult := getPowerAndHalfOfValue(20)
	fmt.Println("Named return value: ", powerResult, " ", halfResult)
}

/* explicit param type declaration */
// explicit param declaration
func add(x int, y int) int {
	return x + y
}

// if param type are same, can omit the first one
func substract(x, y int) int {
	return x - y
}

// return > 1 result
func swap(a, b string) (string, string) {
	return b, a
}

// named return result
func getPowerAndHalfOfValue (value int) (powerResult int, halfResult float64) {
	powerResult = value * value
	halfResult = float64(value) / 2.0
	// omitted returned var here
	return
}
