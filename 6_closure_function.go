package main

import "fmt"

func main(){
	// closure function is a special kind of function that the inner function could capture state of outer function variable after scope end
	a := counter();
	fmt.Printf("calling a, a:%d\n", a())
	fmt.Printf("calling a, a:%d\n", a())
	fmt.Printf("calling a, a:%d\n", a())

	b := counter()
	fmt.Printf("calling b, b:%d\n", b())
	fmt.Printf("calling b, b:%d\n", b())
	fmt.Printf("calling b, b:%d\n", b())

	posAdder := adder()
	posAdder(2)
	posAdder(2)
	x := posAdder(2)
	fmt.Printf("x value now = %d\n", x)

	negAdder := adder()
	negAdder(-2)
	negAdder(-2)
	y := negAdder(-2)
	fmt.Printf("y value now = %d\n", y)

	fmt.Println("Print 10 fibronaci number")
	fib := fibonacci()
	for i:=0 ; i < 10 ; i++ {
		fmt.Println(fib())
	}

}

func counter() func() int{
	i := 0
	return func() int {
		i++
		return i
	}
}

// a closure function with param
// personally i feel the syntax look like JS callback
func adder() func(int)int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	first := 0
	second := 1

	return func() int {
		result := first + second
		first = second
		second = result
		return result
	}
}