/*
go has 25 key words:
* break	 		default 			func 		interface 	select
* case 			defer 				go 			map 				struct
* chan 			else 					goto 		package 		switch
* const 		fallthrough 	if 			range 			type
* continue 	for 					import 	return 			var

go has three dozen predeclared names:
? Constants: 
*	true 	false	 iota	 nil

? Types: 
* int 			int8 			int16 				int32 		int64
* uint 			uint8 		uint16 				uint32 		uint64 		uintptr
* float32 	float64 	complex128 		complex64
* bool 			byte 			rune 					string 		error

? Functions: 
*	make 			len 		cap 	new 	append 	copy	 close	 delete
*	complex 	real 		imag
*	panic 		recover

!^ These predeclared names are not reserved:
	Sometimes redeclaring them makes sense but be careful
	of potential for confusion

*/
package main

import (
	"fmt"
)

func main() {
	nil := "bobo"
	fmt.Println(nil)
}

func false() string {
	return "I did it"
}