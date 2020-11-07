/*
	Reads command line arguments given after invoking the executable,
	uses the os package's Arg value which is a slice of strings,
	where each string is a passed argument to the command
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)


func main() {
	start := time.Now()

	// omit the first string, which is the command executed
	echo(os.Args[1:])

	elapsed := time.Since(start)
	fmt.Println("Finished in", elapsed)
}

// Iterates over all the given string arguments
// and builds up a string using concatenation
func echo(args []string) {
	var s, sep string

	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

// Joins all the arguments into one string using
// the strings package
func echo1(args []string) {
	fmt.Println(strings.Join(args, " "))
}
