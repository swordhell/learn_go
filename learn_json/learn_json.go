package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty,-"`
}

func main() {
	person1 := Person{Name: "John Doe", Age: 30, Address: "123 Main St"}
	jsonBytes1, err1 := json.Marshal(person1)
	if err1 != nil {
		// 处理错误
	}
	jsonString1 := string(jsonBytes1)
	fmt.Println(jsonString1) // 输出: {"name":"John Doe","age":30,"address":"123 Main St"}

	person2 := Person{Name: "Jane Doe", Age: 35}
	jsonBytes2, err2 := json.Marshal(person2)
	if err2 != nil {
		// 处理错误
	}
	jsonString2 := string(jsonBytes2)
	fmt.Println(jsonString2) // 输出: {"name":"Jane Doe","age":35}
}
