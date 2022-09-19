package main

import (
	"github.com/sirupsen/logrus"
)

/* 定义相互交换值的函数 */
func swap(x, y int) int {
	var temp int
 
	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/
 
	return temp;
 }
 
 /* 定义交换值函数*/
 func swap(x *int, y *int) {
	var temp int
	temp = *x    /* 保持 x 地址上的值 */
	*x = *y      /* 将 y 值赋给 x */
	*y = temp    /* 将 temp 值赋给 y */
 }

func main() {
	var a int = 100
	var b int= 200
	swap(a,b)
	swap(a,b)
}
