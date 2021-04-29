package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 测试批量删除逻辑；
	reg := regexp.MustCompile(` +`)
	var raw_str string = "Gmc  sm  5204"
	fmt.Println(reg.ReplaceAllString(raw_str, " "))
}
