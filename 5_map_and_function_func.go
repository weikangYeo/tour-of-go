package main

import (
    "fmt"
    "strings"
)

type Vertex struct {
    Lat, Long float64
}

func main() {
    mapDeclaration()
    mapMutation()
    fmt.Println(WordCount("I am learning Go!"))
    useFunctionRunFunction()
}

func WordCount(s string) map[string]int {
    result := make(map[string]int)
    words := strings.Fields(s)

    for _, w := range words {
        _, isExists := result[w]
        if isExists {
            result[w]++
        } else {
            result[w] = 1
        }
    }
    return result
}

func mapMutation() {
    // always declare instance with "make"
    intByString := make(map[string]int)

    // add element
    intByString["TEST"] = 100
    // retrieve element
    fmt.Println(intByString["TEST"])
    // delete element
    delete(intByString, "TEST")
    // retrieve element with isExists
    value, ok := intByString["TEST"]
    fmt.Printf("Map value : %d, is exists: %t", value, ok)

}

func mapDeclaration() {
    // Go map is like java map, a key value map, just the way it define would be slightly different
    // define a map type
    var vertexByKey map[string]Vertex
    // create a map instance
    vertexByKey = make(map[string]Vertex)
    vertexByKey["TEST_KEY"] = Vertex{1.0, 2.0}
    fmt.Println(vertexByKey["TEST_KEY"])
    fmt.Println(vertexByKey)

    // its work too
    intByString := make(map[string]int)
    intByString["TEST"] = 100

    fmt.Println(intByString["TEST"])
    fmt.Println(intByString)

    // map literal
    // map literal with Type specific
    var mLiteral = map[string]Vertex{
        "TEST_KEY_1": Vertex{10.5, 20.5},
        "TEST_KEY_2": Vertex{20.5, 40.5},
    }

    fmt.Println(mLiteral)

    // omit type name (preferred)
    var mLiteral2 = map[string]Vertex{
        "TEST_KEY_1": {10.5, 20.5},
        "TEST_KEY_2": {20.5, 40.5},
    }

    fmt.Println(mLiteral2)
}

func useFunctionRunFunction() {
    fmt.Println("==============================================")
	fmt.Println("use Function call function")
	fmt.Println("==============================================")
	// function is value in go, it can be passed to function to run
	sumFunc := func(x, y int) int {
		return x + y
	}

    result:= executeAndAdd100(sumFunc, 100, 200)
    fmt.Printf("result : %d", result)

}

func executeAndAdd100(fn func(x,y int) int, x int, y int) int {
	firstVal := fn(x,y)
	return firstVal + 100
}
