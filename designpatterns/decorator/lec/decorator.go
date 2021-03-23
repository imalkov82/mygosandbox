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

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("circle if radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("square with size %f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s and color %s", c.Shape.Render(), c.Color)
}

type TransperantShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransperantShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency)
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	trsh := TransperantShape{&redCircle, 0.55}
	fmt.Println(trsh.Render())
}
