package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"
)

type Person struct {
	firstname  string
	lastname   string
	favCountry []string
}
type vehicle struct {
	doors int
	color string
}
type truck struct {
	fourWheel bool
	vehicle
}
type sedan struct {
	luxury bool
	vehicle
}
type SQUARE struct {
	length float64
	width  float64
}
type CIRCLE struct {
	radius float64
}
type SHAPE interface {
	Area() float64
}

func main() {

	m := map[string]Person{}
	pSlice := []Person{
		Person{firstname: "sai",
			lastname: "pamulapati",
			favCountry: []string{
				"usa",
				"india",
				"uk",
			},
		},
		Person{firstname: "Amar",
			lastname: "murukutla",
			favCountry: []string{
				"vietnam",
				"bankkock",
				"uk",
			},
		},
	}
	outputOfSlice(pSlice)
	mappingSliceData(pSlice, m)
	Truck := truck{
		fourWheel: true,
		vehicle: vehicle{
			doors: 4,
			color: "yellow",
		},
	}
	Sedan := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
	}
	fmt.Println(Truck)
	fmt.Println(Truck.color)
	fmt.Println(Truck.doors)
	fmt.Println(Truck.fourWheel)
	fmt.Println(Sedan)
	fmt.Println(Sedan.color)
	fmt.Println(Sedan.doors)
	fmt.Println(Sedan.luxury)

	info(CIRCLE{radius: 9.00})
	info(SQUARE{length: 9.11, width: 8.22})
	var s string
	fmt.Println("enter an url to get data :")
	fmt.Scanln(&s)
	fmt.Println(wordOccurenceCount(createFileAndLoadData(s)))
	// var t SHAPE
	// t =CIRCLE{radius: 9.55}
	// t.Area()

}
func outputOfSlice(pSlice []Person) {

	for c, a := range pSlice {
		if c == 0 {
			fmt.Println("First person struct ", a)
		} else {
			fmt.Println("Second  person struct ", a)
		}
		fmt.Println("Person First Name", a.firstname)
		fmt.Println("Looping over the favourite Country")
		for c, a := range a.favCountry {
			fmt.Println(c, a)
		}
	}

}
func mappingSliceData(pSlice []Person, m map[string]Person) {
	for _, a := range pSlice {
		m[a.lastname] = a

	}

}
func (s SQUARE) Area() float64 {
	return s.length * s.width
}
func (c CIRCLE) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func info(s SHAPE) {
	fmt.Println(s)
	fmt.Println(s.Area())
}
func createFileAndLoadData(s string) string {
	if x, y := http.Get(s); y != nil {
		fmt.Println("can't retrive the data from the url,", y.Error(), x.Status)
		return ""
	} else {
		if f, err := os.Create("assignment6.txt"); err != nil {
			fmt.Println("can't retrive the data from the url,", err.Error())
			return ""
		} else {
			io.Copy(f, x.Body)
			return "assignment6.txt"
		}

	}

}
func wordOccurenceCount(s string) map[string]int {

	var m = map[string]int{}
	if f, err := ioutil.ReadFile(s); err != nil {
		fmt.Println("error while reading the file", err.Error())
		return nil
	} else {

		for _, v := range strings.Split(string(f), " ") {
			if val, ok := m[v]; !ok {
				m[v] = 1
			} else {
				m[v] = val + 1
			}

		}

	}

	return m
}
