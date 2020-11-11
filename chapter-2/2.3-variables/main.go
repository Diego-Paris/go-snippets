package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2


	v, ok := m["three"]

	fmt.Println(v, ok)


}
