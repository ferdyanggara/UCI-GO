package main

import "fmt"

type shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	base, height int
}

func (t Triangle) Area() float64 {
	return 0.5 * float64(t.base) * float64(t.height)
}

func (t Triangle) Perimeter() float64 {
	return 2*float64(t.base) + 2*float64(t.height)
}

func Visible(s shape2D) bool {
	if s.Area() > 3 && s.Perimeter() > 10 {
		return true
	}
	return false
}

func main() {
	newTri := Triangle{base: 5, height: 2}
	fmt.Println(newTri.Area())
	fmt.Println(newTri.Perimeter())
	fmt.Println(Visible(newTri))
}
