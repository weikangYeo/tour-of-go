package main

import "fmt"

func main() {
	arrayDemo()
	sliceDemo()
	sliceLiteral()
}

func sliceLiteral() {
	fmt.Println("==============================================")
	fmt.Println("Slice Literal")
	fmt.Println("==============================================")
	// slice literal is like an array without size
	// Go create an array then create a slice view on top of it
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice literal (without size): %#v\n", slice)
	fmt.Printf("Struct Slice (on demand struct type declaration and array\n")
	structSlice := []struct {
		Num int
		ok bool
	} {
		{1, true},
		{2, true},
		{3, false},
		{4, false},
	}
	fmt.Printf("slice literal (for struct): %#v\n", structSlice)


}

func sliceDemo() {
	fmt.Println("==============================================")
	fmt.Println("====================SLICE====================")
	fmt.Println("==============================================")

	// slide is a "view" of array
	// arr[start:end], which start is inclusive, end is exclusive (like substring)
	arr := [8]int{10, 20, 30, 40, 50, 60}
	var slice1 []int = arr[2:4]
	fmt.Printf("slice of 2 to 4: %v\n", slice1)
	slice2 := arr[:4]
	fmt.Printf("slice of 0 to 4 (implicit): %v\n", slice2)
	slice3 := arr[4:]
	fmt.Printf("slice of 4 to the end (implicit): %v\n", slice3)
	// note: arr[:] is also a valid slice, it mean a slice from top to the end

	// Since slice is view of array, change the value via slice will affect other slice output too
	// since... the underlying array value changed
	fmt.Println("Changing first ellement of slice1 (arr[2])")
	slice1[0] = 1000
	fmt.Println("Reprint all slice")
	fmt.Printf("slice of 2 to 4: %v\n", slice1)
	fmt.Printf("slice of 0 to 4 (implicit): %v\n", slice2)
	fmt.Printf("slice of 4 to the end (implicit): %v\n", slice3)
	fmt.Printf("Original Array: %v\n", arr)

}

func arrayDemo() {
	// array are fixed size
	// go arr size declaration is in front of type, which is contraverse than most of language
	var arr [2]int
	arr[0] = 100
	arr[1] = 200
	fmt.Printf("arr[0]=%d\n", arr[0])
	fmt.Printf("arr[1]=%d\n", arr[1])
	fmt.Println("Print the whole arr")
	fmt.Println(arr)

	// declare and init directly, the rest will just implicit to zero value
	// this mean Array literal
	arr2 := [6]int{1, 2, 3, 4}
	fmt.Printf("Array literal declararion : %#v\n", arr2)
	fmt.Println("Print the other arr")
	fmt.Println(arr2)
}
