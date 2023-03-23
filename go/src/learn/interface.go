package learn

import "fmt"

// something that implements the area function
type figures interface {
  area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base float64
	height float64
}

func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

// shared function
func calculateArea(f figures) {
	fmt.Println("Area: ", f.area())
}

func Interface_ex() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}
	calculateArea(mySquare)
	calculateArea(myRectangle)
}