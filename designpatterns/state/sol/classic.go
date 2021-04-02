package main

import "fmt"

type Switch struct {
	State State
}

func (s *Switch) On() {

}

func (s *Switch) Off() {

}

type State interface {
	On(s *Switch)
	Off(s *Switch)
}

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

func main() {

}
