package main

import "fmt"

// global variable
// global variable cannot use `:=`
var isC, isPython, isJava bool

func main() {
	// default all variable are 0 value.
	var i int
	fmt.Println("var: ", i, isC, isPython, isJava)

	// var declaration with initialization
	// type can be omitted in this case
	var a, b, c = 1, 2, "test"
	fmt.Println("var: ", a, b, c)

	// declaration with initialization 2
	x, y := 100, 200
	fmt.Println("var: ", x, y)

	// type converstion
	num := -42
	// Go doesnt have implicit type converstion
	floatNum := float64(num) + 0.2
	unsignedNum := uint(num)
	fmt.Println("type conversion: ", num, floatNum, unsignedNum)

	// constant, cannot use `:=`
	// Capitalize = Public, all lowercase = private
	const SomeConst = "World"
	const SomeConst2 = 1000
	const SomeConst3 = 2.1413

	fmt.Println("const: ", SomeConst, SomeConst2, SomeConst3)

}
