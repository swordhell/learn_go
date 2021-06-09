package main

import "fmt"

type (
	InterfaceA interface {
		FunA(param interface{})
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

func (a *classA) FunA(param interface{}) {
	fmt.Println("classA.FunA() Age ", a.Age)
	i := param.(InterfaceA)
	i.FunB()
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
		i.FunA(i)
	}
}

func testinherit() {
	var i InterfaceA
	i = &classB{}
	i.FunA(i)
}

func main() {
	testinherit()
}
