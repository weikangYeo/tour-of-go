package main

import "fmt"

// struct, similar like C but without `,`
// private struct, only can be used in same package
type car struct {
	Make  string
	Model string
	// private field, only can access within same package
	serialNum string
}

type Vertex struct {
	X int
	Y int
}

func main() {
	pointerDemo()
	basicStruct()
	strucLiteral()
}

func strucLiteral(){
	fmt.Println("Struct Literal with Defined Type")
	a := Vertex{1, 2} // implicit X = 1, Y =2
	b := Vertex{Y: 1} // implicit X = 0
	c := Vertex{} // implicit X = 0 , Y = 0
	ptr := Vertex{3,4 } // create a struct and pointed by ptr

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(ptr)

	fmt.Println("Struct Literal with OnDemand Type")
	testOnDemandStruct := struct {
		someField string
		someOtherField string
	}{
		"test1",
		"test2",
	}

	fmt.Printf("On demand struct declaration: %#v\n", testOnDemandStruct)
}

func basicStruct(){
	// init a struct
	myCar  := car {"honda", "city", "12345"}
	// it wont show field name in Print
	fmt.Println("Car:")
	fmt.Println(myCar)

	// access struct field via "."
	myCar.Make = "Toyota"
	myCar.Model = "Vios"
	// private field can be access within same package
	myCar.serialNum = "123"
	fmt.Println("Car after change:")
	fmt.Println(myCar)

	// struct can be pointed by struct pointer
	carPtr := &myCar
	fmt.Println("Print car via ptr")
	fmt.Println(*carPtr)
	fmt.Println("Change field via ptr")
	// not like C, when accessing ptr's references property, dont have to do *
	// in C  need *carPtr->Model = "something else"
	// C only can use car.Model if car is direct struct
	// alternatively *(carPtr).Model
	// Go will just auto deferencing, regardless it is struct ptr or direct struct
	carPtr.Model = "Corrolla"
	fmt.Println("Print car after changed")
	fmt.Println(*carPtr)
}

func pointerDemo() {
	// pointer, like C
	a, b := 1, 2

	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)

	var numPtr *int
	numPtr = &a
	fmt.Printf("a, thru ptr: %d\n", *numPtr)
	fmt.Printf("ptr addr = %p\n", numPtr)

	*numPtr = 20
	fmt.Println("After ptr change value")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)

	fmt.Println("PTR can be reassinged")
	numPtr = &b
	fmt.Printf("*ptr now = b = %d\n", *numPtr)
	fmt.Printf("ptr addr = %p\n", numPtr)

	// Go pointer can't do arithmetric, like ptr++
}
