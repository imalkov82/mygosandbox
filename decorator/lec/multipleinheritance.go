/*

Decorator - Facilitates the additional functionality of
			behaviors to individual objects through embedding

1. Want to augment an object with additionalfunctionality
2. Do not want to rewrite or alter existing code (Open/Close Principle)
3. Want to keep new functionality separate (Single Responsibility Principle)
4. Need to be able to interact with existing structures

Pros:
- Can't implement all the interfaces of wrapped classes - lost functionality

Cons:
- Decorators Can be composed

*/
package main

import "fmt"

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }
func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Print("Flying!")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int       { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }
func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

func (d *Dragon) Age() int { return d.bird.Age() }
func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}
func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}
func main() {
	d := NewDragon()
	d.SetAge(10)
	d.Fly()
	d.Crawl()
}
