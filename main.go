package main

import (
	"fmt"
	"reflect"

	"./add"
	"./div"
	"./multi"
	"./sub"
)

var newArr [4]string

func main() {
	var arr [5]int
	var temp int
	temp = len(arr)
	fmt.Println(temp)

	for i := 0; i < len(arr); i++ {
		fmt.Println("enter the value into index:", i)
		fmt.Scanln(&temp)
		arr[i] = temp

	}
	// Using format printing
	for i := 0; i < len(arr); i++ {
		fmt.Printf("the value at index %d is : %d\n", i, arr[i])
	}
	// print out the TYPE of the array
	fmt.Printf("the type of the array is : %v\n", reflect.TypeOf(arr).Elem())

	// Create a SLICE of TYPE int
	var slc []int
	slc = append(slc, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52)
	fmt.Println(slc[0:5])
	fmt.Println(slc[5:10])
	fmt.Println(slc[2:7])
	fmt.Println(slc[1:6])
	y := []int{56, 57, 58, 59, 60}
	y = append(y, slc...)
	fmt.Println("new slice:", y)
	//5. Write a program to pass a pointer to an array as an argument to the function
	newArr = [4]string{"missouri", "michigan", "iowa", "nebraska"}
	newArrPointer := &newArr
	updateThirdElement(newArrPointer)

	fmt.Println(add.Add())
	fmt.Println(sub.Sub())
	fmt.Println(multi.Multi())
	fmt.Println(div.Div())

}
func updateThirdElement(a *[4]string) {
	a[3] = "Texas"
	fmt.Println(*a)

}
