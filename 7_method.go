package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// a function, who supplied to a type (struct), function name (param) return type
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// or we can just as a regular function, basically it is the same
func AbsUsual(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func receiver is apply to non struct type also
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)}

func main() {
	// go don't have class (like java) hence dont have "class method"
	// the way defined it, i feel like is accept a type, and with any type can call this function
	// feel like it is kind of look like Kotlin now
	vertex := Vertex{2.5, 6.6}
	fmt.Println(vertex.Abs())
	fmt.Println(AbsUsual(vertex))

	myFloat := MyFloat(1.23)
	fmt.Println(myFloat.Abs())

}

