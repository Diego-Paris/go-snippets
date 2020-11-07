/*
Takes in multiple filepaths as commandline arguments,
Reads through the files saving each line into a map[string]int,
if the line already exists increment the count instead

Example:
$ go run main.go file1 file2
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Starting...")

	//readByLine()
	readByFile()

	fmt.Println("Ending...")
}

// Reads data from file in one gulp
func readByFile() {
	counts := make(map[string]int)
	
	for _, filename := range os.Args[1:] {

		fmt.Printf("Reading %v\n", filename)
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "An error has ocurred: %v\n", err)
		}

		lines := strings.Split(string(data), "\n")

		for _, line := range lines {
			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Printf("line: %v, total: %v\n", line, n)
	}
}

// Reads each file given in the args line by line
func readByLine() {
	counts := make(map[string]int) // make a map of keys[string] and values[int]
	files := os.Args[1:]           // omit first string

	// iterate over each arg
	for _, arg := range files {
		// use os package to open a file
		f, err := os.Open(arg)

		// if an error ocurred print to stderr
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			continue // move to next arg given
		}

		countLines(f, counts)

		//? we cannot use the defer keyword because that would
		//? close it after the function exits instead of closing
		//? it when we're done with it
		f.Close()
	}
	fmt.Println(counts)
}

func countLines(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)

	// input.Scan() returns true if
	for input.Scan() {
		line := input.Text()
		fmt.Println("just read the line waiting for 5 seconds: " + line)
		time.Sleep(5 * time.Second)
		counts[line]++
	}
	//fmt.Println(counts)
	// NOTE: Ignoring potential errors from input.Err()
}
