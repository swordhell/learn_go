package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 店铺列表
	restaurants := []string{
		"新疆人的面馆",
		"黄焖鸡米饭",
		"猪脚饭",
		"煲仔饭炒菜",
	}

	randomRestaurant := restaurants[rand.Intn(len(restaurants))]
	fmt.Println("今天的午餐推荐是：", randomRestaurant)

}
