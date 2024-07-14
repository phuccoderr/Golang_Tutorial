package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	phuc := person{
		firstName: "Phuc",
		lastName:  "nguyen",
		contact: contactInfo{
			email: "phuc@gmail",
			zipCode: 911,
		},
	}
	phuc.print()

	var phi person
	phi.firstName = "Con Cho"
	phi.lastName = "Phi"
	phi.contact.email = "concho@gmail.com"
	phi.contact.zipCode = 911
	phi.print()


	phi.updateName("Con Bo")
	phi.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v",p)
	fmt.Println()
}