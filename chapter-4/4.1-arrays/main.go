package main

import "fmt"

func main() {

	var a [3]int // array of 3 integers, all zero valued

	fmt.Println(a[0]) // print the first element

	fmt.Println(a[len(a)-1]) // print the last element

	// Print indices and values
	for i, v := range a {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}

	// Print elements only
	for _, v := range a {
		fmt.Printf("v = %v\n", v)
	}

}
