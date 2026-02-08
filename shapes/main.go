package main

import (
	"fmt"
)
	

func main() {
	rect := Rectangle{Width: 4, Height: 5}
	circle := Circle{Radius: 3}
	triangle := Triangle{Base: 6, Height: 4}

	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n\n", rect.Perimeter())

	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n\n", circle.Perimeter())

	fmt.Printf("Triangle Area: %.2f\n", triangle.Area())
	fmt.Printf("Triangle Perimeter: %.2f\n\n", triangle.Perimeter())

	// Scale example
	rect.Scale(2)

	fmt.Println("After scaling rectangle by factor of 2:")
	fmt.Printf("New Width: %.2f, New Height: %.2f\n", rect.Width, rect.Height)
	fmt.Printf("New Area: %.2f\n", rect.Area())
}



