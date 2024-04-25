package struct_method_interface

import "math"

type Shape interface {
	Area() float64
}

// Rectangle 长方形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle 三角形结构体
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//func Area(rectangle Rectangle) float64 {
//	return rectangle.Width * rectangle.Height
//}
