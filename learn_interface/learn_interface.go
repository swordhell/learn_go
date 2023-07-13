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

type MyInterface interface {
	MyMethod()
}

type MyStruct1 struct {
	// 结构体字段
}

func (s *MyStruct1) MyMethod() {
	// 实现 MyInterface 中的方法
}

type MyStruct2 struct {
	// 结构体字段
}

func (s MyStruct2) MyMethod() {
	// 实现 MyInterface 中的方法
}

func (a *classA) FunA(param interface{}) {
	fmt.Println("classA.FunA() Age ", a.Age)
	i := param.(InterfaceA)
	i.FunB()
	a.InterfaceA.FunB()
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

	// 使用 MyInterface 类型来接收不同的结构体对象
	var obj1 MyInterface = &MyStruct1{}
	var obj2 MyInterface = MyStruct2{}
	obj1.MyMethod()
	obj2.MyMethod()
}
