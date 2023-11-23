package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  interface{}
	Height interface{}
}

func (r Rectangle) Area() float64 {
	var height float64
	var Width float64

	if h, ok := r.Height.(float64); ok {
		height = h
	} else {
		fmt.Println("Underline height value is not float64 type")
	}
	if w, ok := r.Width.(float64); ok {
		Width = w
	} else {
		fmt.Println("Underline Width value is not float64 type")
	}

	return height * Width
}

func main() {
	rectangle := Rectangle{
		Width:  5.00,
		Height: 3.00,
	}

	printArea(rectangle)

}

func printArea(shape Shape) {
	fmt.Println("Area", shape.Area())
}
