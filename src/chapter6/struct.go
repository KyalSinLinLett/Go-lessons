package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

// without using pointer
func circleArea(c Circle) float64 {
	return 3.14 * c.r * c.r
}

func circleAreaPtr(c *Circle) float64 {
	return 3.14 * c.r * c.r
}

// a struct method -> c.area()
func (c *Circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func setRadiusCircle(r float64, c *Circle) {
	c.r = r
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))     // pass by reference
	fmt.Println(circleAreaPtr(&c)) // pass by value

	setRadiusCircle(3, &c)
	fmt.Println(circleAreaPtr(&c))

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
