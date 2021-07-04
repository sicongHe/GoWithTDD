package shape

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) (ret float64) {
	ret = (rectangle.Width + rectangle.Height) * 2
	return
}

func (r Rectangle) Area() (ret float64) {
	ret = r.Width * r.Height
	return
}

func (c Circle) Area() (ret float64) {
	ret = math.Pi * c.Radius * c.Radius
	return
}

func (t Triangle) Area() (ret float64) {
	ret = t.Base * t.Height / 2
	return
}
