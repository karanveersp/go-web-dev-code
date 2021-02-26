package main

import "fmt"

type hotdog int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	var t hotdog
	t = 7
	fmt.Println(t)
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	// fmt.Println(p1)

	// p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	// sa1.speak()
	// sa1.person.speak() // can drill down to type methods

	saySomething(p1)
	saySomething(sa1)
}
