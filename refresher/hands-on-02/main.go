package main

import "fmt"

// #3 types
type person struct {
	fName   string
	lName   string
	favFood []string
}

// #6 methods
func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

// #7
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) description() string {
	isFourWheel := "have four wheel drive"
	if !t.fourWheel {
		isFourWheel = "dont have four wheel drive"
	}
	return fmt.Sprintf("I'm a truck with %d doors and %s", t.doors, isFourWheel)
}

func (s sedan) description() string {
	isLuxury := "am"
	if !s.luxury {
		isLuxury = "am not"
	}
	return fmt.Sprintf("I'm a sedan with %d doors and %s a luxury model", s.doors, isLuxury)
}

// #10
type transportation interface {
	description() string
}

func report(t transportation) {
	fmt.Println(t.description())
}

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
}

func main() {
	// #1 - Slices and looping
	// mySlice := []int{1, 2, 3}
	// fmt.Println(mySlice)

	// for i := range mySlice {
	// 	fmt.Println(i)
	// }

	// for i, j := range mySlice {
	// 	fmt.Println(i, j)
	// }

	// #2 - Maps and looping
	// myMap := map[string]int{"a": 1, "b": 2}
	// fmt.Println(myMap)

	// for key := range myMap {
	// 	fmt.Println(key)
	// }

	// for key, value := range myMap {
	// 	fmt.Println(key, value)
	// }

	// #3 types
	// p1 := person{"Ethan", "Hunt", []string{"Hummus", "Biryani"}}
	// fmt.Println(p1)
	// fmt.Println(p1.fName)
	// fmt.Println(p1.favFood)

	// for i, v := range p1.favFood {
	// 	fmt.Println(i, v)
	// }

	// s := p1.walk()
	// fmt.Println(s)

	// myTruck := truck{
	// 	vehicle{
	// 		2,
	// 		"Red",
	// 	},
	// 	true,
	// }

	// mySedan := sedan{
	// 	vehicle: vehicle{
	// 		doors: 2,
	// 		color: "Beige",
	// 	},
	// 	luxury: true,
	// }

	// fmt.Println(mySedan)
	// fmt.Println(myTruck)
	// // note field promotion
	// fmt.Println(myTruck.color)
	// fmt.Println(mySedan.luxury)

	// fmt.Println(myTruck.description())
	// fmt.Println(mySedan.description())

	// report(myTruck)
	// report(mySedan)

	// #11
	// var g1 gator
	// g1 = 42
	// fmt.Printf("%T\n", g1)

	// var x int
	// fmt.Printf("%d %T\n", x, x)

	// x = int(g1)
	// fmt.Println(x)

	// g1.greeting()
}
