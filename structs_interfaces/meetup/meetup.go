package main

import (
	"fmt"
	"time"
)

//https://thenewstack.io/understanding-golang-type-system/
//not that great of an example, the attendee has no field for the conferences he is attending. I also don't like
//that everything is in the same file. I might try to change a few things + add testing(ex. check that the phone number)
//starts with 69 and that is 10 digits after reading a few things from the go-design-principles book.

type People interface {
	SayHello()
	GetDetails()
}

type Person struct {
	name        string
	age         int
	city, phone string
}

func (p Person) SayHello() {
	fmt.Printf("I am %s , %d years old. I live in %s.\n", p.name, p.age, p.city)
}

func (p Person) GetDetails() {
	fmt.Printf("[Name:%s, Age:%d, City:%s, Phone:%s]\n", p.name, p.age, p.city, p.phone)
}

type Speaker struct {
	Person
	speaksOn   []string
	pastEvents []string
}

func (s Speaker) GetDetails() {
	s.Person.GetDetails()
	fmt.Println("Speaker will talk about the following technologies:")
	for _, theme := range s.speaksOn {
		fmt.Println(theme)
	}
	fmt.Println("Speaker has also attented these past events:")
	for _, event := range s.pastEvents {
		fmt.Println(event)
	}
}

type Organizer struct {
	Person
	meetups []string
}

func (o Organizer) GetDetails() {
	o.Person.GetDetails()
	fmt.Println("Organizer, conducting following meetups:")
	for _, meetup := range o.meetups {
		fmt.Println(meetup)
	}
}

type Attendee struct {
	Person
	interests []string
}

func (a Attendee) GetDetails() {
	a.Person.GetDetails()
	fmt.Println("Attendee has the following interests:")
	for _, interest := range a.interests {
		fmt.Println(interest)
	}
}

type Meetup struct {
	location string
	city     string
	date     time.Time
	people   []People
}

func (m Meetup) MeetupPeople() {
	for _, person := range m.people {
		person.SayHello()
		person.GetDetails()
	}
}

func main() {
	nikos := Speaker{Person{"nikos", 25, "athens", "6955555555"}, []string{"Go", "Microservices", "Docker"}, []string{"MS tech days", "FOOEvent", "BAREvent"}}
	giorgos := Organizer{Person{"giorgos", 24, "athens", "6955555553"}, []string{"Goconf", "RubyConf"}}
	alex := Attendee{Person{"alex", 22, "patras", "6955555554"}, []string{"Go", "Ruby"}}
	meetup := Meetup{
		"OAKA",
		"Athens",
		time.Date(2022, time.February, 19, 9, 0, 0, 0, time.UTC),
		[]People{nikos, giorgos, alex},
	}
	meetup.MeetupPeople()
}
