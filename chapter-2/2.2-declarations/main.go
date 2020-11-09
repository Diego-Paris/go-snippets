package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	fmt.Println("Woot woot")
	boiling()
}

func boiling() {
	f := boilingF
	c := (f - 32) * 5 / 9

	fmt.Printf("Boiling point = %gF or %gC\n", f, c)
	// Output:
	// boiling point = 212F or 100C
	convertFtoc()
}

// ftoc, converts fahrenheit to celsius
func convertFtoc() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%gF = %gC\n", freezingF, ftoc(freezingF)) //"32F = 0C"	
	fmt.Printf("%gF = %gC\n", boilingF, ftoc(boilingF))   //"212F" = 100C"

}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}