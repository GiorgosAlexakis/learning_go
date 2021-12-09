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

func (p Person) SayHello() {
	fmt.Printf("Hi, I am %s and I am %d years old.\n", p.name, p.age)
}

func (p Person) ToString() {
	fmt.Printf("[Name : %s, Age : %d, Phone : %s]\n", p.name, p.age, p.phone)
}

type Student struct {
	Person
	university string
	major      string
}

func (s Student) SayHello() {
	fmt.Printf("Hello, I am %s, %d years old. I am a student in %s majoring in %s.\n", s.name, s.age, s.university, s.major)
}

type Developer struct {
	Person
	company           string
	favorite_language string
}

func (d Developer) SayHello() {
	fmt.Printf("Hello, I am %s, %d years old. I am a developer for %s and I like writing code in %s.\n", d.name, d.age, d.company, d.favorite_language)
}

func main() {
	alex := Student{Person{"alex", 24, "6934567890"}, "NTUA", "ECE"}
	john := Developer{Person{"john", 23, "6934567891"}, "Alphabet", "Golang"}
	PeopleArr := [...]People{alex, john}
	for n := range PeopleArr {
		PeopleArr[n].SayHello()
		PeopleArr[n].ToString()
	}
}
