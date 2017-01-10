package main

import "fmt"

// Named collections of method signatures.
type Human interface {
	Speak()
	ToString()
}

type Person struct {
	name         string
	age          int
	emailAddress string
}

func (p Person) Speak() {
	fmt.Printf("Hello. My name is %s and i am %d years old and my email address is %s\n", p.name, p.age, p.emailAddress)
}

func (p Person) ToString() {
	fmt.Printf("[Name: %s, Age: %d, Phone: %s]\n", p.name, p.age, p.emailAddress)
}

type Student struct {
	Person // type embedding for composition
	university string
	division   string
}

type Employee struct {
	Person // type embedding for composition
	company  string
	platform string
}

// Employee's method overrides Person's Speak
func (e Employee) Speak() {
	fmt.Printf("Hello. My name is %s and i am %d years old and my email address is %s\n and i work at %s \n", e.name, e.age, e.emailAddress, e.company)
}

func toStringAHuman(h Human) {
	h.ToString()
}

func main() {
	louis := Student{Person{"Louis", 34, "louis@golang.com"}, "Leicester Uni", "CS"}
	john := Employee{Person{"John", 35, "john@golang.com"}, "Aldi", "Web Team"}
	andrew := Employee{Person{"Andrew", 33, "andrew@golang.com"}, "Aldi", "Sales"}

	allHumans := [...]Human{louis, john, andrew}
	for n := range allHumans {
		allHumans[n].Speak()
		allHumans[n].ToString()
	}

	toStringAHuman(allHumans[0])

}
