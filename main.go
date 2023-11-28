package main

import (
	"fmt"
	"reflect"
)

func main() {
	fun := func() {
		fmt.Println("asd")
	}

	fmt.Println(reflect.TypeOf(fun))
}
