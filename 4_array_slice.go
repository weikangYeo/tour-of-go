package main

import (
	"fmt"
	"strings"
)

func main() {
	arrayDemo()
	sliceDemo()
	sliceLiteral()
	sliceLenthAndCapacity()
	makeSlice()
	sliceOfSlice()
	growSlice()
	rangeDemo()
}

func rangeDemo() {
	fmt.Println("==============================================")
	fmt.Println("Range()")
	fmt.Println("==============================================")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// if want index only
	for i := range pow {
		fmt.Printf("i = %d\n", i)
	}

	// if want value only
	for _ , v := range pow {
		fmt.Printf("v = %d\n", v)
	}
}

func growSlice() {
	fmt.Println("==============================================")
	fmt.Println("Grow Slice")
	fmt.Println("==============================================")
	var s []int
	printSliceLenCap(s)

	s = append(s, 1)
	printSliceLenCap(s)

}

func sliceOfSlice() {
	fmt.Println("==============================================")
	fmt.Println("Slice of Slice")
	fmt.Println("==============================================")
	twoDimSlice := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	twoDimSlice[0][0] = "X"
	twoDimSlice[1][1] = "O"

	for i := 0; i < len(twoDimSlice); i++ {
		fmt.Printf("%s\n", strings.Join(twoDimSlice[i], " "))
	}
}

func makeSlice() {
	fmt.Println("==============================================")
	fmt.Println("make() slice")
	fmt.Println("==============================================")
	// slice can also be created via make() cmd
	//  The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)
	printSliceLenCap(a)

	// make(T, len, cap)
	b := make([]int, 0, 5)
	printSliceLenCap(b)
	b = b[:cap(b)]
	printSliceLenCap(b)
	b = b[2:]
	printSliceLenCap(b)

	// Note on what just happened:
	/*
			1. `b := make([]int, 0, 5)`
		       * This command does two things:
		           1. It allocates a new, hidden array in memory with a size of 5. Because it's an int array, Go initializes all 5 elements to their zero value, which is 0. So, in
		              memory, you have [0, 0, 0, 0, 0].
		           2. It creates a slice b that points to this underlying array.
		       * However, you specified a length of 0. This means the slice b is a "view" of the array that currently shows none of its elements.
		       * At this point: len=0, cap=5, b=[]

		   2. `b = b[:cap(b)]`
		       * This is the crucial step. This operation creates a new slice that shares the same underlying array.
		       * You are saying, "create a new view of the original array that starts at the beginning (:), and extends to the full capacity (cap(b), which is 5)."
		       * This new slice now has a length of 5. Since it's viewing the underlying array that was already initialized to [0, 0, 0, 0, 0], that's the content you see.
		       * At this point: len=5, cap=5, b=[0 0 0 0 0]

		   So, the underlying array was never empty; it was just that your initial slice b had a length of 0, preventing you from seeing the zeroed values that make had already allocated.
		   The second operation simply expanded the slice's "view" to include the entire capacity of that array.
	*/
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
		ok  bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{4, false},
	}
	fmt.Printf("slice literal (for struct): %#v\n", structSlice)
}

func sliceLenthAndCapacity() {
	/*
			* `cap` (Capacity): This determines the size of the memory allocation for the underlying array. It's the maximum number of elements the slice can hold before Go needs to
		     allocate a new, larger array and copy the old data over.
			* `len` (Length): This determines how many of those allocated elements are initially accessible through the slice. It's the initial "view" into the array.
	*/

	fmt.Println("==============================================")
	fmt.Println("Slice Len & Cap")
	fmt.Println("==============================================")

	// slice literal, create an array & slice point to it
	// NOTE: Go is like C and any other language, array is sit side by side in memory
	// and when I point to an array, I am pointing to the FIRST element of array.
	s := []int{1, 2, 3, 4, 5, 6}

	// len is how many element slice VIEW the array
	// cap is size of array it point to
	printSliceLenCap(s)

	// set the slice to 0 elment,
	s = s[:0]
	// notice the print, underlying array is still the same,
	// slice still point to that array, just that the 'view' is diff
	printSliceLenCap(s)

	// Extend the length
	s = s[:4]
	printSliceLenCap(s)

	// drop first 2 value
	// what this happen is, we create a NEW SLICE that
	// point to same underlying array, but start at INDEX 2
	// so now from Slice POV, the array it point to, the first index is INDEX 2 of the original array
	// mean now [3,4,5,6] is the array this slice point to.
	// hence cap of slice = 4,
	// previously `s` has a view of [1,2,3,4] (after the s[:4]) from array [1,2,3,4,5,6]
	// now since s point from index 2, so the view become [3,4] from array [3,4,5,6]
	s = s[2:]
	printSliceLenCap(s)

	// create a new array and slice pts to it
	// notice the len is changed
	s = []int{1}
	printSliceLenCap(s)

	// Zero value slice
	s = []int{}
	printSliceLenCap(s)
}

func printSliceLenCap(s []int) {
	// len is how many element the Slice view it
	// capacity is how many element of the original array that Slice point to
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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

	// Go Array is VALUE instead of pointer like C
	fmt.Println("copy value of arr and change to another value")
	cpyArr := arr
	cpyArr[0] = 10000
	cpyArr[1] = 20000
	fmt.Printf("arr: %#v\n", arr)
	fmt.Printf("cpyArr: %#v\n", cpyArr)

}
