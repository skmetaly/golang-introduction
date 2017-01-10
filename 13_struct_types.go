package main

import "fmt"

// a way to define our own data type
// typed collections of fields
// They're useful for grouping data together to form records.
type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func main() {

	rect1 := Rectangle{leftX:0, topY:50, height:10, width:10}

	// similar with
	// var rect1 Rectangle = Rectangle{leftX:0, topY:50, height:10, width:10}
	// var rect1 Rectangle = Rectangle{0, 50, 10, 10}
	// rect1 = Rectangle{0,50,10,10}

	fmt.Println(rect1)
	fmt.Println("Width:", rect1.width)
}
