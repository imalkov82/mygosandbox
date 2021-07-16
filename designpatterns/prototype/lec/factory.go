package main

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (e *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)
	_ = enc.Encode(e)

	d := gob.NewDecoder(&b)
	result := &Employee{}
	d.Decode(result)
	return result
}

var mainOffice = Employee{
	"", Address{0, "123 East Dr", "London"}}
var auxOffice = Employee{
	"", Address{0, "66 West Dr", "London"}}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {

}
