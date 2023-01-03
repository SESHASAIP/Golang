package sub

import "fmt"

var a, b int

func Sub() int {
	fmt.Println("enter the values to a and b :")
	fmt.Scanln(&a, &b)
	return a - b

}
