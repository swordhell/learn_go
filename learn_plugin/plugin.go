
// plugin.go
package main

import "fmt"

// 声明一个函数，这个函数将在插件中被导出
func Hello() {
    fmt.Println("Hello from the plugin!")
}

// 声明一个导出的变量
var Name string = "PluginName"
