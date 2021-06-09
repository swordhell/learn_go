package main

import "fmt"

func f1(a *int32) int32 {
	defer func() {
		*a = 11
	}()
	*a = 10
	return *a
}

func main() {
	var a int32
	var b int32
	a = 9
	b = f1(&a)
	fmt.Println("a ", a, " b ", b)
}
