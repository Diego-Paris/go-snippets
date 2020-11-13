package main

/*
Variables, examples of different variable
declarations
*/

import (
	"fmt"
)

func main() {

	//? When we declare variable like this
	//? we can omit the the right side of the
	//? expression, go will add a zero value
	//? which is zero for numbers, "" for strings
	//? and nil for reference types
	var s string
	fmt.Println(s) // ""

	// We can do tuple assignments
	var i, j, k int
	fmt.Println(i, j, k) // 0 0 0

	// We can also do tuple assignments
	// with different types omitting types
	// but we need a start value to be determined
	var b, f, l = true, 3, "four"
	fmt.Printf("%T, %T, %T\n", b, f, l) // bool, int, string

	//! 2.3.1 - Short variable declarations
	a := 100		// int
	d := 100.0 	// float64
	fmt.Printf("%T, %T\n", a, d)

	// multiple variable can be initialized and declared
	// in the same short variable declaration
	e, o := 24, 37
	fmt.Printf("e, o, %T, %T\n", e, o)

	// we can also swap values
	fmt.Printf("e, %v\n", e)
	fmt.Printf("o, %v\n", o)
	e, o = o, e
	fmt.Println("After swap")
	fmt.Printf("e, %v\n", e)
	fmt.Printf("o, %v\n", o)

	// short variable declarations also dont
	// always initialize the variables on the 
	// left sometimes it can reassign
	//*	
	//*	in, err := os.Open(file)
	//*	...
	//* out, err := os.Create(file)
	// we can do this because we are declaring out
	// but we can assign to err, in this case we 
	// need to at least have one new variable declared
	//*	f, err := os.Open(file)
	//*	...
	//* f, err := os.Create(file)
	// the above code won't run but the simple
	// fix would be changing the second := for =

	//! 2.3.2 - Pointers

	// we say &x as "address of x"
	x := 1
	y := &x						// y, of type *int, points to x, contains the address of x
	fmt.Println(*y)		// "1"
	*y = 2						// reassigns the variable indirectly
	fmt.Println(x)		// "2"

	//? two pointers are the same if and only if they 
	//? point to the same variable
}
