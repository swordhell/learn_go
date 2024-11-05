package main

import "fmt"

func main() {
	a := map[string]int64{
		"c": 20,
	}
	a["b"] = 9
	a["a"] += 1

	fmt.Println(a)
}
