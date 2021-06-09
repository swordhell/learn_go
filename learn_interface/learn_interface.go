package main

import "fmt"

type (
	InterfaceA interface {
		FunA()
		FunB()
	}
	classA struct {
		InterfaceA
		Age uint32
	}
	classB struct {
		classA
	}
)

func (a *classA) FunA() {
	fmt.Println("classA.FunA() Age ", a.Age)
	a.FunB()
}

func (a *classA) FunB() {
	fmt.Println("classA.FunB()")
}

func (b *classB) FunB() {
	fmt.Println("classB.FunB()")
}

func testConvert() {
	var b *classA
	if i := InterfaceA(b); i != nil {
		i.FunA()
	}
}

func testinherit() {
	var i InterfaceA
	i = &classB{}
	i.FunA()
}

func main() {
	testConvert()
}
