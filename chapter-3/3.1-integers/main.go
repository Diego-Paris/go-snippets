package main

import "fmt"

func main() {

	// Go numeric data types include
	// several of:
	/*
		integers
		floating-point numbers
		complex numbers
	*/

	// Go has signed and unsigned
	// 8, 16, 32, 64 bit integers
	/*
		int8
		int16
		int32
		in64
	*/
	// and their unsigned versions
	/*
		uint8
		uint16
		uint32
		uint64
	*/

	// there is also int/uint and that depends
	// on platform so it is either 32 or 64

	// type rune is an alias for int32
	// type byte is an alias for int8

	// there is also uintptr but this is usually
	// used for low-level programming

	fmt.Println("hello world")
}
