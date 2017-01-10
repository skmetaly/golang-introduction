package main

import "fmt"

// a way to define our own data type
type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect Rectangle) changeWidth(width float64) float64 {
	rect.width= width

	return rect.width
}

func main() {

	rect1 := Rectangle{leftX:0, topY:50, height:10, width:10}

	rect1.changeWidth(131)
	fmt.Println("Area:",rect1.area())
	fmt.Println(rect1)
}
