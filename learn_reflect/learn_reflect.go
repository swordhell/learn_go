package main

import (
	"fmt"
	"reflect"
)

type X int // 取了一个别名
type Y int

type user struct {
	name string `field:"name" type:"varchar(50)"`
	age  int    `field:"age" type:"int"`
}
type manager struct {
	user
	title string
}

func tNameKind() {

}

func tStruct() {
	var m manager
	t := reflect.TypeOf(m)
	name, _ := t.FieldByName("name")
	fmt.Println(name.Name, name.Type)
	age := t.FieldByIndex([]int{0, 1}) // 第一级索引里面第二个元素
	fmt.Println(age.Name, age.Type)
}

func tCreateBaseData() {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Println(a, m)
}

func tGetStructTag() {
	var u user
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}
}

func main() {
	// tStruct()
	// tCreateBaseData()
	tGetStructTag()
}
