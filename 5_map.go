package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	mapDeclaration()
}

func 

func mapDeclaration(){
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
