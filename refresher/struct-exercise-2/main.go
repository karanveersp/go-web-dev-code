package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	gunLicense bool
}

func (p person) pSpeak() {
	fmt.Println("I'm a person!")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("I'm " + sa.fname + " and I have license to carry firearms")
}

func main() {
	p := person{"Harry", "Osborne"}
	sa := secretAgent{person: person{"Ethan", "Hunt"}, gunLicense: true}

	p.pSpeak()
	sa.saSpeak()
}
