package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email    string
	postCode int
}

func main() {

	MrMucus := person{
		firstName: "Yeiling",
		lastName:  "Loops",
		contactInfo: contactInfo{
			email:    "yelpers@screamers.com",
			postCode: 45265,
		},
	}

	MrMucus.changeFirstName("Samovar")

	MrMucus.print()

}

func (p *person) changeFirstName(newfirstName string) {
	p.firstName = newfirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
