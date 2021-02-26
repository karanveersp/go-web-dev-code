package main

import "fmt"

func main() {
	// Maps are similar to Python dictionary (i.e. hashtable, etc.)

	// specify key/value pair
	m := make(map[string]int)
	m["a"] = 0
	m["b"] = 1

	fmt.Println(m)

	// Print the value of a specific key
	fmt.Println(m["a"])
	fmt.Println(m["b"])

	// len() func is overloaded to work with maps
	fmt.Println(len(m))

	// remove key/pair from map (with delete keyword)
	delete(m, "a")
	fmt.Println(m)

	// another way to initialize a map (if you already know values/keys aot)
	m2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(m2)

	// check value and whether the value is present
	val, isValPresent := m["b"]
	fmt.Println(val)
	fmt.Println(isValPresent)

	_, isValPresent2 := m["a"]
	fmt.Println(isValPresent2)
}
