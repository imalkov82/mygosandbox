package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(costcode string) *PersonAddressBuilder {
	b.person.Postcode = costcode
	return b
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) At(company string) *PersonJobBuilder {
	b.person.CompanyName = company
	return b
}

func (b *PersonJobBuilder) As(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Income(income int) *PersonJobBuilder {
	b.person.AnnualIncome = income
	return b
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("Etzel 64").
		In("Holon").
		WithPostcode("SW12BC").
		Works().
		At("Paypal").
		As("Programmer").
		Income(123456)
	person := pb.Build()
	fmt.Println(person)

}
