package main

import "fmt"

type Chargable interface {
	Total() float64
}

type Car struct {
	millege float64
	name string
}

// implement the Chargable interface
func (c Car) Total() float64 {
	return c.millege
}

type AccountLedger struct {
	totalCredit float64
	totalDebit float64
}

// implement the Chargable interface
func (ledger AccountLedger) Total() float64 {
	return ledger.totalCredit - ledger.totalDebit
}

type SomeStruct struct {
	X float64
}

// if the signature is impl via ptr receiver, only ptr variable of that struct could call the interface method
func (s *SomeStruct) Total() float64 {
	return s.X
}


func main(){
	var chargable Chargable

	// Chargable is now nil, would cause null pointer exception (sigsegv)
//	fmt.Println(chargable.Total())

	hondaCar := Car {1000.0, "Honda"}
	accLedger := AccountLedger{3000, 100}

	// same with Java, interface var can hold obj that impl it
	chargable = hondaCar
	fmt.Printf("chargable (interface var) calling its interface method of var hondaCar: %f\n", chargable.Total())
	chargable = accLedger
	fmt.Printf("chargable (interface var) calling its interface method of var accLedger: %f\n", chargable.Total())

	s := SomeStruct{100}
	chargable = &s
	fmt.Printf("chargable (interface var) calling its interface method of var &s: %f\n", chargable.Total())

	describe(chargable)
	describe(hondaCar)
	describe(accLedger)

}

// interface{} = empty interface, it can hold any value
func describe (i interface{} ) {
	fmt.Printf("(%v, %T)\n", i, i)
}