package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is", p.Name)
}

func main() {
	p := new(Person)
	p.Name = "John"
	p.Talk()
	fmt.Println(p)

	a := new(Android)
	a.Person.Name = "Sam"
	a.Model = "lollipop"
	a.Talk()
	fmt.Println(a)
}
