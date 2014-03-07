package main

import ("fmt"; "math")

type Rectangle struct {
  x1, y1, x2, y2 float64
}

type Circle struct {
  x, y, r float64
}

type Shape interface {
  area() float64
  perimeter() float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) length() float64 {
  return distance(r.x1, r.y1, r.x1, r.y2)
}

func (r *Rectangle) width() float64 {
  return distance(r.x1, r.y1, r.x2, r.y1)
}

func (r *Rectangle) area() float64 {
  return r.length() * r.width()
}

func (c *Circle) perimeter() float64 {
  return math.Pi * 2 * c.r
}

func (r *Rectangle) perimeter() float64 {
  return r.length()*2 + r.width()*2
}

func main() {
  a := Rectangle{0, 0, 10, 10}
  b := Circle{0, 0, 5}
  fmt.Println(a.area())
  fmt.Println(a.perimeter())
  fmt.Println(b.area())
  fmt.Println(b.perimeter())
}
