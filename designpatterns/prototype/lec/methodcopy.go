package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{a.StreetAddress,
		a.City,
		a.Country}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Mat"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Backer St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, jane.Address)
	fmt.Println(jane, jane.Address)
}
