package main

import "fmt"

func main() {
	fmt.Print("main")

	ages := map[string]int{
		"alice":   31,
		"Charlie": 34,
	}
	ages["alice"] = 32
	fmt.Println(ages["alice"])
	delete(ages, "alices")
}
