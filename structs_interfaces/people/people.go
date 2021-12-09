package main

import "fmt"

type People interface {
	SayHello()
	ToString()
}

type Person struct {
	name  string
	age   int
	phone string
}

//A person method
func (p Person) SayHello() {
	fmt.Printf("Hi, I am %s, %d years old\n", p.name, p.age)
}

//A person method
func (p Person) ToString() {
	fmt.Printf("[Name: %s, Age: %d, Phone: %s]\n", p.name, p.age, p.phone)
}

type Student struct {
	Person     //type embedding for composition
	university string
	course     string
}
type Developer struct {
	Person   //type embedding for composition
	company  string
	platform string
}

//Developer's method overrides Person's SayHello
func (d Developer) SayHello() {
	fmt.Printf("Hi, I am %s , %d years old, developer working on %s for %s\n",
		d.name, d.age, d.platform, d.company)
}
func main() {
	alex := Student{Person{"alex", 21, "111-222-XXX"}, "MIT", "BS CS"}
	john := Developer{Person{"John", 35, "111-222-XXX"}, "Accel North America", "Golang"}
	jithesh := Developer{Person{"Jithesh", 33, "111-222-XXX"}, "Accel North America", "Hadoop"}
	//An array with People types
	peopleArr := [...]People{alex, john, jithesh}
	//Iterating through the array of People types and call methods.
	for n, _ := range peopleArr {
		peopleArr[n].SayHello()
		peopleArr[n].ToString()
	}
}
