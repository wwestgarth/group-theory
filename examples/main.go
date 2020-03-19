package main

import (
	"fmt"
	"groups"
	"reflect"
)

func goperation(a, b groups.Element) groups.Element {

	aval := reflect.ValueOf(a).Int()
	bval := reflect.ValueOf(b).Int()
	return (aval + bval) % 5
}

func main() {

	var set = []groups.Element{0, 1, 2, 3, 4}
	var myop groups.GroupOperation
	myop = goperation
	my := groups.New()
	my.Add(set)

	my.RegisterOperation(&myop)

	fmt.Println(my.Operate(3, 4))
}
