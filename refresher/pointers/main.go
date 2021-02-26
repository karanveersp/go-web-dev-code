package main

import "fmt"

func main() {
	val := 20
	fmt.Println(val)

	ptr := &val

	// print out the address where the value 20 is stored
	fmt.Println(ptr)

	// print the value stored at the address (dereferencing)
	fmt.Println(*ptr)

	*ptr = 10 // changing the value by dereferencing the pointer
	fmt.Println(*ptr == 10 && val == 10)

	val = 50
	fmt.Println(*ptr)
}
