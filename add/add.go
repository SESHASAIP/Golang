package add

import (
	"fmt"
)

var a, b int

func Add() int {
	fmt.Println("enter the vales for addition")
	fmt.Scanln(&a, &b)
	return a + b

}
