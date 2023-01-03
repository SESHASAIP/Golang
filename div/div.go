package div

import (
	"fmt"
)

var a, b int

func Div() int {
	fmt.Println("enter the values to a and b :")
	fmt.Scanln(&a, &b)
	return a / b
}
