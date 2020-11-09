package main

import (
	"fmt"
)

//


// Point this is a point
type Point struct {
	X, Y int
}

func (p *Point) addX(z int) {
	p.X = z
	fmt.Println("modified x", p.X)
}

type geo interface{
	addX(int)
}

func test(g geo) {

	g.addX(2)
}

func main() {
	fmt.Println("Starting...")

	// Control flow
	// switch statement
	// cases do not fall through from one to the next
	// like other C-like languages
	// use the keyword fallthrough for that
	heads := 0
	tails := 0
	switch coinflip() {
	case "heads":
		heads++
		fmt.Println("Heads!")
		fallthrough // allows for the case below to be executed if "heads" is matched
	case "tails":
		tails++
		fmt.Println("Tails!")
	default:
		fmt.Println("landed on an edge!")
	}

	// You can break out of an outer loop but
	// you need to use labels to do so

	// from stackoverflow
	// Use break {label} to break out of any loop
	// as nested as you want. Just put the label
	// before the for loop you want to break out of.
	// This is fairly similar to the code that does
	// a goto {label} but I think a tad more elegant,
	// but matter of opinion I guess.
	// out:
	// 	for i := 0; i < 10; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if i+j == 5 {
	// 				fmt.Println("broke")
	// 				break out
	// 			}
	// 		}
	// 	}

	// example of the function that uses
	// a switch statement
	fmt.Println("value of signum", signum(0))

	// p is of type Point struct
	// a name type that contains two variables
	var p Point
	fmt.Println(p.X)

	// pointers
	//? & yields the address of a variable
	//? * retrieves the variable that pointer refers to
	a := 8
	b := &a
	c := *b
	fmt.Println(a, b, c)

	fmt.Println(p.X)
	p.addX(8)
	fmt.Println("just checking", p.X)
	
	fmt.Println("running test")
	test(&p)
}

func coinflip() string {
	return "heads"
}

func signum(x int) int {

	// equal to switch true {...}
	fmt.Println("in signum")
	switch {

	// since switch above has no value next to it
	// it defaults to true matching the case below
	case true:
		return 42

	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}

}
