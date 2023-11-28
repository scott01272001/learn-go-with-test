package main

import (
	"fmt"
	"scott/learGoWithTest/structs"
)

func main() {
	circle := structs.Circle{Radius: 10.0}
	fmt.Println(circle.Area())

	fmt.Printf("%g \n", circle.Area())
}
