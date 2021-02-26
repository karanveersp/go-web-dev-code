/*
Wikipedia: a composite data type (or record) declaration that
defines a physically grouped list of variables to be placed under
one name in a block of memory, allowing the different variables to be
accessed via a single pointer, or the struct declared name which
returns the same address.
*/

package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	fmt.Println(employee{firstName: "Homer", lastName: "Simpson", id: 1})
	fmt.Println(employee{"Whalen", "Smithers", 2})
	fmt.Println(employee{firstName: "Frank", lastName: "Grimes"})

	emp := employee{
		firstName: "Montgomery",
		lastName:  "Burns",
		id:        4,
	}
	fmt.Println(emp.firstName)
	fmt.Println(emp.lastName)
	fmt.Println(emp.id)

	empPtr := &emp
	fmt.Println(empPtr.firstName)

	empPtr.firstName = "Marge"
	fmt.Println(emp.firstName)
}
