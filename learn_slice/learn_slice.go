package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 0)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Printf("1 %#v\n", s)

	// 去掉头
	s = s[1:]
	fmt.Println(s)

	// 取出第一个
	s = s[0:1]
	fmt.Printf("2 %#v\n", s)

	// 清空
	s = s[:0]
	fmt.Printf("3 %#v\n", s)

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	// 最后一个元素
	s = s[len(s)-1:]
	fmt.Printf("4 %#v\n", s)
}
