// main.go
package main

import (
	"fmt"
	"plugin"
)

func main() {
	// 打开插件
	p, err := plugin.Open("myplugin.so")
	if err != nil {
		fmt.Println("Failed to load plugin:", err)
		return
	}

	// 查找并使用 Hello 函数
	hello, err := p.Lookup("Hello")
	if err != nil {
		fmt.Println("Failed to find Hello:", err)
		return
	}

	// 通过类型断言转换为函数并调用
	hello.(func())()

	// 查找并使用 Name 变量
	name, err := p.Lookup("Name")
	if err != nil {
		fmt.Println("Failed to find Name:", err)
		return
	}

	// 通过类型断言转换为字符串指针并输出
	fmt.Println("Plugin name:", *name.(*string))
}
