package main

import "fmt"

// 定义一个接口
type Shape interface {
    Area() float64
    Proce()
}

// 定义一个基础结构体
type BaseShape struct {
    Name string
}

// 实现接口方法，这里相当于虚函数
func (b *BaseShape) Area() float64 {
    return 0
}

func (b *BaseShape) Proce() float64 {
    return 0
}

// 定义一个继承自 BaseShape 的结构体，并覆盖 Area 方法
type Circle struct {
    BaseShape
    Radius float64
}

// 覆盖基础结构体的 Area 方法
func (c *Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// 定义另一个继承自 BaseShape 的结构体，并覆盖 Area 方法
type Rectangle struct {
    BaseShape
    Width  float64
    Height float64
}

// 覆盖基础结构体的 Area 方法
func (r *Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // 创建一个 Circle 实例
    circle := Circle{BaseShape: BaseShape{Name: "Circle"}, Radius: 5.0}
    // 调用 Area 方法，会调用 Circle 结构体中的 Area 方法
    fmt.Printf("%s Area: %.2f\n", circle.Name, circle.Area())

    // 创建一个 Rectangle 实例
    rectangle := Rectangle{BaseShape: BaseShape{Name: "Rectangle"}, Width: 4.0, Height: 6.0}
    // 调用 Area 方法，会调用 Rectangle 结构体中的 Area 方法
    fmt.Printf("%s Area: %.2f\n", rectangle.Name, rectangle.Area())
}
