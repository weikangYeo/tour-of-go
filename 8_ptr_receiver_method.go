package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Struct in Go could be value, it first copy value from original struct before run this "method"
func (v Vertex) AbsWithStateChanged() float64 {
	v.X = v.X * 10
	v.Y = v.Y * 10
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Pointer receiver type of func, it receive struct as pointer (addr), hence it would modify underlying value
func (v *Vertex) AbsWithStateChangedWithPtr() float64 {
	v.X = v.X * 10
	v.Y = v.Y * 10
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// pointer receiver, but func version
func scale (v *Vertex) {
	v.X = v.X * 10
	v.Y = v.Y * 10
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Before state change: %#v\n", v)
	fmt.Println(v.AbsWithStateChanged())
	// check result here, even the method try to change the state, origina struct value remain unchanged
	// because it cp value to method before execute it
	fmt.Printf("After state change with func receiver %#v\n", v)

	fmt.Printf("Before state change with ptr receiver: %#v\n", v)
	// suppose it is also need to pass address, as it is ptr receiver, Go make it as an convienient
	// actually it mean = (&v).AbsWithStateChangedWithPtr()
	fmt.Println(v.AbsWithStateChangedWithPtr())
	// notice value here, since it receive addr, original state get changed.
	fmt.Printf("After state change with ptr receiver %#v\n", v)
	// need to pass addresss
	scale(&v)
	fmt.Printf("After state change with ptr parameter function %#v\n", v)

	// reversely speaking, if it is a func receiver, Go make it as convienient as if the ptr variable is point to that struct
	ptr := &v
	fmt.Printf("Call PTR with func function %#v\n", *ptr)
	fmt.Println(ptr.AbsWithStateChanged()) // this is ok, and = to (*ptr).AbsWithStateChanged()

}
