package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	p := Point{2, 3}
	fmt.Println(p.String())
	fmt.Println(Point{3, 5}.String())

	fmt.Println("hello world")
	LogOut()
}
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}
