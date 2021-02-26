package main

import "fmt"

type person struct { // composite type
	fname string
	Lname string // accessible outside package
}

func main() {
	// composite type
	xi := []int{2, 4, 7, 6, 4}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

}
