package main

import "fmt"

func main() {

	var i interface{} = "hello"

	// this will get I and convert to string
	s := i.(string)
	fmt.Println(s)

	// type assertion will return 2 value, value and true false of that type
	s, ok := i.(string)
	fmt.Println(s, ok)

	// return 0 value if type not match
	f, ok := i.(float64)
	fmt.Println(f, ok)

	typeAssert(32)
	typeAssert(32.55)
	typeAssert("32.55")

}

func typeAssert(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Printf("%d is a int\n", value)
	case float64:
		fmt.Printf("%f is a float\n", value)
	default:
		fmt.Println("Unknown type found")
	}
}
