package main

import "fmt"

type (
	InterfaceA interface {
		FunA()
		FunB()
	}

	classA struct {
		InterfaceA
	}

	classB struct {
		classA
	}
)

func (a *classA) FunA() {
	fmt.Println("classA.FunA()")
	a.FunB()
}

func (a *classA) FunB() {
	fmt.Println("classA.FunB()")
}

func (b *classB) FunB() {
	fmt.Println("classB.FunB()")
}

func main() {
	var i InterfaceA
	i = &classB{}
	i.FunA()
}
